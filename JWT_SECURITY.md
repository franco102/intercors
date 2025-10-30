# 🔐 Autenticación JWT - Guía de Uso

## Descripción

Las APIs ahora están protegidas con autenticación JWT (JSON Web Tokens). Esto garantiza que solo usuarios autorizados puedan acceder a los endpoints protegidos.

---

## 📋 Endpoints

### Públicos (Sin autenticación)

| Método | Endpoint | Descripción |
|--------|----------|-------------|
| GET | `/health` | Verificar estado de la API |
| POST | `/login` | Obtener token JWT |

### Protegidos (Requieren JWT)

| Método | Endpoint | Descripción |
|--------|----------|-------------|
| POST | `/api/rotate` | Rotar matriz (Go API) |
| POST | `/api/statistics` | Calcular estadísticas (Node API) |

---

## 🔑 Credenciales por Defecto

```
Username: admin
Password: password
```

⚠️ **IMPORTANTE**: Cambiar estas credenciales en producción.

---

## 🚀 Cómo Usar

### Paso 1: Obtener Token

**Request:**
```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"password"}'
```

**Response:**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwiZXhwIjoxNzMwMzAwMDAwfQ...."
}
```

### Paso 2: Usar Token en Requests

Agrega el token en el header `Authorization` con formato `Bearer <token>`:

**Request:**
```bash
curl -X POST http://localhost:8080/api/rotate \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwiZXhwIjoxNzMwMzAwMDAwfQ...." \
  -d '{"data": [[1,2,3],[4,5,6],[7,8,9]]}'
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

---

## 📊 Flujo de Autenticación

```
1. Cliente → POST /login (username, password)
   ↓
2. Servidor → Valida credenciales
   ↓
3. Servidor → Genera JWT (válido 24 horas)
   ↓
4. Cliente → Recibe token
   ↓
5. Cliente → Envía requests con Authorization: Bearer <token>
   ↓
6. Servidor → Valida token
   ↓
7. Servidor → Procesa request si token es válido
```

---

## 🔒 Características de Seguridad

✅ **Tokens con expiración** - Válidos por 24 horas  
✅ **Firma HMAC-SHA256** - Imposible falsificar  
✅ **CORS habilitado** - Comunicación segura entre dominios  
✅ **Validación de headers** - Rechaza requests sin token  
✅ **Errores descriptivos** - Facilita debugging  

---

## ⚠️ Errores Comunes

### Error 1: Missing Authorization Header
```json
{
  "error": "Missing authorization header"
}
```
**Solución**: Agrega el header `Authorization: Bearer <token>`

### Error 2: Invalid or Expired Token
```json
{
  "error": "Invalid or expired token"
}
```
**Solución**: Obtén un nuevo token con `/login`

### Error 3: Invalid Credentials
```json
{
  "error": "Invalid credentials"
}
```
**Solución**: Verifica username y password

---

## 🧪 Pruebas con Postman

### 1. Crear request de Login

- **Method**: POST
- **URL**: `http://localhost:8080/login`
- **Body** (JSON):
```json
{
  "username": "admin",
  "password": "password"
}
```

### 2. Copiar token de la respuesta

```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

### 3. Crear request de Rotate

- **Method**: POST
- **URL**: `http://localhost:8080/api/rotate`
- **Headers**:
  - `Content-Type: application/json`
  - `Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...`
- **Body** (JSON):
```json
{
  "data": [
    [1, 2, 3],
    [4, 5, 6],
    [7, 8, 9]
  ]
}
```

---

## 🔐 Cambiar Credenciales (Producción)

### En Go API

Edita `go-api/main.go`:

```go
// Cambiar esta línea:
if loginReq.Username == "admin" && loginReq.Password == "password" {

// Por:
if loginReq.Username == "tu-usuario" && loginReq.Password == "tu-contraseña-segura" {
```

### En Node.js API

Edita `node-api/server.js`:

```javascript
// Cambiar esta línea:
if (username === 'admin' && password === 'password') {

// Por:
if (username === 'tu-usuario' && password === 'tu-contraseña-segura') {
```

### Cambiar JWT Secret

**En Go** - Edita `go-api/main.go`:
```go
var jwtSecret = []byte("tu-secret-key-muy-segura-y-larga")
```

**En Node.js** - Edita `node-api/server.js`:
```javascript
const JWT_SECRET = process.env.JWT_SECRET || 'tu-secret-key-muy-segura-y-larga';
```

O usa variables de entorno:
```bash
export JWT_SECRET="tu-secret-key-muy-segura-y-larga"
```

---

## 📝 Estructura del JWT

Un JWT tiene 3 partes separadas por puntos (`.`):

```
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwiZXhwIjoxNzMwMzAwMDAwfQ.signature
│                                      │                                                      │
└─ Header (algoritmo)                  └─ Payload (datos)                                    └─ Signature (firma)
```

**Header:**
```json
{
  "alg": "HS256",
  "typ": "JWT"
}
```

**Payload:**
```json
{
  "username": "admin",
  "iat": 1730216000,
  "exp": 1730302400
}
```

---

## 🚀 Próximos Pasos

1. ✅ Implementar JWT ✓
2. ⏳ Agregar refresh tokens
3. ⏳ Implementar roles y permisos
4. ⏳ Agregar rate limiting
5. ⏳ Implementar logging de accesos

---

## 📚 Referencias

- [JWT.io](https://jwt.io) - Información sobre JWT
- [RFC 7519](https://tools.ietf.org/html/rfc7519) - Especificación JWT
- [OWASP Authentication Cheat Sheet](https://cheatsheetseries.owasp.org/cheatsheets/Authentication_Cheat_Sheet.html)

---

**¡Tus APIs están ahora seguras con autenticación JWT!** 🔐
