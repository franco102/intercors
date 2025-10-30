# Changelog

Todos los cambios notables en este proyecto serán documentados en este archivo.

El formato está basado en [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
y este proyecto adhiere a [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0] - 2025-10-29

### Added
- API en Go para rotación de matrices 90 grados
- API en Node.js/Express para cálculo de estadísticas
- Comunicación HTTP entre APIs
- Dockerización de ambos servicios
- Docker Compose para orquestación
- Pruebas unitarias para ambas APIs
- Cliente web HTML interactivo
- Documentación técnica completa
- Guía de desarrollo local
- Guía de despliegue
- Colección de Postman
- Configuración de CI/CD con GitHub Actions
- OpenAPI/Swagger specification
- Mejores prácticas de seguridad
- Guía de contribución

### Features
- Rotación de matriz 90 grados en sentido horario
- Cálculo de valor máximo
- Cálculo de valor mínimo
- Cálculo de promedio
- Cálculo de suma total
- Detección de matriz diagonal
- Health check endpoints
- CORS habilitado
- Validación de entrada
- Manejo de errores robusto

### Technical Details
- Go 1.21
- Node.js 18
- Express.js 4.18.2
- Docker & Docker Compose
- Alpine Linux para imágenes optimizadas

## [Unreleased]

### Planned Features
- [ ] Autenticación JWT
- [ ] Rate limiting
- [ ] Persistencia de datos
- [ ] Caché de resultados
- [ ] Métricas de Prometheus
- [ ] Logging centralizado
- [ ] Soporte para múltiples formatos de entrada
- [ ] API GraphQL
- [ ] WebSocket para actualizaciones en tiempo real
- [ ] Interfaz web mejorada
- [ ] Soporte para matrices complejas
- [ ] Operaciones adicionales de matrices

### Improvements
- [ ] Optimización de performance
- [ ] Reducción de tamaño de imágenes Docker
- [ ] Mejor manejo de errores
- [ ] Cobertura de tests > 90%
- [ ] Documentación en múltiples idiomas

## Notas de Versión

### v1.0.0
Primera versión estable del proyecto. Incluye:
- Funcionalidad completa de rotación y estadísticas
- Documentación exhaustiva
- Tests unitarios
- Configuración de Docker
- Ejemplos de uso

## Cómo Reportar Cambios

Para reportar cambios en futuras versiones:

1. Crea un Issue describiendo el cambio
2. Incluye ejemplos si es posible
3. Espera feedback de los maintainers
4. Crea un Pull Request con la implementación
5. El cambio será documentado en la próxima versión

## Versionado

Este proyecto sigue [Semantic Versioning](https://semver.org/):

- **MAJOR version** cuando realizas cambios incompatibles en la API
- **MINOR version** cuando agregas funcionalidad compatible hacia atrás
- **PATCH version** cuando corriges bugs compatibles hacia atrás

## Compatibilidad

### Go
- Mínimo: Go 1.21
- Recomendado: Go 1.21 o superior

### Node.js
- Mínimo: Node.js 18
- Recomendado: Node.js 18 o superior

### Docker
- Mínimo: Docker 20.10
- Docker Compose: 1.29 o superior

## Soporte

- Para bugs: Abre un Issue
- Para preguntas: Usa Discussions
- Para seguridad: Contacta a los maintainers

## Licencia

Este proyecto está bajo la licencia MIT.

## Reconocimientos

Gracias a todos los contribuyentes que han ayudado a mejorar este proyecto.

---

**Última actualización**: 2025-10-29
