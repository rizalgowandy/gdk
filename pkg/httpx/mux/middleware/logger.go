package middleware

import (
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/rizalgowandy/gdk/pkg/env"
	"github.com/rizalgowandy/gdk/pkg/httpx/mux"
	"github.com/sirupsen/logrus" // nolint:depguard
)

type timer interface {
	Now() time.Time
	Since(time.Time) time.Duration
}

// realClock save request times
type realClock struct{}

func (rc *realClock) Now() time.Time {
	return time.Now()
}

func (rc *realClock) Since(t time.Time) time.Duration {
	return time.Since(t)
}

// LogOptions logging middleware options
type LogOptions struct {
	Formatter logrus.Formatter
}

// LoggingMiddleware is a middleware handler that logs the request as it goes in and the response as it goes out.
type LoggingMiddleware struct {
	logger         *logrus.Logger
	clock          timer
	enableStarting bool
}

// NewLogger returns a new *LoggingMiddleware, yay!
func NewLogger(opts ...LogOptions) *LoggingMiddleware {
	var opt LogOptions
	if len(opts) == 0 {
		opt = LogOptions{}
	} else {
		opt = opts[0]
	}

	if opt.Formatter == nil {
		opt.Formatter = &logrus.JSONFormatter{
			TimestampFormat:   time.RFC3339,
			DisableTimestamp:  false,
			DisableHTMLEscape: !env.IsProduction(),
			DataKey:           "",
			FieldMap:          nil,
			CallerPrettyfier:  nil,
			PrettyPrint:       false,
		}
	}

	log := logrus.New()
	log.Formatter = opt.Formatter

	return &LoggingMiddleware{
		logger:         log,
		clock:          &realClock{},
		enableStarting: !env.IsProduction(),
	}
}

// realIP get the real IP from http request
func realIP(req *http.Request) string {
	ra := req.RemoteAddr
	if ip := req.Header.Get(mux.HeaderXForwardedFor); ip != "" {
		ra = strings.Split(ip, ", ")[0]
	} else if ip := req.Header.Get(mux.HeaderXRealIP); ip != "" {
		ra = ip
	} else {
		ra, _, _ = net.SplitHostPort(ra)
	}
	return ra
}

// Middleware implement mux middleware interface
func (m *LoggingMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		entry := logrus.NewEntry(m.logger)
		start := m.clock.Now()

		if reqID := r.Header.Get(mux.HeaderXRequestID); reqID != "" {
			entry = entry.WithField("request_id", reqID)
		}

		if remoteAddr := realIP(r); remoteAddr != "" {
			entry = entry.WithField("remote_addr", remoteAddr)
		}

		if m.enableStarting {
			entry.WithFields(logrus.Fields{
				"request": r.RequestURI,
				"method":  r.Method,
			}).Info(fmt.Sprintf("operation %s starting", r.URL.Path))
		}

		lw := mux.NewWriter(w)
		next.ServeHTTP(lw, r)

		if !env.IsProduction() {
			entry = entry.WithField("response", lw.Response)
		}

		entry.WithFields(logrus.Fields{
			"request_uri": r.RequestURI,
			"status_code": lw.StatusCode,
			"took":        m.clock.Since(start).String(),
		}).Info(fmt.Sprintf("operation %s result", r.URL.Path))
	})
}
