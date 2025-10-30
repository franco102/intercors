# 🔧 Arreglar Error de CORS

El error que ves es porque la API Go no tiene CORS habilitado. Aquí te muestro cómo arreglarlo.

## ❌ Problema

```
Access to fetch at 'http://localhost:8080/health' from origin 'http://127.0.0.1:5500' 
has been blocked by CORS policy: No 'Access-Control-Allow-Origin' header is present
```

## ✅ Solución

### Paso 1: Editar `go-api/main.go`

Abre el archivo `go-api/main.go` y busca la función `RotateHandler`.

**Reemplaza esta función:**

```go
// RotateHandler handles the POST request to rotate a matrix
func RotateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var matrix Matrix
	err := json.NewDecoder(r.Body).Decode(&matrix)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Validate matrix
	if len(matrix.Data) == 0 {
		http.Error(w, "Matrix cannot be empty", http.StatusBadRequest)
		return
	}

	// Rotate the matrix
	rotatedMatrix := RotateMatrix(matrix.Data)

	// Send to Node.js API
	statistics, err := SendToNodeAPI(rotatedMatrix)
	if err != nil {
		log.Printf("Error sending to Node.js API: %v", err)
		http.Error(w, "Error processing matrix", http.StatusInternalServerError)
		return
	}

	// Return response with rotated matrix and statistics
	response := map[string]interface{}{
		"rotatedMatrix": rotatedMatrix,
		"statistics":    statistics,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
```

**Por esta función (con CORS):**

```go
// RotateHandler handles the POST request to rotate a matrix
func RotateHandler(w http.ResponseWriter, r *http.Request) {
	// Agregar headers CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Manejar preflight requests
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var matrix Matrix
	err := json.NewDecoder(r.Body).Decode(&matrix)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Validate matrix
	if len(matrix.Data) == 0 {
		http.Error(w, "Matrix cannot be empty", http.StatusBadRequest)
		return
	}

	// Rotate the matrix
	rotatedMatrix := RotateMatrix(matrix.Data)

	// Send to Node.js API
	statistics, err := SendToNodeAPI(rotatedMatrix)
	if err != nil {
		log.Printf("Error sending to Node.js API: %v", err)
		http.Error(w, "Error processing matrix", http.StatusInternalServerError)
		return
	}

	// Return response with rotated matrix and statistics
	response := map[string]interface{}{
		"rotatedMatrix": rotatedMatrix,
		"statistics":    statistics,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
```

### Paso 2: Actualizar HealthHandler

Busca la función `HealthHandler` y reemplázala:

**Antes:**
```go
// HealthHandler handles health check requests
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "healthy"})
}
```

**Después:**
```go
// HealthHandler handles health check requests
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	// Agregar headers CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Manejar preflight requests
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "healthy"})
}
```

### Paso 3: Reiniciar Docker

```bash
# Detener servicios
docker-compose down

# Reconstruir y iniciar
docker-compose up --build
```

### Paso 4: Probar

Abre `client-simple.html` en tu navegador y debería funcionar perfectamente.

---

## 🎯 Resumen de Cambios

| Función | Cambio |
|---------|--------|
| `RotateHandler` | Agregar headers CORS |
| `HealthHandler` | Agregar headers CORS |
| Ambas | Manejar OPTIONS (preflight) |

---

## ✨ Resultado

Después de estos cambios:
- ✅ El cliente web funcionará sin errores CORS
- ✅ Podrás acceder desde cualquier origen
- ✅ Los presets cargarán correctamente
- ✅ Las estadísticas se mostrarán

---

## 🚀 Alternativa: Usar el Cliente Simple

Si no quieres editar el código, simplemente usa:

```
client-simple.html
```

Este cliente ya está optimizado y debería funcionar mejor.

---

## 📝 Notas

- `Access-Control-Allow-Origin: "*"` permite acceso desde cualquier origen
- Las opciones `OPTIONS` son necesarias para preflight requests
- Node.js ya tiene CORS habilitado por defecto

¿Necesitas ayuda con los cambios?
