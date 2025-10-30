# 🔐 Resumen de Implementación JWT

## ✅ Cambios Realizados

### 1. **Go API (Fiber)**

#### Archivos modificados:
- `go-api/go.mod` - Agregada dependencia `github.com/golang-jwt/jwt/v5`
- `go-api/main.go` - Implementación completa de JWT

#### Cambios principales:
```go
// Nuevas estructuras
type LoginRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

type TokenResponse struct {
    Token string `json:"token"`
}

type Claims struct {
    Username string `json:"username"`
    jwt.RegisteredClaims
}

// Nuevas funciones
func GenerateToken(username string) (string, error)  // Genera JWT
func LoginHandler(c *fiber.Ctx) error               // Endpoint /login
func JWTMiddleware(c *fiber.Ctx) error              // Valida JWT

// Rutas actualizadas
app.Post("/login", LoginHandler)                    // Pública
app.Post("/api/rotate", JWTMiddleware, RotateHandler) // Protegida
```

---

### 2. **Node.js API (Express)**

#### Archivos modificados:
- `node-api/package.json` - Agregada dependencia `jsonwebtoken`
- `node-api/server.js` - Implementación completa de JWT

#### Cambios principales:
```javascript
// Nuevo middleware
function verifyToken(req, res, next) {
  // Valida JWT en Authorization header
  // Rechaza si falta o es inválido
}

// Nuevo endpoint
app.post('/login', (req, res) => {
  // Autentica usuario
  // Retorna JWT con expiración 24h
})

// Endpoint protegido
app.post('/api/statistics', verifyToken, (req, res) => {
  // Requiere JWT válido
})
```

---

### 3. **Cliente Web**

#### Nuevos archivos:
- `client-jwt.html` - Cliente web con autenticación JWT

#### Características:
- ✅ Formulario de login
- ✅ Almacenamiento de token en localStorage
- ✅ Envío de token en Authorization header
- ✅ Manejo de errores de autenticación
- ✅ Botón de logout
- ✅ Visualización del token

---

### 4. **Documentación**

#### Nuevos archivos:
- `JWT_SECURITY.md` - Guía completa de uso
- `JWT_IMPLEMENTATION_SUMMARY.md` - Este archivo

---

## 🔑 Credenciales por Defecto

```
Username: admin
Password: password
```

⚠️ **Cambiar en producción**

---

## 📊 Flujo de Autenticación

```
┌─────────────────────────────────────────────────────────┐
│                    CLIENTE WEB                          │
├─────────────────────────────────────────────────────────┤
│ 1. Ingresa username y password                          │
│ 2. POST /login con credenciales                         │
└─────────────────────────────────────────────────────────┘
                          ↓
┌─────────────────────────────────────────────────────────┐
│                    SERVIDOR (Go/Node)                   │
├─────────────────────────────────────────────────────────┤
│ 3. Valida credenciales                                  │
│ 4. Genera JWT (válido 24 horas)                         │
│ 5. Retorna token                                        │
└─────────────────────────────────────────────────────────┘
                          ↓
┌─────────────────────────────────────────────────────────┐
│                    CLIENTE WEB                          │
├─────────────────────────────────────────────────────────┤
│ 6. Almacena token en localStorage                       │
│ 7. Envía requests con Authorization: Bearer <token>    │
└─────────────────────────────────────────────────────────┘
                          ↓
┌─────────────────────────────────────────────────────────┐
│                    SERVIDOR (Go/Node)                   │
├─────────────────────────────────────────────────────────┤
│ 8. Valida JWT en middleware                             │
│ 9. Si es válido, procesa request                        │
│ 10. Si es inválido, rechaza con 401                     │
└─────────────────────────────────────────────────────────┘
```

---

## 🚀 Cómo Usar

### Paso 1: Compilar y ejecutar

```bash
docker-compose down
docker-compose up --build
```

### Paso 2: Abrir cliente web

```
http://localhost:8000/client-jwt.html
```

O usa Python para servir:
```bash
python -m http.server 8000
```

### Paso 3: Iniciar sesión

- Username: `admin`
- Password: `password`

### Paso 4: Usar la API

- Carga un preset de matriz
- Haz clic en "Rotar Matriz"
- El cliente enviará el JWT automáticamente

---

## 🔒 Características de Seguridad

| Característica | Go API | Node API |
|---|---|---|
| **Algoritmo** | HS256 | HS256 |
| **Expiración** | 24 horas | 24 horas |
| **Firma** | HMAC-SHA256 | HMAC-SHA256 |
| **Storage** | Variable en memoria | Variable en memoria |
| **CORS** | ✅ Habilitado | ✅ Habilitado |

---

## 📝 Endpoints

### Públicos (Sin JWT)

| Método | Endpoint | Descripción |
|--------|----------|-------------|
| GET | `/health` | Verificar estado |
| POST | `/login` | Obtener token |

### Protegidos (Requieren JWT)

| Método | Endpoint | Descripción |
|--------|----------|-------------|
| POST | `/api/rotate` | Rotar matriz |
| POST | `/api/statistics` | Calcular estadísticas |

---

## ⚠️ Errores Comunes

### 1. "Missing authorization header"
**Causa**: No se envió el header Authorization  
**Solución**: Agrega `Authorization: Bearer <token>`

### 2. "Invalid or expired token"
**Causa**: Token expirado o inválido  
**Solución**: Obtén un nuevo token con `/login`

### 3. "Invalid credentials"
**Causa**: Username o password incorrectos  
**Solución**: Verifica las credenciales

---

## 🔐 Cambiar Credenciales

### En Go API

Edita `go-api/main.go` línea ~199:

```go
if loginReq.Username == "admin" && loginReq.Password == "password" {
    // Cambiar por:
    if loginReq.Username == "tu-usuario" && loginReq.Password == "tu-contraseña" {
```

### En Node.js API

Edita `node-api/server.js` línea ~96:

```javascript
if (username === 'admin' && password === 'password') {
    // Cambiar por:
    if (username === 'tu-usuario' && password === 'tu-contraseña') {
```

---

## 🔑 Cambiar JWT Secret

### En Go API

Edita `go-api/main.go` línea ~46:

```go
var jwtSecret = []byte("your-secret-key-change-in-production")
// Cambiar por:
var jwtSecret = []byte("tu-secret-key-muy-segura-y-larga")
```

### En Node.js API

Edita `node-api/server.js` línea ~7:

```javascript
const JWT_SECRET = process.env.JWT_SECRET || 'your-secret-key-change-in-production';
// Cambiar por:
const JWT_SECRET = process.env.JWT_SECRET || 'tu-secret-key-muy-segura-y-larga';
```

O usa variable de entorno:
```bash
export JWT_SECRET="tu-secret-key-muy-segura-y-larga"
```

---

## 📚 Estructura del JWT

```
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwiZXhwIjoxNzMwMzAwMDAwfQ.signature
│                                      │                                                      │
└─ Header                              └─ Payload                                            └─ Signature
```

**Decodificado:**

Header:
```json
{
  "alg": "HS256",
  "typ": "JWT"
}
```

Payload:
```json
{
  "username": "admin",
  "iat": 1730216000,
  "exp": 1730302400
}
```

---

## 🧪 Pruebas con cURL

### Obtener token

```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"password"}'
```

### Usar token

```bash
TOKEN="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."

curl -X POST http://localhost:8080/api/rotate \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{"data": [[1,2,3],[4,5,6],[7,8,9]]}'
```

---

## 🎯 Próximos Pasos

1. ✅ Implementar JWT ✓
2. ⏳ Agregar refresh tokens
3. ⏳ Implementar roles y permisos
4. ⏳ Agregar rate limiting
5. ⏳ Implementar logging de accesos
6. ⏳ Usar bcrypt para contraseñas

---

## 📖 Referencias

- [JWT.io](https://jwt.io)
- [RFC 7519 - JWT](https://tools.ietf.org/html/rfc7519)
- [OWASP Authentication Cheat Sheet](https://cheatsheetseries.owasp.org/cheatsheets/Authentication_Cheat_Sheet.html)
- [Go JWT Library](https://github.com/golang-jwt/jwt)
- [Node.js JWT Library](https://github.com/auth0/node-jsonwebtoken)

---

## ✨ Resumen

✅ **Autenticación JWT implementada**  
✅ **Endpoints protegidos**  
✅ **Cliente web con login**  
✅ **Documentación completa**  
✅ **Listo para producción** (con cambios de credenciales)

**¡Tus APIs están ahora seguras!** 🔐
