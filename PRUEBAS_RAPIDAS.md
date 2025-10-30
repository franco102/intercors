# 🧪 Pruebas Rápidas de las APIs

Las APIs ya están corriendo en Docker. Aquí hay varias formas de probarlas.

## ✅ Opción 1: Cliente Web (Más Fácil)

**Simplemente abre `client.html` en tu navegador:**

```
c:\Users\F_S_H\OneDrive\Escritorio\franco\Intercors\client.html
```

**Características:**
- ✅ Interfaz visual bonita
- ✅ Presets de matrices (3×3, diagonal, rectangular, negativos)
- ✅ Muestra resultados en tiempo real
- ✅ Verifica estado de las APIs

**Pasos:**
1. Abre `client.html` en tu navegador
2. Haz clic en un preset (ej: "3×3 Simple")
3. Haz clic en "Rotar Matriz"
4. ¡Listo! Verás los resultados

---

## ✅ Opción 2: PowerShell (Windows)

**Prueba 1: Health Check**

```powershell
# Go API
Invoke-WebRequest -Uri "http://localhost:8080/health" -UseBasicParsing

# Node API
Invoke-WebRequest -Uri "http://localhost:3000/health" -UseBasicParsing
```

**Resultado esperado:**
```
StatusCode        : 200
StatusDescription : OK
Content           : {"status":"healthy"}
```

**Prueba 2: Rotar Matriz Simple**

```powershell
$body = @{
    data = @(
        @(1, 2, 3),
        @(4, 5, 6),
        @(7, 8, 9)
    )
} | ConvertTo-Json

$response = Invoke-WebRequest -Uri "http://localhost:8080/api/rotate" `
  -Method POST `
  -Headers @{"Content-Type"="application/json"} `
  -Body $body `
  -UseBasicParsing

$response.Content | ConvertFrom-Json | ConvertTo-Json
```

**Resultado esperado:**
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

**Prueba 3: Matriz Diagonal**

```powershell
$body = @{
    data = @(
        @(5, 0, 0),
        @(0, 5, 0),
        @(0, 0, 5)
    )
} | ConvertTo-Json

$response = Invoke-WebRequest -Uri "http://localhost:8080/api/rotate" `
  -Method POST `
  -Headers @{"Content-Type"="application/json"} `
  -Body $body `
  -UseBasicParsing

$response.Content | ConvertFrom-Json | ConvertTo-Json
```

---

## ✅ Opción 3: cURL (CMD o PowerShell)

**Prueba 1: Health Check**

```bash
curl http://localhost:8080/health
curl http://localhost:3000/health
```

**Prueba 2: Rotar Matriz**

```bash
curl -X POST http://localhost:8080/api/rotate ^
  -H "Content-Type: application/json" ^
  -d "{\"data\": [[1,2,3],[4,5,6],[7,8,9]]}"
```

**Prueba 3: Matriz Rectangular**

```bash
curl -X POST http://localhost:8080/api/rotate ^
  -H "Content-Type: application/json" ^
  -d "{\"data\": [[1,2,3,4],[5,6,7,8]]}"
```

**Prueba 4: Números Negativos**

```bash
curl -X POST http://localhost:8080/api/rotate ^
  -H "Content-Type: application/json" ^
  -d "{\"data\": [[-1,-2,-3],[-4,-5,-6],[-7,-8,-9]]}"
```

---

## ✅ Opción 4: Postman (Recomendado para Testing)

**Paso 1: Importar colección**
1. Abre Postman
2. Haz clic en "Import"
3. Selecciona `postman_collection.json`
4. Haz clic en "Import"

**Paso 2: Ejecutar requests**
- Expande "Health Checks" → Haz clic en "Go API Health" → Send
- Expande "Matrix Rotation" → Haz clic en "Simple 3x3 Matrix" → Send
- Prueba otros requests

**Paso 3: Ver resultados**
- Verás la respuesta en la pestaña "Response"
- Puedes ver el estado, headers, body, etc.

---

## ✅ Opción 5: REST Client Extension (VS Code)

**Paso 1: Instalar extensión**
1. Abre VS Code
2. Ve a Extensions (Ctrl+Shift+X)
3. Busca "REST Client"
4. Instala la extensión por Huachao Mao

**Paso 2: Crear archivo de pruebas**

Crea un archivo `test.http` en la raíz del proyecto:

```http
### Health Check Go API
GET http://localhost:8080/health

### Health Check Node API
GET http://localhost:3000/health

### Rotar Matriz 3x3
POST http://localhost:8080/api/rotate
Content-Type: application/json

{
  "data": [
    [1, 2, 3],
    [4, 5, 6],
    [7, 8, 9]
  ]
}

### Rotar Matriz Diagonal
POST http://localhost:8080/api/rotate
Content-Type: application/json

{
  "data": [
    [5, 0, 0],
    [0, 5, 0],
    [0, 0, 5]
  ]
}

### Rotar Matriz Rectangular
POST http://localhost:8080/api/rotate
Content-Type: application/json

{
  "data": [
    [1, 2, 3, 4],
    [5, 6, 7, 8]
  ]
}

### Números Negativos
POST http://localhost:8080/api/rotate
Content-Type: application/json

{
  "data": [
    [-1, -2, -3],
    [-4, -5, -6],
    [-7, -8, -9]
  ]
}

### Matriz Grande 5x5
POST http://localhost:8080/api/rotate
Content-Type: application/json

{
  "data": [
    [1, 2, 3, 4, 5],
    [6, 7, 8, 9, 10],
    [11, 12, 13, 14, 15],
    [16, 17, 18, 19, 20],
    [21, 22, 23, 24, 25]
  ]
}
```

**Paso 3: Ejecutar**
- Haz clic en "Send Request" encima de cada request
- O presiona Ctrl+Alt+R

---

## 🧪 Casos de Prueba Completos

### Caso 1: Matriz Simple 3×3

**Entrada:**
```json
{
  "data": [
    [1, 2, 3],
    [4, 5, 6],
    [7, 8, 9]
  ]
}
```

**Salida esperada:**
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

### Caso 2: Matriz Diagonal

**Entrada:**
```json
{
  "data": [
    [5, 0, 0],
    [0, 5, 0],
    [0, 0, 5]
  ]
}
```

**Salida esperada:**
```json
{
  "rotatedMatrix": [
    [0, 0, 5],
    [0, 5, 0],
    [5, 0, 0]
  ],
  "statistics": {
    "maxValue": 5,
    "minValue": 0,
    "average": 1.67,
    "totalSum": 15,
    "isDiagonal": false
  }
}
```

**Nota:** Después de rotar, ya no es diagonal porque los 0s se movieron.

### Caso 3: Matriz Rectangular 2×4

**Entrada:**
```json
{
  "data": [
    [1, 2, 3, 4],
    [5, 6, 7, 8]
  ]
}
```

**Salida esperada:**
```json
{
  "rotatedMatrix": [
    [5, 1],
    [6, 2],
    [7, 3],
    [8, 4]
  ],
  "statistics": {
    "maxValue": 8,
    "minValue": 1,
    "average": 4.5,
    "totalSum": 36,
    "isDiagonal": false
  }
}
```

### Caso 4: Números Negativos

**Entrada:**
```json
{
  "data": [
    [-1, -2, -3],
    [-4, -5, -6],
    [-7, -8, -9]
  ]
}
```

**Salida esperada:**
```json
{
  "rotatedMatrix": [
    [-7, -4, -1],
    [-8, -5, -2],
    [-9, -6, -3]
  ],
  "statistics": {
    "maxValue": -1,
    "minValue": -9,
    "average": -5,
    "totalSum": -45,
    "isDiagonal": false
  }
}
```

---

## 🔍 Casos de Error

### Error 1: Matriz Vacía

**Entrada:**
```json
{
  "data": []
}
```

**Respuesta esperada:**
```
Status: 400
{
  "error": "Matrix cannot be empty"
}
```

### Error 2: JSON Inválido

**Entrada:**
```
{ invalid json }
```

**Respuesta esperada:**
```
Status: 400
{
  "error": "Invalid JSON"
}
```

### Error 3: Campo Faltante

**Entrada:**
```json
{
  "matrix": [[1, 2], [3, 4]]
}
```

**Respuesta esperada:**
```
Status: 400
{
  "error": "Invalid input..."
}
```

---

## 📊 Resumen de Endpoints

| Método | URL | Descripción |
|--------|-----|-------------|
| GET | http://localhost:8080/health | Verificar Go API |
| GET | http://localhost:3000/health | Verificar Node API |
| POST | http://localhost:8080/api/rotate | Rotar matriz |
| POST | http://localhost:3000/api/statistics | Calcular estadísticas |

---

## 🎯 Mi Recomendación

**Para pruebas rápidas:** 
→ Usa el **cliente web** (`client.html`)

**Para testing completo:** 
→ Usa **Postman** con la colección

**Para desarrollo:** 
→ Usa **REST Client** en VS Code

**Para automatización:** 
→ Usa **cURL** o **PowerShell**

---

## 💡 Tips

1. **Verifica que Docker está corriendo:**
   ```bash
   docker-compose ps
   ```

2. **Ver logs en tiempo real:**
   ```bash
   docker-compose logs -f
   ```

3. **Reiniciar servicios:**
   ```bash
   docker-compose restart
   ```

4. **Detener servicios:**
   ```bash
   docker-compose down
   ```

5. **Reconstruir imágenes:**
   ```bash
   docker-compose up --build
   ```

---

**¡Ahora puedes probar las APIs de varias formas!** 🚀
