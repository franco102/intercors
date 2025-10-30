# Índice de Documentación

Bienvenido al proyecto de Rotación de Matrices. Esta página te ayudará a navegar toda la documentación disponible.

## 🚀 Inicio Rápido

**¿Primer contacto?** Comienza aquí:
1. [README.md](README.md) - Descripción general del proyecto
2. [QUICK_START.md](QUICK_START.md) - Guía de inicio rápido (5 minutos)
3. [client.html](client.html) - Prueba la API desde el navegador

## 📚 Documentación Principal

### Para Usuarios
| Documento | Descripción | Tiempo |
|-----------|-------------|--------|
| [README.md](README.md) | Descripción general, instalación y uso | 10 min |
| [QUICK_START.md](QUICK_START.md) | Guía de inicio rápido | 5 min |
| [PROYECTO_RESUMEN.md](PROYECTO_RESUMEN.md) | Resumen ejecutivo del proyecto | 5 min |

### Para Desarrolladores
| Documento | Descripción | Tiempo |
|-----------|-------------|--------|
| [DESARROLLO_LOCAL.md](DESARROLLO_LOCAL.md) | Desarrollo sin Docker | 15 min |
| [DOCUMENTACION_TECNICA.md](DOCUMENTACION_TECNICA.md) | Detalles técnicos profundos | 30 min |
| [SEGURIDAD_Y_MEJORES_PRACTICAS.md](SEGURIDAD_Y_MEJORES_PRACTICAS.md) | Seguridad y mejores prácticas | 20 min |
| [CONTRIBUIR.md](CONTRIBUIR.md) | Cómo contribuir al proyecto | 10 min |

### Para DevOps
| Documento | Descripción | Tiempo |
|-----------|-------------|--------|
| [DESPLIEGUE.md](DESPLIEGUE.md) | Guía de despliegue en diferentes plataformas | 25 min |
| [docker-compose.yml](docker-compose.yml) | Configuración de Docker Compose | - |
| [.github/workflows/ci.yml](.github/workflows/ci.yml) | Pipeline de CI/CD | - |

### Para Testing
| Documento | Descripción | Tiempo |
|-----------|-------------|--------|
| [POSTMAN_SETUP.md](POSTMAN_SETUP.md) | Configuración de Postman | 10 min |
| [postman_collection.json](postman_collection.json) | Colección de requests | - |
| [examples.sh](examples.sh) | Ejemplos de uso con cURL | - |
| [VERIFICACION.md](VERIFICACION.md) | Verificación del proyecto | 15 min |

### Referencia
| Documento | Descripción |
|-----------|-------------|
| [CHANGELOG.md](CHANGELOG.md) | Historial de cambios |
| [LICENSE](LICENSE) | Licencia MIT |
| [openapi.yaml](openapi.yaml) | Especificación OpenAPI/Swagger |

## 🗂️ Estructura de Directorios

```
Intercors/
├── 📄 Documentación
│   ├── README.md                          # Inicio aquí
│   ├── QUICK_START.md                     # 5 minutos
│   ├── PROYECTO_RESUMEN.md                # Resumen ejecutivo
│   ├── DESARROLLO_LOCAL.md                # Desarrollo
│   ├── DOCUMENTACION_TECNICA.md           # Técnica
│   ├── DESPLIEGUE.md                      # Despliegue
│   ├── SEGURIDAD_Y_MEJORES_PRACTICAS.md  # Seguridad
│   ├── CONTRIBUIR.md                      # Contribución
│   ├── POSTMAN_SETUP.md                   # Testing
│   ├── VERIFICACION.md                    # Verificación
│   ├── CHANGELOG.md                       # Cambios
│   ├── LICENSE                            # Licencia
│   └── INDICE_DOCUMENTACION.md            # Este archivo
│
├── 🐹 Go API
│   ├── main.go                            # Código principal
│   ├── main_test.go                       # Pruebas
│   ├── go.mod                             # Módulo
│   ├── Dockerfile                         # Containerización
│   ├── .env                               # Variables
│   └── .dockerignore                      # Exclusiones
│
├── 📦 Node API
│   ├── server.js                          # Código principal
│   ├── test.js                            # Pruebas
│   ├── package.json                       # Dependencias
│   ├── Dockerfile                         # Containerización
│   ├── nodemon.json                       # Desarrollo
│   ├── .env                               # Variables
│   └── .dockerignore                      # Exclusiones
│
├── 🐳 Docker
│   ├── docker-compose.yml                 # Orquestación
│   └── Makefile                           # Comandos
│
├── 🔧 Configuración
│   ├── .github/workflows/ci.yml           # CI/CD
│   ├── .vscode/settings.json              # Editor
│   ├── .vscode/extensions.json            # Extensiones
│   ├── .vscode/launch.json                # Debugging
│   ├── .vscode/tasks.json                 # Tareas
│   ├── .gitignore                         # Git
│   ├── .env.example                       # Ejemplo
│   └── run.ps1                            # Windows
│
├── 🧪 Testing
│   ├── postman_collection.json            # Postman
│   ├── examples.sh                        # cURL
│   └── POSTMAN_SETUP.md                   # Guía
│
├── 📋 Especificaciones
│   ├── openapi.yaml                       # OpenAPI
│   └── DOCUMENTACION_TECNICA.md           # Técnica
│
└── 🌐 Cliente
    └── client.html                        # Web UI
```

## 🎯 Guías por Caso de Uso

### "Quiero empezar rápido"
1. [QUICK_START.md](QUICK_START.md)
2. [client.html](client.html)
3. Listo para usar

### "Quiero desarrollar localmente"
1. [DESARROLLO_LOCAL.md](DESARROLLO_LOCAL.md)
2. Instala Go y Node.js
3. Ejecuta los servicios
4. Modifica el código

### "Quiero desplegar en producción"
1. [DESPLIEGUE.md](DESPLIEGUE.md)
2. Elige tu plataforma (AWS, GCP, etc.)
3. Sigue los pasos específicos
4. Monitorea los servicios

### "Quiero probar las APIs"
1. [POSTMAN_SETUP.md](POSTMAN_SETUP.md)
2. Importa la colección
3. Ejecuta los requests
4. Revisa los resultados

### "Quiero entender la arquitectura"
1. [README.md](README.md) - Visión general
2. [DOCUMENTACION_TECNICA.md](DOCUMENTACION_TECNICA.md) - Detalles
3. [openapi.yaml](openapi.yaml) - Especificación

### "Quiero contribuir"
1. [CONTRIBUIR.md](CONTRIBUIR.md)
2. Fork el repositorio
3. Crea una rama
4. Envía un Pull Request

### "Quiero asegurar la calidad"
1. [VERIFICACION.md](VERIFICACION.md)
2. Ejecuta el checklist
3. Verifica todos los tests
4. Revisa la documentación

## 📖 Documentación por Rol

### 👨‍💼 Gerente de Proyecto
- [PROYECTO_RESUMEN.md](PROYECTO_RESUMEN.md) - Resumen ejecutivo
- [CHANGELOG.md](CHANGELOG.md) - Historial
- [README.md](README.md) - Visión general

### 👨‍💻 Desarrollador Backend
- [DOCUMENTACION_TECNICA.md](DOCUMENTACION_TECNICA.md) - Arquitectura
- [DESARROLLO_LOCAL.md](DESARROLLO_LOCAL.md) - Setup
- [SEGURIDAD_Y_MEJORES_PRACTICAS.md](SEGURIDAD_Y_MEJORES_PRACTICAS.md) - Calidad

### 🧪 QA/Tester
- [POSTMAN_SETUP.md](POSTMAN_SETUP.md) - Testing
- [VERIFICACION.md](VERIFICACION.md) - Checklist
- [examples.sh](examples.sh) - Ejemplos

### 🚀 DevOps/SRE
- [DESPLIEGUE.md](DESPLIEGUE.md) - Despliegue
- [docker-compose.yml](docker-compose.yml) - Configuración
- [.github/workflows/ci.yml](.github/workflows/ci.yml) - CI/CD

### 📚 Documentalista
- [README.md](README.md) - Principal
- [DOCUMENTACION_TECNICA.md](DOCUMENTACION_TECNICA.md) - Técnica
- [openapi.yaml](openapi.yaml) - Especificación

## 🔍 Búsqueda Rápida

### ¿Cómo...?

**...instalar el proyecto?**
→ [QUICK_START.md](QUICK_START.md) o [README.md](README.md)

**...ejecutar los tests?**
→ [VERIFICACION.md](VERIFICACION.md) o [DESARROLLO_LOCAL.md](DESARROLLO_LOCAL.md)

**...desplegar en AWS?**
→ [DESPLIEGUE.md](DESPLIEGUE.md) - Sección "AWS ECS"

**...usar Postman?**
→ [POSTMAN_SETUP.md](POSTMAN_SETUP.md)

**...contribuir código?**
→ [CONTRIBUIR.md](CONTRIBUIR.md)

**...implementar seguridad?**
→ [SEGURIDAD_Y_MEJORES_PRACTICAS.md](SEGURIDAD_Y_MEJORES_PRACTICAS.md)

**...entender la arquitectura?**
→ [DOCUMENTACION_TECNICA.md](DOCUMENTACION_TECNICA.md)

**...verificar que todo funciona?**
→ [VERIFICACION.md](VERIFICACION.md)

## 📊 Estadísticas de Documentación

| Métrica | Valor |
|---------|-------|
| Archivos de documentación | 15+ |
| Líneas de documentación | 5000+ |
| Ejemplos de código | 50+ |
| Diagramas | 5+ |
| Guías paso a paso | 10+ |
| Casos de uso | 20+ |

## 🌐 Recursos Externos

### Tecnologías
- [Go Documentation](https://golang.org/doc/)
- [Node.js Documentation](https://nodejs.org/docs/)
- [Express.js Guide](https://expressjs.com/)
- [Docker Documentation](https://docs.docker.com/)

### Herramientas
- [Postman Learning Center](https://learning.postman.com/)
- [GitHub Actions Documentation](https://docs.github.com/en/actions)
- [VS Code Documentation](https://code.visualstudio.com/docs)

### Estándares
- [OpenAPI Specification](https://swagger.io/specification/)
- [Semantic Versioning](https://semver.org/)
- [Keep a Changelog](https://keepachangelog.com/)
- [Conventional Commits](https://www.conventionalcommits.org/)

## 💡 Tips de Navegación

1. **Usa Ctrl+F** para buscar en los documentos
2. **Sigue los enlaces** entre documentos
3. **Comienza con README.md** si es tu primer contacto
4. **Consulta VERIFICACION.md** si algo no funciona
5. **Lee CONTRIBUIR.md** antes de hacer cambios

## 📞 Soporte

- **Preguntas**: Abre un Issue en GitHub
- **Bugs**: Reporta en GitHub Issues
- **Sugerencias**: Usa GitHub Discussions
- **Contribuciones**: Sigue [CONTRIBUIR.md](CONTRIBUIR.md)

## ✅ Checklist de Lectura

Marca los documentos que has leído:

- [ ] README.md
- [ ] QUICK_START.md
- [ ] PROYECTO_RESUMEN.md
- [ ] DESARROLLO_LOCAL.md
- [ ] DOCUMENTACION_TECNICA.md
- [ ] DESPLIEGUE.md
- [ ] SEGURIDAD_Y_MEJORES_PRACTICAS.md
- [ ] CONTRIBUIR.md
- [ ] POSTMAN_SETUP.md
- [ ] VERIFICACION.md

## 🎓 Ruta de Aprendizaje Recomendada

### Nivel Principiante (1-2 horas)
1. README.md
2. QUICK_START.md
3. client.html (prueba interactiva)

### Nivel Intermedio (3-5 horas)
1. DESARROLLO_LOCAL.md
2. DOCUMENTACION_TECNICA.md
3. POSTMAN_SETUP.md
4. VERIFICACION.md

### Nivel Avanzado (5-10 horas)
1. SEGURIDAD_Y_MEJORES_PRACTICAS.md
2. DESPLIEGUE.md
3. CONTRIBUIR.md
4. Código fuente (main.go, server.js)

---

**Última actualización**: 2025-10-29  
**Versión**: 1.0.0  
**Estado**: ✅ Documentación completa
