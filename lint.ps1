# Script para ejecutar golangci-lint localmente
# Uso: .\lint.ps1

Write-Host "Ejecutando golangci-lint..." -ForegroundColor Green

# Verificar si golangci-lint está instalado
if (-not (Get-Command golangci-lint -ErrorAction SilentlyContinue)) {
    Write-Host "golangci-lint no está instalado." -ForegroundColor Red
    Write-Host "Instálalo con: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest" -ForegroundColor Yellow
    exit 1
}

# Ejecutar golangci-lint con los mismos parámetros que en CI
golangci-lint run --timeout=5m --out-format=colored-line-number

if ($LASTEXITCODE -eq 0) {
    Write-Host "`n✅ Linter pasó sin errores!" -ForegroundColor Green
} else {
    Write-Host "`n❌ Se encontraron errores en el linter." -ForegroundColor Red
    exit 1
}

