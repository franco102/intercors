# Guía de Contribución

¡Gracias por tu interés en contribuir a este proyecto! Esta guía te ayudará a entender cómo colaborar.

## Código de Conducta

- Sé respetuoso con otros contribuyentes
- Acepta críticas constructivas
- Enfócate en lo mejor para la comunidad
- Reporta comportamiento inapropiado

## Cómo Contribuir

### 1. Reportar Bugs

**Antes de reportar:**
- Verifica que el bug no haya sido reportado
- Intenta reproducir el bug
- Recopila información del sistema

**Al reportar:**
- Usa un título descriptivo
- Describe los pasos exactos para reproducir
- Proporciona ejemplos específicos
- Describe el comportamiento observado vs esperado
- Incluye screenshots si es relevante

**Ejemplo:**
```
Título: Rotación de matriz 2x3 produce resultado incorrecto

Pasos para reproducir:
1. Enviar POST a /api/rotate
2. Con matriz [[1,2,3],[4,5,6]]
3. Observar respuesta

Resultado esperado:
[[4,1],[5,2],[6,3]]

Resultado actual:
[[1,4],[2,5],[3,6]]

Versión: 1.0.0
SO: Windows 10
```

### 2. Sugerir Mejoras

**Proporciona:**
- Descripción clara de la mejora
- Caso de uso
- Beneficios esperados
- Posibles alternativas

### 3. Pull Requests

#### Preparación

1. **Fork el repositorio**
```bash
git clone https://github.com/tu-usuario/Intercors.git
cd Intercors
```

2. **Crear rama**
```bash
git checkout -b feature/tu-feature
# o
git checkout -b fix/tu-bug
```

3. **Hacer cambios**
- Sigue el estilo de código existente
- Escribe código limpio y documentado
- Agrega tests para nuevas funcionalidades

4. **Commit**
```bash
git add .
git commit -m "Descripción clara del cambio"
```

Usa mensajes de commit descriptivos:
```
# ✅ BIEN
git commit -m "Agregar validación de tamaño máximo de matriz"

# ❌ MAL
git commit -m "fix bug"
```

5. **Push**
```bash
git push origin feature/tu-feature
```

6. **Crear Pull Request**
- Ve a GitHub
- Haz clic en "Compare & pull request"
- Completa el template
- Espera revisión

#### Template de Pull Request

```markdown
## Descripción
Descripción breve de los cambios

## Tipo de cambio
- [ ] Bug fix
- [ ] Nueva funcionalidad
- [ ] Mejora de documentación
- [ ] Mejora de performance

## Cambios realizados
- Cambio 1
- Cambio 2
- Cambio 3

## Testing realizado
Describe las pruebas que realizaste

## Checklist
- [ ] Mi código sigue el estilo del proyecto
- [ ] He realizado una auto-revisión
- [ ] He comentado código complejo
- [ ] He actualizado la documentación
- [ ] He agregado tests
- [ ] Los tests pasan localmente
```

## Estándares de Código

### Go

```go
// Nombres en CamelCase
func RotateMatrix(matrix [][]int) [][]int {
    // Implementación
}

// Documentación
// RotateMatrix rotates a matrix 90 degrees clockwise
func RotateMatrix(matrix [][]int) [][]int {
    // ...
}

// Manejo de errores
if err != nil {
    log.Printf("Error: %v", err)
    return nil, err
}

// Formato
go fmt ./...
```

### Node.js

```javascript
// Nombres en camelCase
function calculateStatistics(matrix) {
    // Implementación
}

// Documentación JSDoc
/**
 * Calculates statistics for a matrix
 * @param {number[][]} matrix - 2D array
 * @returns {Object} Statistics
 */
function calculateStatistics(matrix) {
    // ...
}

// Manejo de errores
try {
    // código
} catch (error) {
    console.error('Error:', error);
    throw error;
}

// Formato
npx prettier --write .
```

## Proceso de Revisión

1. **Verificación automática**
   - Tests deben pasar
   - Linters deben pasar
   - Coverage debe ser > 80%

2. **Revisión manual**
   - Mínimo 1 aprobación requerida
   - Cambios solicitados deben ser abordados
   - Conversación constructiva

3. **Merge**
   - Squash commits si es necesario
   - Merge a main/develop

## Desarrollo Local

### Configurar entorno

```bash
# Clonar
git clone https://github.com/tu-usuario/Intercors.git
cd Intercors

# Go
cd go-api
go mod download

# Node.js
cd ../node-api
npm install
```

### Ejecutar tests

```bash
# Go
cd go-api
go test -v

# Node.js
cd ../node-api
node test.js
```

### Ejecutar linters

```bash
# Go
cd go-api
golangci-lint run

# Node.js
cd ../node-api
npx eslint .
```

## Documentación

### Actualizar README

Si cambias funcionalidad, actualiza el README:

```markdown
## Nueva Funcionalidad

Descripción de la funcionalidad

### Uso

```bash
curl -X POST http://localhost:8080/api/nueva
```
```

### Actualizar DOCUMENTACION_TECNICA.md

Para cambios técnicos significativos:

```markdown
## Nueva Sección

Explicación técnica del cambio
```

## Versioning

Seguimos [Semantic Versioning](https://semver.org/):

- **MAJOR**: Cambios incompatibles (1.0.0)
- **MINOR**: Nuevas funcionalidades compatibles (1.1.0)
- **PATCH**: Bug fixes (1.1.1)

## Licencia

Al contribuir, aceptas que tu código esté bajo la licencia MIT.

## Preguntas

- Abre un Issue para preguntas
- Usa Discussions para ideas
- Contacta a los maintainers

## Reconocimiento

Los contribuyentes serán reconocidos en:
- README.md
- CONTRIBUTORS.md
- Releases

## Proceso de Merge

1. Todos los tests deben pasar
2. Mínimo 1 aprobación
3. Sin conflictos de merge
4. Squash commits (opcional)
5. Merge a develop
6. Merge a main en releases

## Después del Merge

- Tu rama será eliminada
- Tu cambio aparecerá en la próxima release
- Serás reconocido como contribuyente

## Recursos Útiles

- [Git Documentation](https://git-scm.com/doc)
- [GitHub Flow](https://guides.github.com/introduction/flow/)
- [Conventional Commits](https://www.conventionalcommits.org/)
- [Keep a Changelog](https://keepachangelog.com/)

## Gracias

¡Gracias por contribuir a hacer este proyecto mejor! 🎉
