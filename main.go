package main

import (
	"fmt"
	"log"
	"net/http"

	"helloworld/config"
	"helloworld/repositories"
	"helloworld/routes"
	"helloworld/services"

	_ "helloworld/docs" // docs generados por swag
)

// @title           API de Usuarios
// @version         1.0
// @description     API REST para gestión de usuarios siguiendo principios SOLID
// @termsOfService  http://swagger.io/terms/

// @contact.name   Soporte API
// @contact.email  soporte@ejemplo.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @schemes http https
func main() {
	// Cargar configuración
	cfg := config.LoadConfig()

	// Inicializar repositorio MySQL (requerido)
	log.Println("Conectando a MySQL...")
	userRepo, err := repositories.NewMySQLUserRepository(cfg.GetDSN())
	if err != nil {
		log.Fatalf("Error al conectar con MySQL: %v. La aplicación requiere MySQL para funcionar.", err)
	}
	log.Println("Conectado a MySQL exitosamente")

	// Cerrar conexión al finalizar
	defer func() {
		if err := userRepo.Close(); err != nil {
			log.Printf("Error al cerrar conexión MySQL: %v", err)
		}
	}()

	userService := services.NewUserService(userRepo)

	// Configurar rutas
	handler := routes.SetupRoutes(userService)

	// Iniciar servidor
	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("Servidor iniciado en http://localhost%s", addr)
	log.Printf("Health check: http://localhost%s/health", addr)
	log.Printf("API endpoints: http://localhost%s/api/v1/users", addr)
	log.Printf("Swagger UI: http://localhost%s/swagger/index.html", addr)

	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}

