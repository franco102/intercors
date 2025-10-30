# Guía de Despliegue

Esta guía te ayudará a desplegar el proyecto en diferentes entornos.

## Despliegue Local con Docker

### Requisitos
- Docker instalado
- Docker Compose instalado

### Pasos

1. **Clonar o descargar el proyecto**
```bash
cd /ruta/del/proyecto
```

2. **Construir las imágenes**
```bash
docker-compose build
```

3. **Iniciar los servicios**
```bash
docker-compose up -d
```

4. **Verificar que está corriendo**
```bash
docker-compose ps
```

Deberías ver:
```
NAME                COMMAND             STATUS
go-api-service      "./main"            Up 2 minutes
node-api-service    "npm start"         Up 2 minutes
```

5. **Ver logs**
```bash
docker-compose logs -f
```

6. **Detener servicios**
```bash
docker-compose down
```

## Despliegue en la Nube

### Opción 1: AWS ECS

#### Preparación

1. **Crear repositorio ECR**
```bash
aws ecr create-repository --repository-name matrix-go-api
aws ecr create-repository --repository-name matrix-node-api
```

2. **Autenticarse en ECR**
```bash
aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin <account-id>.dkr.ecr.us-east-1.amazonaws.com
```

3. **Construir y pushear imágenes**
```bash
# Go API
docker build -t matrix-go-api:latest ./go-api
docker tag matrix-go-api:latest <account-id>.dkr.ecr.us-east-1.amazonaws.com/matrix-go-api:latest
docker push <account-id>.dkr.ecr.us-east-1.amazonaws.com/matrix-go-api:latest

# Node API
docker build -t matrix-node-api:latest ./node-api
docker tag matrix-node-api:latest <account-id>.dkr.ecr.us-east-1.amazonaws.com/matrix-node-api:latest
docker push <account-id>.dkr.ecr.us-east-1.amazonaws.com/matrix-node-api:latest
```

#### Crear Task Definition

Crea un archivo `ecs-task-definition.json`:

```json
{
  "family": "matrix-rotation-api",
  "networkMode": "awsvpc",
  "requiresCompatibilities": ["FARGATE"],
  "cpu": "256",
  "memory": "512",
  "containerDefinitions": [
    {
      "name": "go-api",
      "image": "<account-id>.dkr.ecr.us-east-1.amazonaws.com/matrix-go-api:latest",
      "portMappings": [
        {
          "containerPort": 8080,
          "hostPort": 8080,
          "protocol": "tcp"
        }
      ],
      "environment": [
        {
          "name": "NODE_API_URL",
          "value": "http://node-api:3000"
        }
      ]
    },
    {
      "name": "node-api",
      "image": "<account-id>.dkr.ecr.us-east-1.amazonaws.com/matrix-node-api:latest",
      "portMappings": [
        {
          "containerPort": 3000,
          "hostPort": 3000,
          "protocol": "tcp"
        }
      ]
    }
  ]
}
```

#### Registrar Task Definition

```bash
aws ecs register-task-definition --cli-input-json file://ecs-task-definition.json
```

### Opción 2: Google Cloud Run

#### Preparación

1. **Autenticarse en Google Cloud**
```bash
gcloud auth login
gcloud config set project <project-id>
```

2. **Construir y pushear imágenes**
```bash
# Go API
gcloud builds submit --tag gcr.io/<project-id>/matrix-go-api ./go-api
gcloud run deploy matrix-go-api \
  --image gcr.io/<project-id>/matrix-go-api \
  --platform managed \
  --region us-central1 \
  --set-env-vars NODE_API_URL=http://matrix-node-api:3000

# Node API
gcloud builds submit --tag gcr.io/<project-id>/matrix-node-api ./node-api
gcloud run deploy matrix-node-api \
  --image gcr.io/<project-id>/matrix-node-api \
  --platform managed \
  --region us-central1
```

### Opción 3: Heroku

#### Preparación

1. **Instalar Heroku CLI**
```bash
# Windows
choco install heroku-cli

# macOS
brew tap heroku/brew && brew install heroku

# Linux
curl https://cli-assets.heroku.com/install.sh | sh
```

2. **Autenticarse**
```bash
heroku login
```

3. **Crear aplicaciones**
```bash
heroku create matrix-go-api
heroku create matrix-node-api
```

4. **Pushear código**
```bash
# Go API
git subtree push --prefix go-api heroku main

# Node API
git subtree push --prefix node-api heroku main
```

### Opción 4: DigitalOcean App Platform

1. **Conectar repositorio GitHub**
   - Ir a DigitalOcean App Platform
   - Conectar tu repositorio GitHub
   - Seleccionar este proyecto

2. **Configurar app.yaml**

Crea un archivo `app.yaml`:

```yaml
name: matrix-rotation-api
services:
- name: go-api
  github:
    repo: <username>/<repo>
    branch: main
  build_command: cd go-api && go build -o main .
  run_command: ./main
  http_port: 8080
  envs:
  - key: NODE_API_URL
    value: http://node-api:3000

- name: node-api
  github:
    repo: <username>/<repo>
    branch: main
  build_command: cd node-api && npm install
  run_command: npm start
  http_port: 3000
```

3. **Desplegar**
   - Hacer push a GitHub
   - DigitalOcean desplegará automáticamente

## Despliegue con Kubernetes

### Crear manifiestos

**go-api-deployment.yaml**
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: go-api
spec:
  replicas: 2
  selector:
    matchLabels:
      app: go-api
  template:
    metadata:
      labels:
        app: go-api
    spec:
      containers:
      - name: go-api
        image: matrix-go-api:latest
        ports:
        - containerPort: 8080
        env:
        - name: NODE_API_URL
          value: http://node-api:3000
---
apiVersion: v1
kind: Service
metadata:
  name: go-api
spec:
  selector:
    app: go-api
  ports:
  - protocol: TCP
    port: 8080
    targetPort: 8080
  type: LoadBalancer
```

**node-api-deployment.yaml**
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: node-api
spec:
  replicas: 2
  selector:
    matchLabels:
      app: node-api
  template:
    metadata:
      labels:
        app: node-api
    spec:
      containers:
      - name: node-api
        image: matrix-node-api:latest
        ports:
        - containerPort: 3000
---
apiVersion: v1
kind: Service
metadata:
  name: node-api
spec:
  selector:
    app: node-api
  ports:
  - protocol: TCP
    port: 3000
    targetPort: 3000
  type: ClusterIP
```

### Desplegar en Kubernetes

```bash
kubectl apply -f go-api-deployment.yaml
kubectl apply -f node-api-deployment.yaml

# Verificar estado
kubectl get deployments
kubectl get services
kubectl get pods
```

## Monitoreo

### Logs

```bash
# Docker Compose
docker-compose logs -f go-api
docker-compose logs -f node-api

# Kubernetes
kubectl logs -f deployment/go-api
kubectl logs -f deployment/node-api
```

### Health Checks

```bash
# Go API
curl http://localhost:8080/health

# Node API
curl http://localhost:3000/health
```

### Métricas

Implementa Prometheus para monitoreo:

```yaml
# prometheus.yml
global:
  scrape_interval: 15s

scrape_configs:
  - job_name: 'go-api'
    static_configs:
      - targets: ['localhost:8080']
  
  - job_name: 'node-api'
    static_configs:
      - targets: ['localhost:3000']
```

## Escalado

### Docker Compose

```bash
# Escalar Node API a 3 instancias
docker-compose up -d --scale node-api=3
```

### Kubernetes

```bash
# Escalar Go API a 5 réplicas
kubectl scale deployment go-api --replicas=5
```

## Rollback

### Docker

```bash
# Volver a versión anterior
docker-compose down
git checkout <commit-anterior>
docker-compose up -d
```

### Kubernetes

```bash
# Ver historial
kubectl rollout history deployment/go-api

# Hacer rollback
kubectl rollout undo deployment/go-api
```

## Seguridad

### Variables de Entorno Sensibles

Usa secrets en lugar de variables de entorno:

```bash
# Kubernetes
kubectl create secret generic api-secrets \
  --from-literal=NODE_API_URL=http://node-api:3000

# Docker Compose
# Usa archivos .env.production (no commitar a git)
```

### SSL/TLS

Configura certificados SSL:

```bash
# Let's Encrypt con Certbot
certbot certonly --standalone -d tudominio.com
```

### Firewall

Configura reglas de firewall:
- Solo permitir tráfico en puertos 8080 y 3000 desde IPs autorizadas
- Usar VPN para acceso administrativo

## Backup

### Datos

Como el proyecto no persiste datos, no hay necesidad de backup de datos.

### Código

```bash
# Backup del repositorio
git clone --mirror <repo-url> backup.git
```

## Troubleshooting

### Problema: "Connection refused"

**Solución**: Verifica que ambos servicios estén corriendo y en la misma red.

### Problema: "Port already in use"

**Solución**: Cambia los puertos en la configuración o detén otros servicios.

### Problema: "Out of memory"

**Solución**: Aumenta la memoria asignada en docker-compose.yml o Kubernetes.

## Checklist de Despliegue

- [ ] Código testeado localmente
- [ ] Imágenes Docker construidas
- [ ] Variables de entorno configuradas
- [ ] Health checks funcionando
- [ ] Logs monitoreados
- [ ] Backups configurados
- [ ] Seguridad verificada
- [ ] Documentación actualizada

## Próximos Pasos

1. Elige tu plataforma de despliegue
2. Sigue los pasos específicos
3. Verifica que todo funciona
4. Configura monitoreo
5. Establece alertas
