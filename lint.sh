#!/bin/bash

# Script para ejecutar golangci-lint localmente
# Uso: ./lint.sh

echo "Ejecutando golangci-lint..."

# Verificar si golangci-lint está instalado
if ! command -v golangci-lint &> /dev/null; then
    echo "golangci-lint no está instalado."
    echo "Instálalo con: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"
    exit 1
fi

# Ejecutar golangci-lint con los mismos parámetros que en CI
golangci-lint run --timeout=5m --out-format=colored-line-number

if [ $? -eq 0 ]; then
    echo ""
    echo "✅ Linter pasó sin errores!"
else
    echo ""
    echo "❌ Se encontraron errores en el linter."
    exit 1
fi

