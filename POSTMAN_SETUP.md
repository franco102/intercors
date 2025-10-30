# Configuración de Postman

Esta guía te ayudará a importar y usar la colección de Postman para probar las APIs.

## Instalación de Postman

1. Descarga Postman desde [postman.com](https://www.postman.com/downloads/)
2. Instala la aplicación
3. Crea una cuenta (opcional pero recomendado)

## Importar la Colección

### Método 1: Importar archivo JSON

1. Abre Postman
2. Haz clic en **Import** (esquina superior izquierda)
3. Selecciona **Upload Files**
4. Busca y selecciona `postman_collection.json`
5. Haz clic en **Import**

### Método 2: Importar desde URL

1. Abre Postman
2. Haz clic en **Import**
3. Selecciona **Link**
4. Pega la URL del archivo (si está en un servidor)
5. Haz clic en **Continue** y luego **Import**

## Estructura de la Colección

La colección está organizada en las siguientes carpetas:

### 1. Health Checks
- **Go API Health**: Verifica el estado de la API Go
- **Node.js API Health**: Verifica el estado de la API Node.js

### 2. Matrix Rotation
- **Simple 3x3 Matrix**: Matriz básica 3x3
- **Diagonal Matrix**: Matriz diagonal
- **Rectangular Matrix**: Matriz rectangular
- **Matrix with Negative Numbers**: Matriz con números negativos
- **Large Matrix**: Matriz 5x5

### 3. Statistics
- **Calculate Statistics**: Calcula estadísticas de una matriz

### 4. Error Cases
- **Empty Matrix**: Prueba con matriz vacía
- **Invalid JSON**: Prueba con JSON inválido
- **Missing Data Field**: Prueba sin el campo "data"

## Uso de la Colección

### Ejecutar una solicitud individual

1. Expande la carpeta deseada
2. Haz clic en la solicitud
3. Revisa los detalles en el panel derecho
4. Haz clic en **Send**
5. Revisa la respuesta en la pestaña **Response**

### Ejecutar toda la colección

1. Haz clic en los tres puntos (...) al lado del nombre de la colección
2. Selecciona **Run collection**
3. Configura las opciones:
   - **Iterations**: Número de veces a ejecutar
   - **Delay**: Tiempo entre solicitudes (ms)
4. Haz clic en **Run Matrix Rotation API**

## Variables de Entorno

### Crear un Environment

1. Haz clic en **Environments** (lado izquierdo)
2. Haz clic en **Create New**
3. Nombre: `Local Development`
4. Agrega las siguientes variables:

| Variable | Initial Value | Current Value |
|----------|---------------|---------------|
| go_url | http://localhost:8080 | http://localhost:8080 |
| node_url | http://localhost:3000 | http://localhost:3000 |

5. Haz clic en **Save**

### Usar variables en solicitudes

En lugar de escribir URLs completas, usa:
- `{{go_url}}/api/rotate`
- `{{node_url}}/api/statistics`

## Ejemplos de Pruebas

### Prueba 1: Verificar que ambas APIs están activas

1. Ejecuta **Go API Health**
2. Ejecuta **Node.js API Health**
3. Ambas deben devolver `{"status":"healthy"}`

### Prueba 2: Rotar una matriz simple

1. Abre **Simple 3x3 Matrix**
2. Haz clic en **Send**
3. Verifica que la matriz se rotó correctamente
4. Verifica que las estadísticas se calcularon

### Prueba 3: Probar casos de error

1. Abre **Empty Matrix**
2. Haz clic en **Send**
3. Verifica que devuelve un error 400
4. Repite con **Invalid JSON** y **Missing Data Field**

## Escritura de Tests

Postman permite escribir tests para validar respuestas.

### Ejemplo: Validar que maxValue es 9

1. Abre **Simple 3x3 Matrix**
2. Ve a la pestaña **Tests**
3. Agrega el siguiente código:

```javascript
pm.test("Max value should be 9", function () {
    var jsonData = pm.response.json();
    pm.expect(jsonData.statistics.maxValue).to.equal(9);
});
```

4. Haz clic en **Send**
5. Verifica que el test pasó en la pestaña **Test Results**

### Más ejemplos de tests

```javascript
// Validar que la respuesta es exitosa
pm.test("Status code is 200", function () {
    pm.response.to.have.status(200);
});

// Validar que isDiagonal es false
pm.test("Is not diagonal", function () {
    var jsonData = pm.response.json();
    pm.expect(jsonData.statistics.isDiagonal).to.equal(false);
});

// Validar que totalSum es 45
pm.test("Total sum is 45", function () {
    var jsonData = pm.response.json();
    pm.expect(jsonData.statistics.totalSum).to.equal(45);
});

// Validar estructura de respuesta
pm.test("Response has required fields", function () {
    var jsonData = pm.response.json();
    pm.expect(jsonData).to.have.property('rotatedMatrix');
    pm.expect(jsonData).to.have.property('statistics');
});
```

## Automatización con Newman

Newman es la herramienta CLI de Postman para ejecutar colecciones.

### Instalación

```bash
npm install -g newman
```

### Ejecutar colección

```bash
newman run postman_collection.json
```

### Ejecutar con environment

```bash
newman run postman_collection.json -e environment.json
```

### Generar reporte HTML

```bash
newman run postman_collection.json -r html
```

## Troubleshooting

### Error: "Could not get any response"

**Causa**: Las APIs no están corriendo

**Solución**: Verifica que ambas APIs estén ejecutándose

### Error: "Invalid JSON"

**Causa**: El JSON en el body no es válido

**Solución**: Usa el formateador JSON de Postman (Ctrl+I)

### Error: "Port already in use"

**Causa**: El puerto está siendo usado por otro proceso

**Solución**: Cambia el puerto en las variables de entorno

## Consejos Útiles

1. **Usa Pre-request Scripts** para configurar datos antes de enviar
2. **Usa Tests** para validar respuestas automáticamente
3. **Usa Variables** para evitar repetir URLs
4. **Usa Environments** para cambiar fácilmente entre desarrollo y producción
5. **Exporta resultados** para compartir con el equipo

## Recursos Adicionales

- [Documentación de Postman](https://learning.postman.com/)
- [Postman API Testing](https://learning.postman.com/docs/writing-scripts/test-scripts/)
- [Newman Documentation](https://github.com/postmanlabs/newman)

## Próximos Pasos

1. Importa la colección
2. Configura el environment
3. Ejecuta las pruebas
4. Agrega tus propios tests
5. Automatiza con Newman
