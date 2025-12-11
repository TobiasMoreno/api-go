# Script para formatear todos los archivos Go
# Uso: .\format.ps1

Write-Host "Formateando archivos Go..." -ForegroundColor Green

# Formatear todos los archivos .go (excepto docs que es generado)
Get-ChildItem -Recurse -Filter "*.go" | Where-Object { 
    $_.FullName -notlike "*\docs\*" 
} | ForEach-Object {
    Write-Host "Formateando: $($_.Name)" -ForegroundColor Yellow
    gofmt -w $_.FullName
}

Write-Host "`n✅ Archivos formateados!" -ForegroundColor Green

