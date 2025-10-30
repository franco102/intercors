# Verificación del Proyecto

Esta guía te ayudará a verificar que todo está correctamente configurado y funcionando.

## ✅ Checklist de Verificación

### 1. Estructura de Archivos

```bash
# Verificar que todos los archivos existen
ls -la go-api/main.go
ls -la node-api/server.js
ls -la docker-compose.yml
ls -la README.md
```

### 2. Configuración de Go

```bash
cd go-api

# Verificar go.mod
cat go.mod

# Verificar que el código compila
go build -o main

# Ejecutar pruebas
go test -v

# Limpiar
rm main
cd ..
```

**Resultado esperado:**
```
✓ go.mod contiene: module matrix-rotation-api
✓ Compilación exitosa
✓ Todos los tests pasan
```

### 3. Configuración de Node.js

```bash
cd node-api

# Verificar package.json
cat package.json

# Instalar dependencias
npm install

# Ejecutar pruebas
node test.js

# Limpiar
rm -rf node_modules
cd ..
```

**Resultado esperado:**
```
✓ package.json contiene express y cors
✓ Instalación exitosa
✓ Todos los tests pasan (6/6)
```

### 4. Docker

```bash
# Verificar que Docker está instalado
docker --version
docker-compose --version

# Construir imágenes
docker-compose build

# Verificar imágenes
docker images | grep matrix
```

**Resultado esperado:**
```
✓ Docker version 20.10+
✓ Docker Compose version 1.29+
✓ Dos imágenes construidas: go-api y node-api
```

### 5. Ejecución con Docker Compose

```bash
# Iniciar servicios
docker-compose up -d

# Esperar 5 segundos
sleep 5

# Verificar que están corriendo
docker-compose ps

# Verificar logs
docker-compose logs
```

**Resultado esperado:**
```
✓ go-api-service: Up
✓ node-api-service: Up
✓ Logs muestran "Go API server starting on port 8080"
✓ Logs muestran "Node.js API server running on port 3000"
```

### 6. Health Checks

```bash
# Verificar Go API
curl http://localhost:8080/health

# Verificar Node API
curl http://localhost:3000/health
```

**Resultado esperado:**
```json
{"status":"healthy"}
```

### 7. Test de Rotación

```bash
# Enviar matriz a Go API
curl -X POST http://localhost:8080/api/rotate \
  -H "Content-Type: application/json" \
  -d '{"data": [[1,2,3],[4,5,6],[7,8,9]]}'
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

### 8. Detener Servicios

```bash
# Detener servicios
docker-compose down

# Verificar que se detuvieron
docker-compose ps
```

**Resultado esperado:**
```
✓ No hay contenedores corriendo
```

## 🧪 Pruebas Completas

### Prueba 1: Matriz Simple 3×3

```bash
docker-compose up -d
sleep 5

curl -X POST http://localhost:8080/api/rotate \
  -H "Content-Type: application/json" \
  -d '{"data": [[1,2,3],[4,5,6],[7,8,9]]}'
```

**Verificar:**
- ✓ Status code 200
- ✓ rotatedMatrix es [[7,4,1],[8,5,2],[9,6,3]]
- ✓ maxValue es 9
- ✓ isDiagonal es false

### Prueba 2: Matriz Diagonal

```bash
curl -X POST http://localhost:8080/api/rotate \
  -H "Content-Type: application/json" \
  -d '{"data": [[5,0,0],[0,5,0],[0,0,5]]}'
```

**Verificar:**
- ✓ Status code 200
- ✓ isDiagonal es true
- ✓ maxValue es 5

### Prueba 3: Matriz Rectangular

```bash
curl -X POST http://localhost:8080/api/rotate \
  -H "Content-Type: application/json" \
  -d '{"data": [[1,2,3,4],[5,6,7,8]]}'
```

**Verificar:**
- ✓ Status code 200
- ✓ Dimensiones rotadas correctamente (2×4 → 4×2)

### Prueba 4: Casos de Error

```bash
# Matriz vacía
curl -X POST http://localhost:8080/api/rotate \
  -H "Content-Type: application/json" \
  -d '{"data": []}'

# JSON inválido
curl -X POST http://localhost:8080/api/rotate \
  -H "Content-Type: application/json" \
  -d '{ invalid }'
```

**Verificar:**
- ✓ Status code 400
- ✓ Mensaje de error descriptivo

## 📊 Verificación de Documentación

```bash
# Verificar que existen todos los archivos de documentación
ls -la README.md
ls -la QUICK_START.md
ls -la DOCUMENTACION_TECNICA.md
ls -la DESPLIEGUE.md
ls -la SEGURIDAD_Y_MEJORES_PRACTICAS.md
ls -la CONTRIBUIR.md
ls -la CHANGELOG.md
ls -la LICENSE
```

**Verificar:**
- ✓ README.md > 500 líneas
- ✓ DOCUMENTACION_TECNICA.md > 300 líneas
- ✓ Todos los archivos están en formato Markdown
- ✓ LICENSE contiene MIT License

## 🔧 Verificación de Configuración

```bash
# Verificar .gitignore
cat .gitignore

# Verificar docker-compose.yml
cat docker-compose.yml

# Verificar Makefile
cat Makefile
```

**Verificar:**
- ✓ .gitignore excluye node_modules y .env
- ✓ docker-compose.yml define ambos servicios
- ✓ Makefile tiene targets útiles

## 📝 Verificación de Código

### Go

```bash
cd go-api

# Verificar formato
go fmt ./...

# Verificar que compila
go build

# Ejecutar tests
go test -v

cd ..
```

**Verificar:**
- ✓ Código bien formateado
- ✓ Compila sin errores
- ✓ Todos los tests pasan

### Node.js

```bash
cd node-api

# Instalar dependencias
npm install

# Ejecutar tests
node test.js

cd ..
```

**Verificar:**
- ✓ Dependencias instaladas
- ✓ Todos los tests pasan (6/6)

## 🌐 Verificación de Cliente Web

```bash
# Abrir client.html en navegador
# Verificar que:
# - Se carga correctamente
# - Los botones de preset funcionan
# - Se puede enviar una matriz
# - Se muestran los resultados
```

**Verificar:**
- ✓ Interfaz se carga sin errores
- ✓ Presets cargan matrices correctas
- ✓ Botón "Rotar Matriz" funciona
- ✓ Resultados se muestran correctamente

## 📦 Verificación de Postman

```bash
# Importar postman_collection.json en Postman
# Ejecutar todas las pruebas
# Verificar que todas pasan
```

**Verificar:**
- ✓ Colección importa correctamente
- ✓ Todos los requests funcionan
- ✓ Respuestas son correctas

## 🚀 Verificación de CI/CD

```bash
# Verificar que el archivo existe
cat .github/workflows/ci.yml

# Verificar que contiene:
# - Test de Go
# - Test de Node.js
# - Build de Docker
# - Lint
```

**Verificar:**
- ✓ Archivo existe
- ✓ Contiene todos los jobs
- ✓ Sintaxis YAML correcta

## 📋 Resumen de Verificación

| Componente | Estado | Notas |
|-----------|--------|-------|
| Estructura | ✅ | Todos los archivos presentes |
| Go API | ✅ | Compila y tests pasan |
| Node API | ✅ | Dependencias OK, tests pasan |
| Docker | ✅ | Imágenes construidas |
| Docker Compose | ✅ | Servicios corren correctamente |
| Documentación | ✅ | Completa y detallada |
| Cliente Web | ✅ | Funcional |
| Postman | ✅ | Colección importable |
| CI/CD | ✅ | Configurado |
| Seguridad | ✅ | Mejores prácticas aplicadas |

## ✨ Verificación Final

Si todos los items anteriores están ✅, entonces:

```
✅ El proyecto está completamente funcional
✅ Listo para desarrollo
✅ Listo para despliegue
✅ Documentación completa
✅ Tests implementados
✅ Seguridad considerada
```

## 🐛 Troubleshooting

Si algo no funciona:

1. **Verifica los logs**
   ```bash
   docker-compose logs -f
   ```

2. **Verifica que los puertos están disponibles**
   ```bash
   netstat -an | grep 8080
   netstat -an | grep 3000
   ```

3. **Limpia y reconstruye**
   ```bash
   docker-compose down -v
   docker-compose build --no-cache
   docker-compose up
   ```

4. **Consulta la documentación**
   - README.md
   - DESARROLLO_LOCAL.md
   - DOCUMENTACION_TECNICA.md

## 📞 Soporte

Si encuentras problemas:
1. Revisa los logs
2. Consulta la documentación
3. Abre un Issue en GitHub
4. Contacta a los maintainers

---

**Última actualización**: 2025-10-29  
**Estado**: ✅ Proyecto verificado y funcional
