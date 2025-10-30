# Script para ejecutar el proyecto en Windows PowerShell

param(
    [string]$Command = "help"
)

function Show-Help {
    Write-Host "======================================"
    Write-Host "Comandos disponibles:"
    Write-Host "======================================"
    Write-Host ""
    Write-Host "  .\run.ps1 build       - Construir imágenes Docker"
    Write-Host "  .\run.ps1 up          - Iniciar servicios"
    Write-Host "  .\run.ps1 down        - Detener servicios"
    Write-Host "  .\run.ps1 logs        - Ver logs"
    Write-Host "  .\run.ps1 test        - Ejecutar pruebas"
    Write-Host "  .\run.ps1 clean       - Limpiar todo"
    Write-Host "  .\run.ps1 health      - Verificar estado"
    Write-Host ""
}

function Build-Images {
    Write-Host "Construyendo imágenes Docker..." -ForegroundColor Green
    docker-compose build
}

function Start-Services {
    Write-Host "Iniciando servicios..." -ForegroundColor Green
    docker-compose up -d
    Write-Host ""
    Write-Host "Servicios iniciados:" -ForegroundColor Green
    Write-Host "  Go API: http://localhost:8080"
    Write-Host "  Node.js API: http://localhost:3000"
    Write-Host ""
    Write-Host "Esperando a que los servicios estén listos..." -ForegroundColor Yellow
    Start-Sleep -Seconds 5
    Write-Host "Servicios listos!" -ForegroundColor Green
}

function Stop-Services {
    Write-Host "Deteniendo servicios..." -ForegroundColor Green
    docker-compose down
}

function Show-Logs {
    Write-Host "Mostrando logs..." -ForegroundColor Green
    docker-compose logs -f
}

function Run-Tests {
    Write-Host "Ejecutando pruebas..." -ForegroundColor Green
    Write-Host ""
    
    Write-Host "Pruebas de Go:" -ForegroundColor Cyan
    Set-Location go-api
    go test -v
    Set-Location ..
    
    Write-Host ""
    Write-Host "Pruebas de Node.js:" -ForegroundColor Cyan
    Set-Location node-api
    node test.js
    Set-Location ..
}

function Clean-All {
    Write-Host "Limpiando..." -ForegroundColor Green
    docker-compose down -v
    Remove-Item -Path "go-api/main" -ErrorAction SilentlyContinue
    Remove-Item -Path "node_modules" -Recurse -ErrorAction SilentlyContinue
    Write-Host "Limpieza completada!" -ForegroundColor Green
}

function Check-Health {
    Write-Host "Verificando estado de los servicios..." -ForegroundColor Green
    Write-Host ""
    
    Write-Host "Go API (8080):" -ForegroundColor Cyan
    try {
        $response = Invoke-WebRequest -Uri "http://localhost:8080/health" -UseBasicParsing
        Write-Host $response.Content -ForegroundColor Green
    } catch {
        Write-Host "No disponible" -ForegroundColor Red
    }
    
    Write-Host ""
    Write-Host "Node.js API (3000):" -ForegroundColor Cyan
    try {
        $response = Invoke-WebRequest -Uri "http://localhost:3000/health" -UseBasicParsing
        Write-Host $response.Content -ForegroundColor Green
    } catch {
        Write-Host "No disponible" -ForegroundColor Red
    }
}

# Execute command
switch ($Command.ToLower()) {
    "build" { Build-Images }
    "up" { Start-Services }
    "down" { Stop-Services }
    "logs" { Show-Logs }
    "test" { Run-Tests }
    "clean" { Clean-All }
    "health" { Check-Health }
    default { Show-Help }
}
