package handlers

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"runtime/debug"
	"time"

	"github.com/getsentry/sentry-go"
)

// CORSMiddleware validates the Origin header and injects CORS headers.
func CORSMiddleware(allowedOrigin string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")

		if allowedOrigin != "" && origin != allowedOrigin {
			jsonResponse(w, http.StatusForbidden, "error", "forbidden")
			return
		}

		if origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		} else {
			w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
		}
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Max-Age", "86400")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next(w, r)
	}
}

// RecoveryMiddleware catches panics, reports them to Sentry, and returns 500.
func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				sentry.CurrentHub().Recover(rec)
				sentry.Flush(2 * time.Second)

				err := fmt.Errorf("panic: %v", rec)
				SetRequestError(r.Context(), err)
				// The stack goes in a single string field: one event, one line.
				Logger(r.Context()).Error("panic recovered",
					"error", err.Error(),
					"stack", string(debug.Stack()),
				)
				jsonResponse(w, http.StatusInternalServerError, "error", "internal server error")
			}
		}()
		next.ServeHTTP(w, r)
	})
}

// healthPath is polled by the container healthcheck every 30s. Logging it at
// INFO would be pure noise, so it is logged at DEBUG (off in production).
const healthPath = "/api/health"

// LoggingMiddleware emits exactly one structured line per request, and puts a
// request_id and a request-scoped logger into the context so every other line
// logged while serving the request carries the same id.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		requestID := newRequestID()
		state := &requestState{}
		ctx := r.Context()
		ctx = context.WithValue(ctx, ctxKeyRequestID, requestID)
		ctx = context.WithValue(ctx, ctxKeyRequestState, state)
		ctx = context.WithValue(ctx, ctxKeyLogger, slog.Default().With("request_id", requestID))
		r = r.WithContext(ctx)

		w.Header().Set("X-Request-Id", requestID)
		sw := &statusWriter{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(sw, r)

		attrs := []any{
			"request_id", requestID,
			"method", r.Method,
			"path", r.URL.Path,
			"query", r.URL.RawQuery,
			"status", sw.status,
			"duration_ms", float64(time.Since(start).Microseconds()) / 1000.0,
			"remote_ip", realIP(r),
			"user_agent", r.UserAgent(),
			"referer", r.Referer(),
		}

		err := state.getError()
		// 5xx and unhandled errors are a distinct event name, so alerts do not
		// have to regex the status. 4xx stays normal traffic at INFO.
		if sw.status >= 500 || err != nil {
			msg := "unhandled error"
			if err != nil {
				msg = err.Error()
			}
			slog.Error("request failed", append(attrs, "error", msg)...)
			return
		}

		if r.URL.Path == healthPath {
			slog.Debug("request", attrs...)
			return
		}
		slog.Info("request", attrs...)
	})
}

// statusWriter wraps http.ResponseWriter to capture the status code.
type statusWriter struct {
	http.ResponseWriter
	status int
}

func (sw *statusWriter) WriteHeader(code int) {
	sw.status = code
	sw.ResponseWriter.WriteHeader(code)
}

func (sw *statusWriter) Unwrap() http.ResponseWriter {
	return sw.ResponseWriter
}

// Redirect301 returns a handler that redirects to the given path with 301 status.
func Redirect301(to string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, to, http.StatusMovedPermanently)
	}
}

// NotFoundHandler returns a handler that renders the 404 template.
func NotFoundHandler(render func(w http.ResponseWriter, r *http.Request, name string, data any)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		render(w, r, "404", PageData{
			Title:       "Page Not Found",
			Description: "The page you're looking for doesn't exist.",
		})
	}
}

// PageData holds data passed to page templates.
type PageData struct {
	Title       string
	Description string
	Phone       string
	Email       string
	Site        SiteData
	ActiveNav   string
}

// SiteData holds site-wide configuration.
type SiteData struct {
	Name    string
	Tagline string
	Phone   string
	Email   string
	BaseURL string
	CSSFile string
}

func NewSiteData() SiteData {
	return SiteData{
		Name:    "Firefly Software",
		Tagline: "Websites, fixes, and software — built for small businesses.",
		Phone:   "+1 (406) 871-9875",
		Email:   "logan@fireflysoftware.dev",
		BaseURL: "https://fireflysoftware.dev",
		CSSFile: cssFilePath("static"),
	}
}

// cssFilePath computes a fingerprinted CSS filename from static/css/output.css.
// Returns "/static/css/output.<hash>.css" in production, or "/static/css/output.css" as fallback.
func cssFilePath(staticDir string) string {
	path := staticDir + "/css/output.css"
	f, err := os.Open(path)
	if err != nil {
		slog.Warn("css fingerprint failed", "path", path, "error", err.Error(), "fallback", "/static/css/output.css")
		return "/static/css/output.css"
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		slog.Warn("css fingerprint failed", "path", path, "error", err.Error(), "fallback", "/static/css/output.css")
		return "/static/css/output.css"
	}

	hash := fmt.Sprintf("%x", h.Sum(nil))[:8]
	hashedName := fmt.Sprintf("/static/css/output.%s.css", hash)

	// Create the hashed file on disk so the static file server can serve it
	src := staticDir + "/css/output.css"
	dst := fmt.Sprintf("%s/css/output.%s.css", staticDir, hash)
	if _, err := os.Stat(dst); os.IsNotExist(err) {
		data, err := os.ReadFile(src)
		if err != nil {
			slog.Warn("css fingerprint failed", "path", src, "error", err.Error(), "fallback", "/static/css/output.css")
			return "/static/css/output.css"
		}
		if err := os.WriteFile(dst, data, 0644); err != nil {
			slog.Warn("css fingerprint failed", "path", dst, "error", err.Error(), "fallback", "/static/css/output.css")
			return "/static/css/output.css"
		}
		slog.Info("css fingerprinted", "path", dst)
	}

	return hashedName
}

func NewPageData(title, description string, site SiteData) PageData {
	return PageData{
		Title:       title,
		Description: description,
		Phone:       site.Phone,
		Email:       site.Email,
		Site:        site,
	}
}

// WithNav returns a copy of PageData with ActiveNav set.
func (p PageData) WithNav(nav string) PageData {
	p.ActiveNav = nav
	return p
}

// PhoneHref returns the tel: link for the phone number.
func (s SiteData) PhoneHref() string {
	return fmt.Sprintf("tel:+14068719875")
}
