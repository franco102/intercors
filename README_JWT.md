# 🔐 Sistema de Rotación de Matrices - Con Autenticación JWT

> Sistema distribuido con autenticación JWT para proteger las APIs de rotación de matrices y cálculo de estadísticas.

## 📋 Tabla de Contenidos

- [Descripción](#descripción)
- [Características](#características)
- [Requisitos](#requisitos)
- [Instalación](#instalación)
- [Uso](#uso)
- [Autenticación JWT](#autenticación-jwt)
- [Endpoints](#endpoints)
- [Ejemplos](#ejemplos)
- [Documentación](#documentación)

---

## 📝 Descripción

Sistema completo de APIs distribuidas con:
- **Go API (Fiber)**: Rotación de matrices 90 grados
- **Node.js API (Express)**: Cálculo de estadísticas
- **Autenticación JWT**: Protección de endpoints
- **Cliente Web**: Interfaz interactiva con login
- **Docker**: Containerización completa

---

## ✨ Características

### Seguridad
- ✅ **Autenticación JWT** - Tokens con expiración 24h
- ✅ **Endpoints protegidos** - Solo usuarios autenticados
- ✅ **CORS habilitado** - Comunicación segura entre dominios
- ✅ **Validación de entrada** - Prevención de errores

### APIs
- ✅ **Go API (Fiber)** - Rotación de matrices
- ✅ **Node.js API (Express)** - Estadísticas
- ✅ **Health checks** - Monitoreo de servicios
- ✅ **Comunicación inter-APIs** - Integración completa

### Desarrollo
- ✅ **Docker Compose** - Orquestación de servicios
- ✅ **Cliente web** - Interfaz moderna
- ✅ **Scripts de prueba** - Validación automática
- ✅ **Documentación completa** - Guías detalladas

---

## 🔧 Requisitos

- **Docker** (v20.10+)
- **Docker Compose** (v1.29+)
- **Git** (opcional)

O para desarrollo local:
- **Go** (v1.21+)
- **Node.js** (v18+)
- **npm** (v9+)

---

## 🚀 Instalación

### Opción 1: Con Docker (Recomendado)

```bash
# Clonar o descargar el proyecto
cd Intercors

# Compilar y ejecutar
docker-compose up --build

# Esperar a que ambos servicios estén listos
# Go API: http://localhost:8080
# Node API: http://localhost:3000
```

### Opción 2: Desarrollo Local

**Go API:**
```bash
cd go-api
go mod download
go run main.go
# Escucha en http://localhost:8080
```

**Node.js API:**
```bash
cd node-api
npm install
npm start
# Escucha en http://localhost:3000
```

---

## 🔐 Autenticación JWT

### Credenciales por Defecto

```
Username: admin
Password: password
```

### Flujo de Autenticación

```
1. POST /login → Obtener token
2. Guardar token en localStorage
3. Enviar token en Authorization header
4. Servidor valida token
5. Procesa request si es válido
```

### Ejemplo de Login

```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"password"}'
```

**Respuesta:**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

### Usar Token en Requests

```bash
curl -X POST http://localhost:8080/api/rotate \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -d '{"data": [[1,2,3],[4,5,6],[7,8,9]]}'
```

---

## 📡 Endpoints

### Públicos (Sin JWT)

| Método | Endpoint | Descripción |
|--------|----------|-------------|
| GET | `/health` | Verificar estado |
| POST | `/login` | Obtener token JWT |

### Protegidos (Requieren JWT)

| Método | Endpoint | Descripción | API |
|--------|----------|-------------|-----|
| POST | `/api/rotate` | Rotar matriz | Go |
| POST | `/api/statistics` | Calcular estadísticas | Node |

---

## 💡 Ejemplos

### 1. Iniciar Sesión

```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"password"}'
```

### 2. Rotar Matriz (Con JWT)

```bash
TOKEN="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."

curl -X POST http://localhost:8080/api/rotate \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
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

### 3. Usar Cliente Web

```bash
# Servir archivos estáticos (Python)
python -m http.server 8000

# Abrir en navegador
http://localhost:8000/client-jwt.html
```

---

## 🧪 Pruebas

### Con PowerShell (Windows)

```powershell
.\test-jwt.ps1
```

### Con Bash (Linux/Mac)

```bash
chmod +x test-jwt.sh
./test-jwt.sh
```

### Pruebas Manuales

1. **Test Login:**
   ```bash
   curl -X POST http://localhost:8080/login \
     -H "Content-Type: application/json" \
     -d '{"username":"admin","password":"password"}'
   ```

2. **Test Sin Token (Debe fallar):**
   ```bash
   curl -X POST http://localhost:8080/api/rotate \
     -H "Content-Type: application/json" \
     -d '{"data": [[1,2,3],[4,5,6],[7,8,9]]}'
   ```

3. **Test Con Token (Debe funcionar):**
   ```bash
   TOKEN="<tu-token-aqui>"
   curl -X POST http://localhost:8080/api/rotate \
     -H "Content-Type: application/json" \
     -H "Authorization: Bearer $TOKEN" \
     -d '{"data": [[1,2,3],[4,5,6],[7,8,9]]}'
   ```

---

## 📚 Documentación

### Archivos de Documentación

- **JWT_SECURITY.md** - Guía completa de JWT
- **JWT_IMPLEMENTATION_SUMMARY.md** - Detalles técnicos
- **DOCUMENTACION_TECNICA.md** - Arquitectura del sistema
- **DESARROLLO_LOCAL.md** - Desarrollo sin Docker

### Estructura del Proyecto

```
Intercors/
├── go-api/
│   ├── main.go              # API Go con JWT
│   ├── main_test.go         # Pruebas unitarias
│   ├── go.mod               # Módulo Go
│   ├── go.sum               # Dependencias
│   └── Dockerfile           # Imagen Docker
├── node-api/
│   ├── server.js            # API Node con JWT
│   ├── test.js              # Pruebas unitarias
│   ├── package.json         # Dependencias
│   └── Dockerfile           # Imagen Docker
├── client-jwt.html          # Cliente web con JWT
├── test-jwt.sh              # Script de pruebas (Bash)
├── test-jwt.ps1             # Script de pruebas (PowerShell)
├── docker-compose.yml       # Orquestación
├── JWT_SECURITY.md          # Guía JWT
└── README_JWT.md            # Este archivo
```

---

## 🔒 Seguridad en Producción

### Cambiar Credenciales

**Go API** - Edita `go-api/main.go`:
```go
if loginReq.Username == "admin" && loginReq.Password == "password" {
    // Cambiar por tus credenciales
}
```

**Node.js API** - Edita `node-api/server.js`:
```javascript
if (username === 'admin' && password === 'password') {
    // Cambiar por tus credenciales
}
```

### Cambiar JWT Secret

**Go API** - Edita `go-api/main.go`:
```go
var jwtSecret = []byte("your-secret-key-change-in-production")
// Cambiar por una clave segura
```

**Node.js API** - Edita `node-api/server.js`:
```javascript
const JWT_SECRET = process.env.JWT_SECRET || 'your-secret-key-change-in-production';
```

O usa variables de entorno:
```bash
export JWT_SECRET="tu-secret-key-muy-segura"
```

### Mejoras Recomendadas

- [ ] Usar bcrypt para contraseñas
- [ ] Implementar refresh tokens
- [ ] Agregar rate limiting
- [ ] Implementar logging de accesos
- [ ] Usar HTTPS en producción
- [ ] Implementar roles y permisos

---

## 🐛 Solución de Problemas

### Error: "Missing authorization header"
**Causa**: No se envió el header Authorization  
**Solución**: Agrega `Authorization: Bearer <token>`

### Error: "Invalid or expired token"
**Causa**: Token expirado o inválido  
**Solución**: Obtén un nuevo token con `/login`

### Error: "Invalid credentials"
**Causa**: Username o password incorrectos  
**Solución**: Verifica las credenciales (admin/password)

### Error: "Connection refused"
**Causa**: Las APIs no están corriendo  
**Solución**: Ejecuta `docker-compose up --build`

---

## 📞 Soporte

Para más información, consulta:
- `JWT_SECURITY.md` - Guía de autenticación
- `DOCUMENTACION_TECNICA.md` - Detalles técnicos
- `JWT_IMPLEMENTATION_SUMMARY.md` - Resumen de cambios

---

## 📄 Licencia

MIT License - Ver LICENSE para detalles

---

## ✅ Checklist de Verificación

- [ ] Docker Compose está instalado
- [ ] Ejecutaste `docker-compose up --build`
- [ ] Go API responde en http://localhost:8080/health
- [ ] Node API responde en http://localhost:3000/health
- [ ] Puedes obtener token en http://localhost:8080/login
- [ ] Puedes usar token para rotar matriz
- [ ] Cliente web funciona en http://localhost:8000/client-jwt.html

---

**¡Sistema listo para producción con autenticación JWT!** 🔐
