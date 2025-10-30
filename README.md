# Sistema de Rotación de Matrices y Cálculo de Estadísticas

Este proyecto implementa dos APIs que trabajan en conjunto:
- **API en Go**: Recibe una matriz, realiza la rotación 90 grados y envía los datos a la API en Node.js
- **API en Node.js/Express**: Recibe la matriz rotada, calcula estadísticas y devuelve los resultados

## Arquitectura

```
Cliente HTTP
    ↓
[API Go - Puerto 8080]
    ↓ (HTTP)
[API Node.js/Express - Puerto 3000]
    ↓
Respuesta con estadísticas
```

## Requisitos

- Docker
- Docker Compose
- O alternativamente: Go 1.21+ y Node.js 18+

## Instalación y Ejecución

### Opción 1: Con Docker Compose (Recomendado)

```bash
# Desde la raíz del proyecto
docker-compose up --build
```

Esto iniciará ambos servicios:
- Go API: http://localhost:8080
- Node.js API: http://localhost:3000

### Opción 2: Ejecución Local

#### API en Go

```bash
cd go-api
go mod download
go run main.go
```

#### API en Node.js

```bash
cd node-api
npm install
npm start
```

## 🔐 Autenticación

### Obtener Token de Acceso

**Endpoint:** `POST /login`

**Request:**
```json
{
  "username": "admin",
  "password": "password"
}
```

**Response:**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

### Uso del Token

Incluir el token en el header `Authorization` de las peticiones:
```
Authorization: Bearer <token>
```

## Endpoints

### API Go

#### POST /api/rotate (Protegido)
**Requiere autenticación JWT**

**Headers:**
- `Authorization: Bearer <token>`

**Body:**
Recibe una matriz, la rota 90 grados y envía a la API Node.js para calcular estadísticas.

**Request:**
```json
{
  "data": [
    [1, 2, 3],
    [4, 5, 6],
    [7, 8, 9]
  ]
}
```

**Response:**
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

#### GET /health
Verifica el estado de la API Go.

**Response:**
```json
{
  "status": "healthy"
}
```

### API Node.js

#### POST /api/statistics
Recibe una matriz rotada y calcula estadísticas.

**Request:**
```json
{
  "data": [
    [7, 4, 1],
    [8, 5, 2],
    [9, 6, 3]
  ]
}
```

**Response:**
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

#### GET /health
Verifica el estado de la API Node.js.

**Response:**
```json
{
  "status": "healthy"
}
```

## Funcionalidades

### Rotación de Matriz (Go)
- Rota una matriz 90 grados en sentido horario
- Valida que la matriz no esté vacía
- Soporta matrices de cualquier tamaño

### Cálculo de Estadísticas (Node.js)
- **Valor máximo**: El valor más grande en la matriz
- **Valor mínimo**: El valor más pequeño en la matriz
- **Promedio**: La media de todos los valores
- **Suma total**: La suma de todos los valores
- **Matriz diagonal**: Verifica si la matriz es diagonal (solo valores en la diagonal principal)

## Ejemplos de Uso

### Usando cURL

```bash
# Enviar matriz a la API Go
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

### Usando JavaScript/Fetch

```javascript
const matrix = {
  data: [
    [1, 2, 3],
    [4, 5, 6],
    [7, 8, 9]
  ]
};

fetch('http://localhost:8080/api/rotate', {
  method: 'POST',
  headers: {
    'Content-Type': 'application/json'
  },
  body: JSON.stringify(matrix)
})
.then(response => response.json())
.then(data => console.log(data))
.catch(error => console.error('Error:', error));
```

## Estructura del Proyecto

```
.
├── go-api/
│   ├── main.go              # Código principal de la API Go
│   ├── go.mod               # Módulo de Go
│   └── Dockerfile           # Dockerfile para Go
├── node-api/
│   ├── server.js            # Código principal de la API Node.js
│   ├── package.json         # Dependencias de Node.js
│   └── Dockerfile           # Dockerfile para Node.js
├── docker-compose.yml       # Orquestación de servicios
└── README.md                # Este archivo
```

## Consideraciones Técnicas

- **Lenguajes**: Go para la API de rotación, Node.js/Express para estadísticas
- **Comunicación**: HTTP entre APIs
- **Containerización**: Docker para facilitar despliegue
- **Validación**: Ambas APIs validan la estructura de entrada
- **CORS**: Habilitado en la API Node.js para permitir solicitudes desde diferentes orígenes
- **Manejo de errores**: Respuestas HTTP apropiadas con mensajes descriptivos

## Pruebas

### Matriz simple 3x3
```json
{
  "data": [
    [1, 2, 3],
    [4, 5, 6],
    [7, 8, 9]
  ]
}
```

### Matriz diagonal
```json
{
  "data": [
    [5, 0, 0],
    [0, 5, 0],
    [0, 0, 5]
  ]
}
```

### Matriz rectangular
```json
{
  "data": [
    [1, 2, 3, 4],
    [5, 6, 7, 8]
  ]
}
```

## Detención de Servicios

### Con Docker Compose
```bash
docker-compose down
```

### Limpieza completa
```bash
docker-compose down -v
```

## Notas

- Las APIs están diseñadas para trabajar juntas pero también pueden usarse independientemente
- La API Go espera que la API Node.js esté disponible en `http://node-api:3000` (en Docker)
- Para desarrollo local, ajusta la variable de entorno `NODE_API_URL` en la API Go
