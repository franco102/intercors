# Guía Rápida de Inicio

## Requisitos Previos

- Docker y Docker Compose instalados
- O alternativamente: Go 1.21+ y Node.js 18+

## Opción 1: Con Docker (Recomendado)

### 1. Iniciar los servicios

```bash
docker-compose up --build
```

Espera a ver:
```
go-api-service    | Go API server starting on port 8080
node-api-service  | Node.js API server running on port 3000
```

### 2. Probar la API

En otra terminal:

```bash
# Test simple
curl -X POST http://localhost:8080/api/rotate \
  -H "Content-Type: application/json" \
  -d '{"data": [[1,2,3],[4,5,6],[7,8,9]]}'
```

### 3. Detener los servicios

```bash
docker-compose down
```

## Opción 2: Ejecución Local

### API Go

```bash
cd go-api
go run main.go
```

En otra terminal, configura la variable de entorno:

```bash
# Windows PowerShell
$env:NODE_API_URL = "http://localhost:3000"
```

### API Node.js

```bash
cd node-api
npm install
npm start
```

## Ejemplos de Uso

### Matriz Simple 3×3

```bash
curl -X POST http://localhost:8080/api/rotate \
  -H "Content-Type: application/json" \
  -d '{
    "data": [
      [1, 2, 3],
      [4, 5, 6],
      [7, 8, 9]
    ]
  }'
```

**Respuesta:**
```json
{
  "rotatedMatrix": [
    [7, 4, 1],
    [8, 5, 2],
    [9, 6, 3]
  ],
  "statistics": {
    "maxValue": 9,
    "minValue": 1,
    "average": 5,
    "totalSum": 45,
    "isDiagonal": false
  }
}
```

### Matriz Diagonal

```bash
curl -X POST http://localhost:8080/api/rotate \
  -H "Content-Type: application/json" \
  -d '{
    "data": [
      [5, 0, 0],
      [0, 5, 0],
      [0, 0, 5]
    ]
  }'
```

### Verificar Estado

```bash
# Go API
curl http://localhost:8080/health

# Node.js API
curl http://localhost:3000/health
```

## Comandos Útiles (con Make)

```bash
make build       # Construir imágenes
make up          # Iniciar servicios
make down        # Detener servicios
make logs        # Ver logs
make test        # Ejecutar pruebas
make health-check # Verificar estado
make clean       # Limpiar todo
```

## Estructura de Archivos

```
.
├── go-api/
│   ├── main.go           # API Go
│   ├── main_test.go      # Pruebas
│   ├── go.mod            # Módulo Go
│   └── Dockerfile
├── node-api/
│   ├── server.js         # API Node.js
│   ├── test.js           # Pruebas
│   ├── package.json
│   └── Dockerfile
├── docker-compose.yml    # Orquestación
├── README.md             # Documentación completa
├── DOCUMENTACION_TECNICA.md
├── Makefile
└── examples.sh
```

## Puertos

- **Go API**: 8080
- **Node.js API**: 3000

## Solución de Problemas

### "Connection refused"
- Asegurar que Docker está corriendo
- Esperar 5 segundos después de iniciar los servicios

### "Port already in use"
- Cambiar los puertos en `docker-compose.yml`
- O detener otros servicios que usen esos puertos

### "Invalid JSON"
- Verificar que el JSON esté bien formado
- Usar herramientas como `jq` para validar

## Próximos Pasos

1. Lee `README.md` para documentación completa
2. Lee `DOCUMENTACION_TECNICA.md` para detalles técnicos
3. Ejecuta `bash examples.sh` para ver más ejemplos
4. Modifica el código según tus necesidades

## Soporte

Para más información, consulta:
- `README.md` - Documentación general
- `DOCUMENTACION_TECNICA.md` - Detalles técnicos
- `examples.sh` - Ejemplos de uso
