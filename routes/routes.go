package routes

import (
	"net/http"

	"helloworld/handlers"
	"helloworld/middleware"
	"helloworld/services"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// SetupRoutes configura todas las rutas de la API
func SetupRoutes(userService services.UserService) http.Handler {
	router := mux.NewRouter()

	// Inicializar handlers
	userHandler := handlers.NewUserHandler(userService)

	// Rutas de usuarios
	api := router.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	api.HandleFunc("/users", userHandler.GetAllUsers).Methods("GET")
	api.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")
	api.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT")
	api.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")

	// Ruta de health check
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		// nolint:errcheck // Error de escritura en respuesta HTTP, no hay recuperación posible
		_, _ = w.Write([]byte("OK"))
	}).Methods("GET")

	// Ruta de Swagger UI
	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), // URL del archivo JSON generado
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods("GET")

	// Aplicar middleware (orden inverso: primero CORS, luego logging)
	var handler http.Handler = middleware.NewCORSMiddleware(router)
	handler = middleware.NewLoggingMiddleware(handler)

	return handler
}
