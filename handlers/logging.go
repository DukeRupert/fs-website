package handlers

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"log/slog"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
)

type ctxKey int

const (
	ctxKeyLogger ctxKey = iota
	ctxKeyRequestID
	ctxKeyRequestState
)

// InitLogger installs a JSON slog handler writing one object per line to
// stdout, and makes it the process default. slog's JSON handler emits `time`
// (RFC3339 with nanoseconds), uppercase `level`, and `msg` natively, which is
// exactly the fleet logging standard.
//
// slog.SetDefault also redirects the standard library `log` package through
// this handler, so any stray log.Printf in a dependency becomes JSON rather
// than plain text on the same stdout.
func InitLogger(debug bool) *slog.Logger {
	level := slog.LevelInfo
	if debug {
		level = slog.LevelDebug
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			// slog emits RFC3339 with nanoseconds already, but in the local
			// zone. The standard requires UTC, so normalise it here rather than
			// depending on the container's TZ.
			if len(groups) == 0 && a.Key == slog.TimeKey && a.Value.Kind() == slog.KindTime {
				a.Value = slog.TimeValue(a.Value.Time().UTC())
			}
			return a
		},
	}))
	slog.SetDefault(logger)
	return logger
}

// newRequestID returns a short random hex id for correlating every line
// emitted while serving one request.
func newRequestID() string {
	var b [8]byte
	if _, err := rand.Read(b[:]); err != nil {
		return "0000000000000000"
	}
	return hex.EncodeToString(b[:])
}

// realIP returns the client IP address, without a port.
//
// The Go server runs in Docker behind a host-level Caddy reverse proxy, so the
// socket address is always the bridge network (172.x / 192.168.x) and is
// useless. The visitor address arrives in X-Forwarded-For.
//
// Trusted hops: 1 — the host-level Caddy. The in-container Caddy passes
// X-Forwarded-For through unchanged (see Caddyfile), so the right-most entry is
// the one our own proxy wrote. Take the right-most entry, never the left-most:
// anything to the left of our proxy's entry was supplied by the caller and can
// be forged. Today the host Caddy runs with trusted_proxies unset, so it
// replaces the header and there is exactly one entry — but that flips the
// moment a CDN is put in front, and right-most stays correct either way.
func realIP(r *http.Request) string {
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		parts := strings.Split(xff, ",")
		candidate := strings.TrimSpace(parts[len(parts)-1])
		if candidate != "" {
			return stripPort(candidate)
		}
	}
	return stripPort(r.RemoteAddr)
}

// stripPort removes a trailing :port from an address if present.
func stripPort(addr string) string {
	if host, _, err := net.SplitHostPort(addr); err == nil {
		return host
	}
	return addr
}

// requestState carries per-request data that handlers can contribute to and
// the logging middleware reads when emitting the single request line.
type requestState struct {
	mu  sync.Mutex
	err error
}

func (rs *requestState) setError(err error) {
	rs.mu.Lock()
	defer rs.mu.Unlock()
	if rs.err == nil {
		rs.err = err
	}
}

func (rs *requestState) getError() error {
	rs.mu.Lock()
	defer rs.mu.Unlock()
	return rs.err
}

// SetRequestError records an error against the in-flight request. The logging
// middleware promotes the request line to msg="request failed" at ERROR and
// attaches the error string.
func SetRequestError(ctx context.Context, err error) {
	if err == nil {
		return
	}
	if rs, ok := ctx.Value(ctxKeyRequestState).(*requestState); ok {
		rs.setError(err)
	}
}

// Logger returns the request-scoped logger, which carries request_id. Falls
// back to the default logger outside of a request.
func Logger(ctx context.Context) *slog.Logger {
	if l, ok := ctx.Value(ctxKeyLogger).(*slog.Logger); ok {
		return l
	}
	return slog.Default()
}

// RequestID returns the id assigned to the in-flight request, or "".
func RequestID(ctx context.Context) string {
	if id, ok := ctx.Value(ctxKeyRequestID).(string); ok {
		return id
	}
	return ""
}
