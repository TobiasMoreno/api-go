package middleware

import "net/http"

// CORSMiddleware maneja los headers CORS
type CORSMiddleware struct {
	handler http.Handler
}

// NewCORSMiddleware crea una nueva instancia del middleware CORS
func NewCORSMiddleware(handler http.Handler) *CORSMiddleware {
	return &CORSMiddleware{handler: handler}
}

// ServeHTTP implementa http.Handler
func (m *CORSMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	m.handler.ServeHTTP(w, r)
}
