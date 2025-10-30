.PHONY: help build up down logs test clean

help:
	@echo "Comandos disponibles:"
	@echo "  make build       - Construir las imágenes Docker"
	@echo "  make up          - Iniciar los servicios"
	@echo "  make down        - Detener los servicios"
	@echo "  make logs        - Ver logs de los servicios"
	@echo "  make test        - Ejecutar pruebas"
	@echo "  make clean       - Limpiar todo"
	@echo "  make go-test     - Pruebas unitarias de Go"
	@echo "  make node-test   - Pruebas unitarias de Node.js"

build:
	docker-compose build

up:
	docker-compose up -d
	@echo "Servicios iniciados:"
	@echo "  Go API: http://localhost:8080"
	@echo "  Node.js API: http://localhost:3000"

down:
	docker-compose down

logs:
	docker-compose logs -f

logs-go:
	docker-compose logs -f go-api

logs-node:
	docker-compose logs -f node-api

test: go-test node-test

go-test:
	@echo "Ejecutando pruebas de Go..."
	cd go-api && go test -v

node-test:
	@echo "Ejecutando pruebas de Node.js..."
	cd node-api && node test.js

clean:
	docker-compose down -v
	rm -rf go-api/main
	rm -rf node_modules/

restart: down up

status:
	docker-compose ps

shell-go:
	docker-compose exec go-api sh

shell-node:
	docker-compose exec node-api sh

health-check:
	@echo "Verificando Go API..."
	curl -s http://localhost:8080/health | jq .
	@echo "\nVerificando Node.js API..."
	curl -s http://localhost:3000/health | jq .
