# Etapa 1: Build - Compilar la aplicación
FROM golang:1.21-alpine AS builder

# Instalar dependencias del sistema necesarias para compilar
RUN apk add --no-cache git

# Establecer directorio de trabajo
WORKDIR /app

# Copiar archivos de dependencias
COPY go.mod go.sum ./

# Descargar dependencias (se cachean si no cambian go.mod/go.sum)
RUN go mod download

# Copiar el código fuente
COPY . .

# Generar documentación Swagger (si existe swag)
RUN go install github.com/swaggo/swag/cmd/swag@latest || true
RUN swag init -g main.go -o ./docs || true

# Compilar la aplicación
# CGO_ENABLED=0: deshabilita CGO para crear binario estático
# -ldflags="-s -w": reduce el tamaño del binario
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/api main.go

# Etapa 2: Runtime - Imagen final mínima
FROM alpine:latest

# Instalar ca-certificates y wget para healthcheck
RUN apk --no-cache add ca-certificates wget

# Crear usuario no-root para seguridad
RUN addgroup -g 1000 appuser && \
    adduser -D -u 1000 -G appuser appuser

WORKDIR /app

# Copiar el binario compilado desde la etapa builder
COPY --from=builder /app/api .

# Cambiar ownership al usuario no-root
RUN chown -R appuser:appuser /app

# Cambiar al usuario no-root
USER appuser

# Exponer el puerto
EXPOSE 8080

# Healthcheck
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://localhost:8080/health || exit 1

# Variable de entorno por defecto
ENV PORT=8080

# Comando para ejecutar la aplicación
CMD ["./api"]

