# Guía de Desarrollo Local

Esta guía te ayudará a ejecutar el proyecto sin Docker en tu máquina local.

## Requisitos

### Para la API Go
- Go 1.21 o superior
- [Descargar Go](https://golang.org/dl/)

### Para la API Node.js
- Node.js 18 o superior
- npm o yarn
- [Descargar Node.js](https://nodejs.org/)

## Instalación

### 1. Verificar instalaciones

```bash
# Verificar Go
go version

# Verificar Node.js
node --version
npm --version
```

## Ejecución

### Opción A: Dos terminales (Recomendado)

#### Terminal 1: API Go

```bash
cd go-api
go run main.go
```

Deberías ver:
```
Go API server starting on port 8080
```

#### Terminal 2: API Node.js

```bash
cd node-api
npm install
npm start
```

Deberías ver:
```
Node.js API server running on port 3000
```

### Opción B: Una terminal con ejecución en background

#### Windows PowerShell

```powershell
# Terminal 1 - API Go
cd go-api
go run main.go

# En otra ventana PowerShell
cd node-api
npm install
npm start
```

#### Linux/macOS

```bash
# Terminal 1 - API Go
cd go-api
go run main.go &

# Terminal 2 - API Node.js
cd node-api
npm install
npm start
```

## Variables de Entorno

### API Go

Crea un archivo `.env` en `go-api/`:

```
PORT=8080
NODE_API_URL=http://localhost:3000
```

### API Node.js

Crea un archivo `.env` en `node-api/`:

```
PORT=3000
NODE_ENV=development
```

## Pruebas

### Pruebas Unitarias Go

```bash
cd go-api
go test -v
```

Salida esperada:
```
=== RUN   TestRotateMatrix
=== RUN   TestRotateMatrix/3x3_matrix
--- PASS: TestRotateMatrix/3x3_matrix (0.00s)
=== RUN   TestRotateMatrix/2x3_matrix
--- PASS: TestRotateMatrix/2x3_matrix (0.00s)
...
PASS
ok      matrix-rotation-api     0.005s
```

### Pruebas Unitarias Node.js

```bash
cd node-api
node test.js
```

Salida esperada:
```
Running tests...

✓ Test 1 passed: Basic 3x3 matrix
✓ Test 2 passed: Diagonal matrix
✓ Test 3 passed: Single element matrix
✓ Test 4 passed: Rectangular matrix
✓ Test 5 passed: Empty matrix
✓ Test 6 passed: Matrix with negative numbers

✅ All tests passed!
```

## Pruebas de Integración

### Con cURL

```bash
# Test simple
curl -X POST http://localhost:8080/api/rotate \
  -H "Content-Type: application/json" \
  -d '{"data": [[1,2,3],[4,5,6],[7,8,9]]}'
```

### Con PowerShell

```powershell
$body = @{
    data = @(
        @(1, 2, 3),
        @(4, 5, 6),
        @(7, 8, 9)
    )
} | ConvertTo-Json

Invoke-WebRequest -Uri "http://localhost:8080/api/rotate" `
  -Method POST `
  -Headers @{"Content-Type"="application/json"} `
  -Body $body
```

## Verificar Estado

### Health Check Go API

```bash
curl http://localhost:8080/health
```

Respuesta esperada:
```json
{"status":"healthy"}
```

### Health Check Node.js API

```bash
curl http://localhost:3000/health
```

Respuesta esperada:
```json
{"status":"healthy"}
```

## Cliente Web

Abre `client.html` en tu navegador:

```bash
# Windows
start client.html

# macOS
open client.html

# Linux
xdg-open client.html
```

O simplemente arrastra el archivo a tu navegador.

## Solución de Problemas

### Error: "Port already in use"

**Problema**: El puerto 8080 o 3000 ya está en uso.

**Soluciones**:

1. Encontrar y matar el proceso:

```bash
# Windows PowerShell
Get-Process | Where-Object {$_.Port -eq 8080}
Stop-Process -Id <PID>

# Linux/macOS
lsof -i :8080
kill -9 <PID>
```

2. Cambiar el puerto en el código o variables de entorno

### Error: "Cannot find module 'express'"

**Problema**: Las dependencias de Node.js no están instaladas.

**Solución**:
```bash
cd node-api
npm install
```

### Error: "go: no Go files in /path/to/go-api"

**Problema**: El archivo `main.go` no existe o está en la ubicación incorrecta.

**Solución**: Verifica que `main.go` esté en la carpeta `go-api/`

### Error: "Connection refused"

**Problema**: Una de las APIs no está corriendo.

**Solución**: Verifica que ambas APIs estén ejecutándose en las terminales correctas.

## Desarrollo

### Editar el código Go

1. Edita `go-api/main.go`
2. Detén el servidor (Ctrl+C)
3. Ejecuta `go run main.go` nuevamente

### Editar el código Node.js

1. Edita `node-api/server.js`
2. El servidor se reiniciará automáticamente si tienes `nodemon` instalado
3. Si no, detén (Ctrl+C) y ejecuta `npm start` nuevamente

### Instalar dependencias adicionales

**Go**:
```bash
cd go-api
go get <package-name>
```

**Node.js**:
```bash
cd node-api
npm install <package-name>
```

## Debugging

### Go

Usa `log.Printf()` para imprimir valores:

```go
log.Printf("Debug: %v", variable)
```

### Node.js

Usa `console.log()` para imprimir valores:

```javascript
console.log('Debug:', variable);
```

## Performance

### Monitorear recursos

**Windows PowerShell**:
```powershell
Get-Process go, node | Select-Object ProcessName, CPU, Memory
```

**Linux/macOS**:
```bash
top -p $(pgrep -f "go run|node")
```

## Próximos Pasos

1. Modifica el código según tus necesidades
2. Ejecuta las pruebas después de cada cambio
3. Consulta `DOCUMENTACION_TECNICA.md` para detalles de implementación
4. Cuando estés listo, usa Docker para desplegar

## Recursos Útiles

- [Go Documentation](https://golang.org/doc/)
- [Express.js Guide](https://expressjs.com/en/starter/basic-routing.html)
- [Node.js Best Practices](https://nodejs.org/en/docs/guides/)
- [Go Testing](https://golang.org/pkg/testing/)
