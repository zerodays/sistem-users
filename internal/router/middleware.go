package router

import (
	"github.com/justinas/alice"
	"github.com/rs/zerolog/hlog"
	"github.com/zerodays/sistem-users/internal/logger"
	"net/http"
	"time"
)

// createLoggerMiddleware creates logger middleware.
func createLoggerMiddleware() func(handler http.Handler) http.Handler {
	return hlog.AccessHandler(func(r *http.Request, status, size int, duration time.Duration) {
		hlog.FromRequest(r).Info().
			Str("method", r.Method).
			Stringer("url", r.URL).
			Int("status", status).
			Int("size", size).
			Dur("duration", duration).
			Send()
	})
}

// writeCorsHeader writes CORS header to allow all connections.
func writeCorsHeader(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "DNT,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,Range,Authorization")
}

// corsHandler handles OPTIONS requests. Writes CORS header fields
// and responds with empty body and status 200.
func corsHandler(w http.ResponseWriter, _ *http.Request) {
	writeCorsHeader(w)

	_, _ = w.Write([]byte{})
}

// corsMiddleware writes CORS headers before further processing the request.
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		writeCorsHeader(w)

		next.ServeHTTP(w, r)
	})
}

// responseTypeHeaderMiddleware sets Content-Type header to "application/json" before further processing the request.
func responseTypeHeaderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(w, r)
	})
}

// addMiddleware adds necessary middleware to the handle.
func addMiddleware(handler http.Handler, customContentType bool) http.Handler {
	c := alice.New(hlog.NewHandler(logger.Log), createLoggerMiddleware())

	if !customContentType {
		c = c.Append(responseTypeHeaderMiddleware)
	}

	c = c.Append(corsMiddleware)

	return c.Then(handler)
}
