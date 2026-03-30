package handlers

import (
	"fmt"
	"log"
	"net/http"
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
			if err := recover(); err != nil {
				sentry.CurrentHub().Recover(err)
				sentry.Flush(2 * time.Second)
				jsonResponse(w, http.StatusInternalServerError, "error", "internal server error")
			}
		}()
		next.ServeHTTP(w, r)
	})
}

// LoggingMiddleware logs each request method, path, status, and duration.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		sw := &statusWriter{ResponseWriter: w, status: 200}
		next.ServeHTTP(sw, r)
		log.Printf("[http] %s %s %d %s", r.Method, r.URL.Path, sw.status, time.Since(start).Round(time.Millisecond))
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
}

func NewSiteData() SiteData {
	return SiteData{
		Name:    "Firefly Software",
		Tagline: "Websites, fixes, and software — built for small businesses.",
		Phone:   "+1 (406) 871-9875",
		Email:   "logan@fireflysoftware.dev",
		BaseURL: "https://fireflysoftware.dev",
	}
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
