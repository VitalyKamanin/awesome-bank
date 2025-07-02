package middleware

import (
	"awesome-bank/services/utils"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		ww := NewWrappedResponseWriter(w)

		next.ServeHTTP(ww, r)

		utils.Logger.WithField("method", r.Method).
			WithField("path", r.URL.Path).
			WithField("ip", r.RemoteAddr).
			WithField("status", ww.Status).
			WithField("duration", time.Since(start).String()).
			Info("Request processed")
	})
}

type WrappedResponseWriter struct {
	http.ResponseWriter
	Status int
}

func NewWrappedResponseWriter(w http.ResponseWriter) *WrappedResponseWriter {
	return &WrappedResponseWriter{ResponseWriter: w, Status: http.StatusOK}
}

func (ww *WrappedResponseWriter) WriteHeader(code int) {
	ww.Status = code
	ww.ResponseWriter.WriteHeader(code)
}
