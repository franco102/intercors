# 🎉 Bienvenido al Proyecto Intercors

## Sistema de Rotación de Matrices y Cálculo de Estadísticas

```
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│    🔄 ROTACIÓN DE MATRICES + 📊 ESTADÍSTICAS              │
│                                                             │
│    Go API (8080)  ←→  Node.js API (3000)                  │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

## ⚡ Inicio Rápido (5 minutos)

### Opción 1: Con Docker (Recomendado)

```bash
# Clonar o descargar el proyecto
cd Intercors

# Iniciar servicios
docker-compose up --build

# Abrir en navegador
# http://localhost:8080 (Go API)
# http://localhost:3000 (Node API)
```

### Opción 2: Sin Docker

```bash
# Terminal 1 - Go API
cd go-api
go run main.go

# Terminal 2 - Node API
cd node-api
npm install
npm start
```

### Opción 3: Cliente Web

```bash
# Simplemente abre client.html en tu navegador
# ¡No necesitas nada más!
```

## 📚 Documentación

### 🚀 Para Empezar
- **[README.md](README.md)** - Descripción general
- **[QUICK_START.md](QUICK_START.md)** - Guía rápida
- **[INDICE_DOCUMENTACION.md](INDICE_DOCUMENTACION.md)** - Índice completo

### 💻 Para Desarrolladores
- **[DESARROLLO_LOCAL.md](DESARROLLO_LOCAL.md)** - Setup local
- **[DOCUMENTACION_TECNICA.md](DOCUMENTACION_TECNICA.md)** - Detalles técnicos
- **[SEGURIDAD_Y_MEJORES_PRACTICAS.md](SEGURIDAD_Y_MEJORES_PRACTICAS.md)** - Seguridad

### 🚀 Para DevOps
- **[DESPLIEGUE.md](DESPLIEGUE.md)** - Despliegue en la nube
- **[docker-compose.yml](docker-compose.yml)** - Configuración Docker

### 🧪 Para Testing
- **[POSTMAN_SETUP.md](POSTMAN_SETUP.md)** - Usar Postman
- **[VERIFICACION.md](VERIFICACION.md)** - Verificar proyecto
- **[client.html](client.html)** - Cliente web

### 🤝 Para Contribuir
- **[CONTRIBUIR.md](CONTRIBUIR.md)** - Guía de contribución
- **[CHANGELOG.md](CHANGELOG.md)** - Historial

## 🎯 Características

```
✅ Rotación de matriz 90 grados
✅ Cálculo de estadísticas (max, min, promedio, suma)
✅ Detección de matriz diagonal
✅ Validación de entrada robusta
✅ Docker y Docker Compose
✅ Pruebas unitarias (100% cobertura)
✅ Cliente web interactivo
✅ Documentación exhaustiva
✅ CI/CD con GitHub Actions
✅ Mejores prácticas de seguridad
```

## 🏗️ Arquitectura

```
┌──────────────────────────────────────────────────────────┐
│                     Cliente HTTP                         │
└────────────────────────┬─────────────────────────────────┘
                         │
                    POST /api/rotate
                    (Matriz Original)
                         │
                         ▼
        ┌────────────────────────────────┐
        │      API Go (Puerto 8080)      │
        │  • Rotación de matriz          │
        │  • Validación de entrada       │
        │  • Health check                │
        └────────────┬───────────────────┘
                     │
              POST /api/statistics
              (Matriz Rotada)
                     │
                     ▼
        ┌────────────────────────────────┐
        │   API Node.js (Puerto 3000)    │
        │  • Cálculo de estadísticas     │
        │  • Validación de matriz        │
        │  • Health check                │
        └────────────┬───────────────────┘
                     │
                     ▼
        ┌────────────────────────────────┐
        │    Respuesta JSON              │
        │  • Matriz rotada               │
        │  • Estadísticas                │
        └────────────────────────────────┘
```

## 📊 Ejemplo de Uso

### Entrada
```json
{
  "data": [
    [1, 2, 3],
    [4, 5, 6],
    [7, 8, 9]
  ]
}
```

### Salida
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

## 🛠️ Tecnologías

| Componente | Tecnología |
|-----------|-----------|
| API 1 | Go 1.21 |
| API 2 | Node.js 18 + Express.js |
| Containerización | Docker & Docker Compose |
| Testing | Go test + Node.js test |
| CI/CD | GitHub Actions |
| Documentación | Markdown + OpenAPI |

## 📁 Estructura

```
Intercors/
├── go-api/              # API Go
├── node-api/            # API Node.js
├── .github/workflows/   # CI/CD
├── .vscode/             # Configuración VS Code
├── client.html          # Cliente web
├── docker-compose.yml   # Orquestación
├── README.md            # Documentación principal
└── [10+ archivos más]   # Documentación completa
```

## 🚀 Próximos Pasos

### 1. Primero
```bash
# Lee la documentación
cat README.md

# O ve directamente al cliente web
open client.html
```

### 2. Luego
```bash
# Inicia los servicios
docker-compose up --build

# O ejecuta localmente
cd go-api && go run main.go
cd node-api && npm start
```

### 3. Finalmente
```bash
# Prueba la API
curl -X POST http://localhost:8080/api/rotate \
  -H "Content-Type: application/json" \
  -d '{"data": [[1,2,3],[4,5,6],[7,8,9]]}'
```

## 💡 Comandos Útiles

```bash
# Iniciar servicios
docker-compose up -d

# Ver logs
docker-compose logs -f

# Ejecutar pruebas
make test

# Detener servicios
docker-compose down

# Limpiar todo
docker-compose down -v
```

## 🧪 Testing

### Cliente Web
- Abre `client.html` en tu navegador
- Carga un preset
- Haz clic en "Rotar Matriz"
- ¡Listo!

### Postman
1. Importa `postman_collection.json`
2. Ejecuta los requests
3. Revisa los resultados

### cURL
```bash
curl -X POST http://localhost:8080/api/rotate \
  -H "Content-Type: application/json" \
  -d '{"data": [[1,2,3],[4,5,6],[7,8,9]]}'
```

## 📞 Soporte

- **Preguntas**: Abre un Issue
- **Bugs**: Reporta en GitHub
- **Contribuciones**: Lee [CONTRIBUIR.md](CONTRIBUIR.md)
- **Documentación**: Ver [INDICE_DOCUMENTACION.md](INDICE_DOCUMENTACION.md)

## 📋 Checklist Rápido

- [ ] He leído README.md
- [ ] He ejecutado docker-compose up
- [ ] He probado el cliente web
- [ ] He ejecutado los tests
- [ ] He leído la documentación técnica

## 🎓 Aprendizaje

Este proyecto demuestra:
- ✅ Arquitectura de microservicios
- ✅ Comunicación entre APIs
- ✅ Containerización con Docker
- ✅ CI/CD con GitHub Actions
- ✅ Documentación profesional
- ✅ Testing y calidad
- ✅ Mejores prácticas de seguridad

## 📄 Licencia

MIT License - Ver [LICENSE](LICENSE)

## 🎉 ¡Listo para Empezar!

```
┌─────────────────────────────────────────┐
│                                         │
│   ¡Bienvenido al Proyecto Intercors!   │
│                                         │
│   Elige tu camino:                      │
│   1. Cliente Web → client.html          │
│   2. Docker → docker-compose up         │
│   3. Local → DESARROLLO_LOCAL.md        │
│                                         │
│   ¡Que disfrutes!                       │
│                                         │
└─────────────────────────────────────────┘
```

---

**Última actualización**: 2025-10-29  
**Versión**: 1.0.0  
**Estado**: ✅ Listo para usar
