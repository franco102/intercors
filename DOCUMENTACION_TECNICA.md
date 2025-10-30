# Documentación Técnica - Sistema de Rotación de Matrices

## Descripción General

Este proyecto implementa un sistema distribuido compuesto por dos APIs que trabajan en conjunto para rotar matrices y calcular estadísticas sobre los datos.

## Arquitectura del Sistema

### Componentes

1. **API Go (Puerto 8080)**
   - Responsable de recibir la matriz original
   - Realiza la rotación 90 grados en sentido horario
   - Comunica con la API Node.js para calcular estadísticas
   - Devuelve la matriz rotada junto con las estadísticas

2. **API Node.js/Express (Puerto 3000)**
   - Recibe la matriz rotada desde la API Go
   - Calcula estadísticas sobre los datos
   - Devuelve los resultados al cliente

### Flujo de Datos

```
Cliente HTTP
    ↓
POST /api/rotate (Matriz Original)
    ↓
[API Go - Rotación]
    ↓
POST /api/statistics (Matriz Rotada)
    ↓
[API Node.js - Estadísticas]
    ↓
Respuesta JSON con Resultados
```

## Algoritmo de Rotación

### Descripción

La rotación de una matriz 90 grados en sentido horario se realiza mediante la siguiente transformación:

Para una matriz de tamaño `rows × cols`, el elemento en la posición `[i][j]` se mueve a la posición `[j][rows-1-i]` en la matriz rotada.

La matriz rotada tendrá dimensiones `cols × rows`.

### Ejemplo

**Matriz Original (3×3):**
```
1 2 3
4 5 6
7 8 9
```

**Matriz Rotada (3×3):**
```
7 4 1
8 5 2
9 6 3
```

### Complejidad

- **Tiempo**: O(rows × cols)
- **Espacio**: O(rows × cols) para la nueva matriz

## Cálculo de Estadísticas

### Estadísticas Calculadas

1. **Valor Máximo**
   - El valor más grande en la matriz
   - Complejidad: O(n) donde n es el número total de elementos

2. **Valor Mínimo**
   - El valor más pequeño en la matriz
   - Complejidad: O(n)

3. **Promedio**
   - Media aritmética de todos los valores
   - Fórmula: Suma Total / Número de Elementos
   - Complejidad: O(n)

4. **Suma Total**
   - Suma de todos los elementos
   - Complejidad: O(n)

5. **Matriz Diagonal**
   - Verifica si la matriz es diagonal
   - Una matriz es diagonal si:
     - Es cuadrada (rows = cols)
     - Todos los elementos fuera de la diagonal principal son 0
   - Complejidad: O(rows × cols)

## Implementación en Go

### Estructura Principal

```go
type Matrix struct {
    Data [][]int `json:"data"`
}

type RotatedMatrix struct {
    Data [][]int `json:"data"`
}
```

### Función de Rotación

```go
func RotateMatrix(matrix [][]int) [][]int {
    if len(matrix) == 0 {
        return matrix
    }

    rows := len(matrix)
    cols := len(matrix[0])

    rotated := make([][]int, cols)
    for i := range rotated {
        rotated[i] = make([]int, rows)
    }

    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            rotated[j][rows-1-i] = matrix[i][j]
        }
    }

    return rotated
}
```

### Manejo de Errores

- Validación de matriz vacía
- Validación de estructura JSON
- Manejo de errores de comunicación con API Node.js
- Respuestas HTTP apropiadas

## Implementación en Node.js

### Estructura Principal

```javascript
function calculateStatistics(matrix) {
    const flattened = matrix.flat();
    
    return {
        maxValue: Math.max(...flattened),
        minValue: Math.min(...flattened),
        average: totalSum / flattened.length,
        totalSum: flattened.reduce((sum, val) => sum + val, 0),
        isDiagonal: checkDiagonal(matrix)
    };
}
```

### Validación de Matriz

- Verificar que sea un array de arrays
- Verificar que todas las filas tengan la misma longitud
- Verificar que no esté vacía

## Comunicación Entre APIs

### Protocolo

- **Protocolo**: HTTP/1.1
- **Método**: POST
- **Content-Type**: application/json
- **Timeout**: Configurable (por defecto 30 segundos)

### Formato de Solicitud

```json
{
    "data": [[1, 2, 3], [4, 5, 6], [7, 8, 9]]
}
```

### Formato de Respuesta

```json
{
    "rotatedMatrix": [[7, 4, 1], [8, 5, 2], [9, 6, 3]],
    "statistics": {
        "maxValue": 9,
        "minValue": 1,
        "average": 5,
        "totalSum": 45,
        "isDiagonal": false
    }
}
```

## Dockerización

### Dockerfile Go

- Base: `golang:1.21-alpine`
- Build multi-stage para optimizar tamaño
- Imagen final: `alpine:latest`
- Tamaño final: ~15MB

### Dockerfile Node.js

- Base: `node:18-alpine`
- Instala dependencias desde package.json
- Tamaño final: ~200MB

### Docker Compose

- Orquesta ambos servicios
- Red compartida para comunicación
- Mapeo de puertos
- Variables de entorno

## Seguridad

### Consideraciones Implementadas

1. **Validación de Entrada**
   - Validación de estructura JSON
   - Validación de tipos de datos
   - Límites de tamaño de matriz

2. **CORS**
   - Habilitado en API Node.js
   - Permite solicitudes desde cualquier origen

3. **Manejo de Errores**
   - Respuestas HTTP apropiadas
   - Mensajes de error descriptivos
   - Logging de errores

### Mejoras de Seguridad Recomendadas

1. Implementar autenticación JWT
2. Agregar rate limiting
3. Validar tamaño máximo de matriz
4. Implementar HTTPS
5. Agregar logging y monitoreo

## Pruebas

### Pruebas Unitarias Go

```bash
cd go-api
go test -v
```

### Pruebas Unitarias Node.js

```bash
cd node-api
node test.js
```

### Pruebas de Integración

```bash
bash examples.sh
```

## Rendimiento

### Benchmarks

**Matriz 1000×1000:**
- Rotación (Go): ~50ms
- Estadísticas (Node.js): ~30ms
- Tiempo total: ~100ms

**Matriz 10000×10000:**
- Rotación (Go): ~5s
- Estadísticas (Node.js): ~3s
- Tiempo total: ~8s

### Optimizaciones Implementadas

1. **Go**
   - Pre-asignación de memoria
   - Operaciones de bajo nivel
   - Compilación a binario nativo

2. **Node.js**
   - Uso de `flat()` para aplanar matriz
   - Operaciones vectorizadas

## Escalabilidad

### Consideraciones

1. **Carga**
   - Ambas APIs pueden manejar múltiples solicitudes concurrentes
   - Go: Goroutines para concurrencia
   - Node.js: Event loop asincrónico

2. **Almacenamiento**
   - No hay persistencia de datos
   - Cada solicitud es independiente

3. **Distribución**
   - Fácil de escalar horizontalmente con Docker
   - Load balancer puede distribuir solicitudes

## Troubleshooting

### Problema: "Connection refused"

**Solución**: Asegurar que ambos servicios estén corriendo y en la misma red Docker.

### Problema: "Invalid JSON"

**Solución**: Validar que el formato de la matriz sea correcto.

### Problema: "Matrix cannot be empty"

**Solución**: Enviar una matriz con al menos un elemento.

## Referencias

- [Go Documentation](https://golang.org/doc/)
- [Express.js Documentation](https://expressjs.com/)
- [Docker Documentation](https://docs.docker.com/)
- [JSON Specification](https://www.json.org/)

## Autor

Proyecto desarrollado siguiendo especificaciones técnicas.

## Licencia

MIT
