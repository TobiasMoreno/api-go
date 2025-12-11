#!/bin/bash

echo "Generando documentación Swagger..."
swag init -g main.go -o ./docs

echo ""
echo "Documentación generada en ./docs"
echo ""
echo "Ahora puedes ejecutar: go run main.go"
echo "Y acceder a Swagger UI en: http://localhost:8080/swagger/index.html"

