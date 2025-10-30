#!/bin/bash

# Script para probar las APIs

echo "======================================"
echo "Ejemplos de uso de las APIs"
echo "======================================"
echo ""

# Esperar a que los servicios estén listos
echo "Esperando a que los servicios estén listos..."
sleep 5

# Test 1: Health check Go API
echo "1. Health check - Go API"
echo "GET http://localhost:8080/health"
curl -X GET http://localhost:8080/health
echo -e "\n"

# Test 2: Health check Node API
echo "2. Health check - Node.js API"
echo "GET http://localhost:3000/health"
curl -X GET http://localhost:3000/health
echo -e "\n"

# Test 3: Rotate simple 3x3 matrix
echo "3. Rotación de matriz 3x3"
echo "POST http://localhost:8080/api/rotate"
curl -X POST http://localhost:8080/api/rotate \
  -H "Content-Type: application/json" \
  -d '{
    "data": [
      [1, 2, 3],
      [4, 5, 6],
      [7, 8, 9]
    ]
  }'
echo -e "\n\n"

# Test 4: Rotate diagonal matrix
echo "4. Rotación de matriz diagonal"
echo "POST http://localhost:8080/api/rotate"
curl -X POST http://localhost:8080/api/rotate \
  -H "Content-Type: application/json" \
  -d '{
    "data": [
      [5, 0, 0],
      [0, 5, 0],
      [0, 0, 5]
    ]
  }'
echo -e "\n\n"

# Test 5: Rotate rectangular matrix
echo "5. Rotación de matriz rectangular"
echo "POST http://localhost:8080/api/rotate"
curl -X POST http://localhost:8080/api/rotate \
  -H "Content-Type: application/json" \
  -d '{
    "data": [
      [1, 2, 3, 4],
      [5, 6, 7, 8]
    ]
  }'
echo -e "\n\n"

# Test 6: Rotate matrix with negative numbers
echo "6. Rotación de matriz con números negativos"
echo "POST http://localhost:8080/api/rotate"
curl -X POST http://localhost:8080/api/rotate \
  -H "Content-Type: application/json" \
  -d '{
    "data": [
      [-1, -2, -3],
      [-4, -5, -6],
      [-7, -8, -9]
    ]
  }'
echo -e "\n\n"

echo "======================================"
echo "Pruebas completadas"
echo "======================================"
