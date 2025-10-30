# Seguridad y Mejores Prácticas

## Seguridad

### 1. Validación de Entrada

**Implementado:**
- Validación de estructura JSON
- Validación de tipos de datos
- Validación de matriz no vacía

**Recomendaciones adicionales:**
```go
// Limitar tamaño de matriz
const MAX_MATRIX_SIZE = 10000

func ValidateMatrixSize(matrix [][]int) error {
    if len(matrix) * len(matrix[0]) > MAX_MATRIX_SIZE {
        return fmt.Errorf("matrix too large")
    }
    return nil
}
```

### 2. Autenticación y Autorización

**Implementar JWT:**

```go
import "github.com/golang-jwt/jwt/v5"

func GenerateToken(userID string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": userID,
        "exp": time.Now().Add(time.Hour * 24).Unix(),
    })
    return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func ValidateToken(tokenString string) error {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return []byte(os.Getenv("JWT_SECRET")), nil
    })
    if err != nil || !token.Valid {
        return fmt.Errorf("invalid token")
    }
    return nil
}
```

### 3. Rate Limiting

**Implementar en Go:**

```go
import "golang.org/x/time/rate"

var limiter = rate.NewLimiter(rate.Limit(100), 10)

func RateLimitMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if !limiter.Allow() {
            http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
            return
        }
        next.ServeHTTP(w, r)
    })
}
```

**Implementar en Node.js:**

```javascript
const rateLimit = require('express-rate-limit');

const limiter = rateLimit({
    windowMs: 15 * 60 * 1000, // 15 minutos
    max: 100, // límite de 100 requests por ventana
    message: 'Too many requests from this IP'
});

app.use(limiter);
```

### 4. HTTPS/TLS

**Configurar en Go:**

```go
func main() {
    http.HandleFunc("/api/rotate", RotateHandler)
    
    // Usar HTTPS
    log.Fatal(http.ListenAndServeTLS(":443", 
        "cert.pem", 
        "key.pem", 
        nil))
}
```

**Configurar en Node.js:**

```javascript
const https = require('https');
const fs = require('fs');
const express = require('express');

const app = express();
const options = {
    key: fs.readFileSync('key.pem'),
    cert: fs.readFileSync('cert.pem')
};

https.createServer(options, app).listen(3000);
```

### 5. CORS Seguro

**Configurar en Node.js:**

```javascript
const cors = require('cors');

const corsOptions = {
    origin: process.env.ALLOWED_ORIGINS?.split(',') || ['http://localhost:3000'],
    credentials: true,
    optionsSuccessStatus: 200,
    methods: ['GET', 'POST'],
    allowedHeaders: ['Content-Type', 'Authorization']
};

app.use(cors(corsOptions));
```

### 6. Protección contra Ataques Comunes

**CSRF Protection:**

```javascript
const csrf = require('csurf');
const cookieParser = require('cookie-parser');

app.use(cookieParser());
app.use(csrf({ cookie: true }));

app.post('/api/rotate', (req, res) => {
    // CSRF token validado automáticamente
    // ...
});
```

**XSS Protection:**

```javascript
const helmet = require('helmet');
app.use(helmet());
```

**SQL Injection (si usas base de datos):**

```go
// Usar prepared statements
stmt, err := db.Prepare("SELECT * FROM users WHERE id = ?")
defer stmt.Close()
rows, err := stmt.Query(userID)
```

### 7. Logging y Auditoría

**Go:**

```go
import "log"

func LogRequest(r *http.Request) {
    log.Printf("[%s] %s %s from %s", 
        time.Now().Format(time.RFC3339),
        r.Method,
        r.URL.Path,
        r.RemoteAddr)
}
```

**Node.js:**

```javascript
const morgan = require('morgan');
app.use(morgan('combined'));

// Log personalizado
app.use((req, res, next) => {
    console.log(`[${new Date().toISOString()}] ${req.method} ${req.path}`);
    next();
});
```

### 8. Manejo de Errores Seguro

**No exponer detalles internos:**

```go
// ❌ MAL
http.Error(w, fmt.Sprintf("Database error: %v", err), 500)

// ✅ BIEN
log.Printf("Internal error: %v", err)
http.Error(w, "Internal server error", 500)
```

## Mejores Prácticas

### 1. Código Limpio

**Go:**
```go
// ✅ BIEN: Nombres descriptivos
func RotateMatrixClockwise(matrix [][]int) [][]int {
    // ...
}

// ❌ MAL: Nombres genéricos
func rotate(m [][]int) [][]int {
    // ...
}
```

**Node.js:**
```javascript
// ✅ BIEN
function calculateStatistics(matrix) {
    const flattened = matrix.flat();
    const maxValue = Math.max(...flattened);
    // ...
}

// ❌ MAL
function calc(m) {
    const f = m.flat();
    const max = Math.max(...f);
    // ...
}
```

### 2. Documentación

**Go:**
```go
// RotateMatrix rotates a matrix 90 degrees clockwise.
// It takes a 2D array of integers and returns the rotated matrix.
// Time complexity: O(rows * cols)
// Space complexity: O(rows * cols)
func RotateMatrix(matrix [][]int) [][]int {
    // ...
}
```

**Node.js:**
```javascript
/**
 * Calculates statistics for a given matrix
 * @param {number[][]} matrix - 2D array of numbers
 * @returns {Object} Statistics object with max, min, average, sum, isDiagonal
 * @throws {Error} If matrix is invalid
 */
function calculateStatistics(matrix) {
    // ...
}
```

### 3. Testing

**Go:**
```go
func TestRotateMatrix(t *testing.T) {
    tests := []struct {
        name     string
        input    [][]int
        expected [][]int
    }{
        {
            name: "3x3 matrix",
            input: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
            expected: [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}},
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := RotateMatrix(tt.input)
            if !reflect.DeepEqual(result, tt.expected) {
                t.Errorf("expected %v, got %v", tt.expected, result)
            }
        })
    }
}
```

### 4. Versionado de API

**Go:**
```go
func main() {
    http.HandleFunc("/api/v1/rotate", RotateHandlerV1)
    http.HandleFunc("/api/v2/rotate", RotateHandlerV2)
}
```

### 5. Configuración

**Usar variables de entorno:**

```go
port := os.Getenv("PORT")
if port == "" {
    port = "8080"
}

nodeAPIURL := os.Getenv("NODE_API_URL")
if nodeAPIURL == "" {
    nodeAPIURL = "http://localhost:3000"
}
```

### 6. Manejo de Dependencias

**Go:**
```bash
# Usar go.mod para versionar dependencias
go get -u github.com/package@v1.2.3
go mod tidy
```

**Node.js:**
```bash
# Usar package-lock.json
npm ci  # En lugar de npm install en producción
npm audit  # Verificar vulnerabilidades
```

### 7. Performance

**Go:**
```go
// Pre-asignar memoria
rotated := make([][]int, cols)
for i := range rotated {
    rotated[i] = make([]int, rows)
}

// Usar buffers
buf := new(bytes.Buffer)
json.NewEncoder(buf).Encode(data)
```

**Node.js:**
```javascript
// Usar streams para datos grandes
fs.createReadStream('large-file.json')
    .pipe(JSONStream.parse('*'))
    .on('data', (data) => {
        // Procesar datos
    });
```

### 8. Monitoreo y Observabilidad

**Implementar métricas:**

```go
import "github.com/prometheus/client_golang/prometheus"

var (
    requestsTotal = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "requests_total",
            Help: "Total number of requests",
        },
        []string{"method", "endpoint"},
    )
)

func init() {
    prometheus.MustRegister(requestsTotal)
}
```

### 9. Estructura del Proyecto

```
project/
├── go-api/
│   ├── main.go
│   ├── handlers/
│   │   ├── rotate.go
│   │   └── health.go
│   ├── models/
│   │   └── matrix.go
│   ├── utils/
│   │   └── validation.go
│   ├── tests/
│   │   └── rotate_test.go
│   └── Dockerfile
├── node-api/
│   ├── server.js
│   ├── routes/
│   │   └── statistics.js
│   ├── controllers/
│   │   └── statisticsController.js
│   ├── middleware/
│   │   └── validation.js
│   ├── tests/
│   │   └── statistics.test.js
│   └── Dockerfile
└── docker-compose.yml
```

### 10. CI/CD

**Implementar en GitHub Actions:**

```yaml
- name: Run tests
  run: |
    cd go-api && go test -v
    cd ../node-api && npm test

- name: Run linters
  run: |
    cd go-api && golangci-lint run
    cd ../node-api && eslint .

- name: Build Docker images
  run: docker-compose build
```

## Checklist de Seguridad

- [ ] Validación de entrada implementada
- [ ] Autenticación configurada
- [ ] Rate limiting habilitado
- [ ] HTTPS/TLS configurado
- [ ] CORS restringido
- [ ] Protección CSRF implementada
- [ ] XSS protection habilitada
- [ ] Logging y auditoría configurados
- [ ] Errores no exponen detalles internos
- [ ] Dependencias actualizadas
- [ ] Tests de seguridad implementados
- [ ] Documentación de seguridad completada

## Recursos

- [OWASP Top 10](https://owasp.org/www-project-top-ten/)
- [Go Security Best Practices](https://golang.org/doc/effective_go)
- [Express.js Security](https://expressjs.com/en/advanced/best-practice-security.html)
- [Node.js Security Best Practices](https://nodejs.org/en/docs/guides/security/)
