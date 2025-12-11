.PHONY: help swagger run build deps install-swag test lint docker-build docker-up docker-down docker-logs docker-clean docker-rebuild

# Variables
BINARY_NAME=api
DOCKER_IMAGE=go-api

# Colores para output
GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
RESET  := $(shell tput -Txterm sgr0)

help: ## Mostrar esta ayuda
	@echo "$(GREEN)Comandos disponibles:$(RESET)"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  $(YELLOW)%-20s$(RESET) %s\n", $$1, $$2}'

# ============================================
# Desarrollo
# ============================================

deps: ## Instalar dependencias
	@echo "$(GREEN)Instalando dependencias...$(RESET)"
	@go mod download
	@go mod tidy

install-swag: ## Instalar Swagger CLI
	@echo "$(GREEN)Instalando Swagger CLI...$(RESET)"
	@go install github.com/swaggo/swag/cmd/swag@latest

swagger: ## Generar documentación Swagger
	@echo "$(GREEN)Generando documentación Swagger...$(RESET)"
	@swag init -g main.go -o ./docs
	@echo "$(GREEN)Documentación generada en ./docs$(RESET)"

run: ## Ejecutar la aplicación
	@echo "$(GREEN)Ejecutando aplicación...$(RESET)"
	@go run main.go

build: ## Compilar la aplicación
	@echo "$(GREEN)Compilando aplicación...$(RESET)"
	@go build -ldflags="-s -w" -o bin/$(BINARY_NAME) main.go
	@echo "$(GREEN)Binario creado en bin/$(BINARY_NAME)$(RESET)"

# ============================================
# Testing
# ============================================

test: ## Ejecutar tests
	@echo "$(GREEN)Ejecutando tests...$(RESET)"
	@go test -v -race -coverprofile=coverage.out ./...

test-coverage: test ## Ejecutar tests con cobertura
	@echo "$(GREEN)Generando reporte de cobertura...$(RESET)"
	@go tool cover -html=coverage.out -o coverage.html
	@echo "$(GREEN)Reporte generado en coverage.html$(RESET)"

test-short: ## Ejecutar tests rápidos
	@go test -short ./...

# ============================================
# Linting
# ============================================

lint: ## Ejecutar linter
	@echo "$(GREEN)Ejecutando golangci-lint...$(RESET)"
	@golangci-lint run

lint-fix: ## Ejecutar linter y corregir automáticamente
	@echo "$(GREEN)Corrigiendo código con linter...$(RESET)"
	@golangci-lint run --fix

fmt: ## Formatear código
	@echo "$(GREEN)Formateando código...$(RESET)"
	@go fmt ./...
	@goimports -w .

vet: ## Ejecutar go vet
	@echo "$(GREEN)Ejecutando go vet...$(RESET)"
	@go vet ./...

# ============================================
# Docker
# ============================================

docker-build: ## Construir imagen Docker
	@echo "$(GREEN)Construyendo imagen Docker...$(RESET)"
	@docker-compose build

docker-up: ## Ejecutar con docker-compose
	@echo "$(GREEN)Iniciando contenedores...$(RESET)"
	@docker-compose up -d
	@echo "$(GREEN)API disponible en http://localhost:8080$(RESET)"

docker-down: ## Detener contenedores
	@echo "$(GREEN)Deteniendo contenedores...$(RESET)"
	@docker-compose down

docker-logs: ## Ver logs
	@docker-compose logs -f api

docker-clean: ## Limpiar recursos Docker
	@echo "$(GREEN)Limpiando recursos Docker...$(RESET)"
	@docker-compose down -v
	@docker rmi $$(docker images -q $(DOCKER_IMAGE)) 2>/dev/null || true
	@echo "$(GREEN)Limpieza completada$(RESET)"

docker-rebuild: ## Reconstruir y ejecutar
	@docker-compose up --build -d

# ============================================
# CI/CD
# ============================================

ci: lint test ## Ejecutar pipeline CI local
	@echo "$(GREEN)✅ CI completado exitosamente$(RESET)"

# ============================================
# Utilidades
# ============================================

clean: ## Limpiar archivos generados
	@echo "$(GREEN)Limpiando archivos...$(RESET)"
	@rm -rf bin/
	@rm -f coverage.out coverage.html
	@go clean

.PHONY: install-tools
install-tools: ## Instalar herramientas de desarrollo
	@echo "$(GREEN)Instalando herramientas...$(RESET)"
	@go install github.com/swaggo/swag/cmd/swag@latest
	@go install golang.org/x/tools/cmd/goimports@latest
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@echo "$(GREEN)Herramientas instaladas$(RESET)"
