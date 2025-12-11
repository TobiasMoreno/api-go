package middleware

import (
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware registra información sobre las peticiones HTTP
type LoggingMiddleware struct {
	handler http.Handler
}

// NewLoggingMiddleware crea una nueva instancia del middleware de logging
func NewLoggingMiddleware(handler http.Handler) *LoggingMiddleware {
	return &LoggingMiddleware{handler: handler}
}

// ServeHTTP implementa http.Handler
func (m *LoggingMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	// Crear un ResponseWriter personalizado para capturar el status code
	wrapped := &responseWriter{
		ResponseWriter: w,
		statusCode:     http.StatusOK,
	}

	m.handler.ServeHTTP(wrapped, r)

	duration := time.Since(start)
	log.Printf(
		"%s %s %s %d %v",
		r.Method,
		r.RequestURI,
		r.RemoteAddr,
		wrapped.statusCode,
		duration,
	)
}

// responseWriter envuelve http.ResponseWriter para capturar el status code
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}
