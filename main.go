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

	// Inicializar repositorio (MySQL si está configurado, sino memoria)
	var userRepo repositories.UserRepository
	var err error

	// Intentar usar MySQL si está configurado
	if cfg.DBHost != "" && cfg.DBHost != "localhost" || cfg.DBPassword != "" {
		log.Println("Conectando a MySQL...")
		userRepo, err = repositories.NewMySQLUserRepository(cfg.GetDSN())
		if err != nil {
			log.Printf("Error al conectar con MySQL: %v. Usando repositorio en memoria.", err)
			userRepo = repositories.NewInMemoryUserRepository()
		} else {
			log.Println("Conectado a MySQL exitosamente")
			// Cerrar conexión al finalizar
			defer func() {
				if mysqlRepo, ok := userRepo.(*repositories.MySQLUserRepository); ok {
					if err := mysqlRepo.Close(); err != nil {
						log.Printf("Error al cerrar conexión MySQL: %v", err)
					}
				}
			}()
		}
	} else {
		log.Println("Usando repositorio en memoria")
		userRepo = repositories.NewInMemoryUserRepository()
	}

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

