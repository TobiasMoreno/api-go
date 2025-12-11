# API REST con Go - Principios SOLID

API REST desarrollada en Go siguiendo principios SOLID, especialmente el Single Responsibility Principle (SRP).

## Estructura del Proyecto

```
.
├── config/          # Configuración de la aplicación
├── handlers/        # Manejo de peticiones HTTP
├── middleware/      # Middleware (CORS, Logging)
├── models/          # Modelos de datos
├── repositories/    # Capa de acceso a datos
├── routes/          # Configuración de rutas
├── services/        # Lógica de negocio
├── docs/            # Documentación Swagger (generada)
├── main.go          # Punto de entrada
└── go.mod           # Dependencias

```

## Instalación

1. Clonar el repositorio
```bash
git clone https://github.com/tu-usuario/go-api.git
cd go-api
```

2. Copiar archivo de entorno:
```bash
# Opción 1: Si tienes .env.example
cp .env.example .env

# Opción 2: Si tienes env.example
cp env.example .env

# Editar .env con tus configuraciones
```

3. Instalar dependencias:
```bash
make deps
# O manualmente:
go mod download
go mod tidy
```

4. Instalar herramientas de desarrollo (opcional):
```bash
make install-tools
```

**Nota**: Si usas MySQL, asegúrate de tener el driver instalado. Se instala automáticamente con `make deps`.

3. Instalar Swagger CLI (solo la primera vez):
```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

4. Generar documentación Swagger:
```bash
make swagger
# O manualmente:
swag init -g main.go -o ./docs
```

## Ejecución

### Opción 1: Ejecutar localmente

```bash
# Ejecutar la aplicación
make run
# O
go run main.go
```

El servidor se iniciará en `http://localhost:8080`

### Opción 2: Usando Docker (con MySQL)

#### Construir y ejecutar con Docker Compose

```bash
# Construir y ejecutar (incluye MySQL)
docker-compose up --build

# Ejecutar en segundo plano
docker-compose up -d

# Ver logs de la API
docker-compose logs -f api

# Ver logs de MySQL
docker-compose logs -f mysql

# Detener
docker-compose down

# Detener y eliminar volúmenes (elimina datos de BD)
docker-compose down -v
```

**Configuración de MySQL:**
- Host: `mysql` (dentro de Docker) o `localhost` (desde fuera)
- Puerto: `3306`
- Base de datos: `usersdb`
- Usuario: `appuser`
- Contraseña: `apppassword`
- Root password: `rootpassword`

#### Usando Docker directamente

```bash
# Construir la imagen
docker build -t go-api .

# Ejecutar el contenedor
docker run -p 8080:8080 --name go-api go-api

# Ejecutar con variable de entorno personalizada
docker run -p 8080:8080 -e PORT=3000 --name go-api go-api
```

#### Desarrollo con Docker

```bash
# Usar docker-compose para desarrollo
docker-compose -f docker-compose.dev.yml up --build
```

## Documentación Swagger

Una vez generada la documentación, puedes acceder a la interfaz de Swagger UI en:

**http://localhost:8080/swagger/index.html**

## Endpoints

### Usuarios

- `POST /api/v1/users` - Crear usuario
- `GET /api/v1/users` - Obtener todos los usuarios
- `GET /api/v1/users/{id}` - Obtener usuario por ID
- `PUT /api/v1/users/{id}` - Actualizar usuario
- `DELETE /api/v1/users/{id}` - Eliminar usuario

### Health Check

- `GET /health` - Verificar estado del servidor

## Ejemplo de Uso

### Crear un usuario

```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Juan Pérez",
    "email": "juan@example.com",
    "age": 30
  }'
```

### Obtener todos los usuarios

```bash
curl http://localhost:8080/api/v1/users
```

### Obtener un usuario por ID

```bash
curl http://localhost:8080/api/v1/users/{id}
```

## Principios Aplicados

- **SRP (Single Responsibility Principle)**: Cada paquete tiene una única responsabilidad
- **Dependency Inversion**: Uso de interfaces para desacoplar capas
- **Separation of Concerns**: Separación clara entre handlers, services y repositories

## Comandos Make

```bash
make help          # Ver todos los comandos disponibles
make deps          # Instalar dependencias
make run           # Ejecutar aplicación
make test          # Ejecutar tests
make lint          # Ejecutar linter
make build         # Compilar aplicación
make docker-up     # Ejecutar con Docker
make ci            # Ejecutar pipeline CI local
```

Ver todos los comandos: `make help`

## Testing

```bash
# Ejecutar todos los tests
make test

# Tests con cobertura
make test-coverage

# Tests rápidos
make test-short
```

## Linting

```bash
# Verificar código
make lint

# Corregir automáticamente
make lint-fix

# Formatear código
make fmt
```

## CI/CD

El proyecto incluye GitHub Actions para:
- ✅ Linting automático
- ✅ Tests automáticos
- ✅ Build de Docker
- ✅ Security scanning
- ✅ Release automático

Ver workflows en `.github/workflows/`

## Tecnologías

- Go 1.21+
- Gorilla Mux (router)
- Swagger/OpenAPI (documentación)
- UUID (generación de IDs)
- MySQL 8.0 (base de datos)
- Docker & Docker Compose

## Base de Datos

La aplicación soporta dos modos de almacenamiento:

### 1. MySQL (Producción/Recomendado)
- Se usa automáticamente cuando se configuran las variables de entorno de BD
- Persistencia de datos
- Escalable y robusto

### 2. Memoria (Desarrollo/Testing)
- Se usa cuando no hay configuración de MySQL
- Datos se pierden al reiniciar
- Útil para desarrollo rápido

### Variables de Entorno para MySQL

```bash
DB_HOST=mysql          # Host de MySQL
DB_PORT=3306          # Puerto de MySQL
DB_USER=appuser        # Usuario de MySQL
DB_PASSWORD=apppassword # Contraseña
DB_NAME=usersdb        # Nombre de la base de datos
```

### Conectar a MySQL desde fuera de Docker

```bash
# Usando mysql client
mysql -h localhost -P 3306 -u appuser -papppassword usersdb

# O usando cualquier cliente MySQL (DBeaver, MySQL Workbench, etc.)
# Host: localhost
# Port: 3306
# User: appuser
# Password: apppassword
# Database: usersdb
```

## Docker

### Características del Dockerfile

- **Multi-stage build**: Reduce el tamaño de la imagen final
- **Imagen Alpine**: Imagen base mínima (~5MB)
- **Usuario no-root**: Ejecuta con permisos limitados para seguridad
- **Healthcheck**: Verifica el estado del servicio automáticamente
- **Optimizado**: Binario estático sin dependencias CGO

### Comandos Docker útiles

```bash
# Ver logs del contenedor
docker-compose logs -f api

# Ejecutar comandos dentro del contenedor
docker-compose exec api sh

# Reconstruir sin cache
docker-compose build --no-cache

# Ver estado de los contenedores
docker-compose ps

# Limpiar (elimina contenedores, redes, volúmenes)
docker-compose down -v
```

