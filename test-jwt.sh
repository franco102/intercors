#!/bin/bash

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${BLUE}================================${NC}"
echo -e "${BLUE}JWT Authentication Test Suite${NC}"
echo -e "${BLUE}================================${NC}\n"

# API URLs
GO_API="http://localhost:8080"
NODE_API="http://localhost:3000"

# Test 1: Health Check (Public)
echo -e "${YELLOW}Test 1: Health Check (Public)${NC}"
echo "GET $GO_API/health"
curl -s "$GO_API/health" | jq .
echo ""

# Test 2: Login - Go API
echo -e "${YELLOW}Test 2: Login - Go API${NC}"
echo "POST $GO_API/login"
GO_TOKEN=$(curl -s -X POST "$GO_API/login" \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"password"}' | jq -r '.token')

if [ "$GO_TOKEN" != "null" ] && [ -n "$GO_TOKEN" ]; then
  echo -e "${GREEN}✓ Token obtained successfully${NC}"
  echo "Token: ${GO_TOKEN:0:50}..."
else
  echo -e "${RED}✗ Failed to obtain token${NC}"
  exit 1
fi
echo ""

# Test 3: Login - Node API
echo -e "${YELLOW}Test 3: Login - Node API${NC}"
echo "POST $NODE_API/login"
NODE_TOKEN=$(curl -s -X POST "$NODE_API/login" \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"password"}' | jq -r '.token')

if [ "$NODE_TOKEN" != "null" ] && [ -n "$NODE_TOKEN" ]; then
  echo -e "${GREEN}✓ Token obtained successfully${NC}"
  echo "Token: ${NODE_TOKEN:0:50}..."
else
  echo -e "${RED}✗ Failed to obtain token${NC}"
  exit 1
fi
echo ""

# Test 4: Rotate Matrix - Without Token (Should Fail)
echo -e "${YELLOW}Test 4: Rotate Matrix - Without Token (Should Fail)${NC}"
echo "POST $GO_API/api/rotate (no token)"
RESPONSE=$(curl -s -X POST "$GO_API/api/rotate" \
  -H "Content-Type: application/json" \
  -d '{"data": [[1,2,3],[4,5,6],[7,8,9]]}')

if echo "$RESPONSE" | grep -q "error"; then
  echo -e "${GREEN}✓ Correctly rejected request without token${NC}"
  echo "Response: $RESPONSE"
else
  echo -e "${RED}✗ Should have rejected request without token${NC}"
fi
echo ""

# Test 5: Rotate Matrix - With Token (Should Succeed)
echo -e "${YELLOW}Test 5: Rotate Matrix - With Token (Should Succeed)${NC}"
echo "POST $GO_API/api/rotate (with token)"
curl -s -X POST "$GO_API/api/rotate" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $GO_TOKEN" \
  -d '{"data": [[1,2,3],[4,5,6],[7,8,9]]}' | jq .
echo ""

# Test 6: Statistics - Without Token (Should Fail)
echo -e "${YELLOW}Test 6: Statistics - Without Token (Should Fail)${NC}"
echo "POST $NODE_API/api/statistics (no token)"
RESPONSE=$(curl -s -X POST "$NODE_API/api/statistics" \
  -H "Content-Type: application/json" \
  -d '{"data": [[1,2,3],[4,5,6],[7,8,9]]}')

if echo "$RESPONSE" | grep -q "error"; then
  echo -e "${GREEN}✓ Correctly rejected request without token${NC}"
  echo "Response: $RESPONSE"
else
  echo -e "${RED}✗ Should have rejected request without token${NC}"
fi
echo ""

# Test 7: Statistics - With Token (Should Succeed)
echo -e "${YELLOW}Test 7: Statistics - With Token (Should Succeed)${NC}"
echo "POST $NODE_API/api/statistics (with token)"
curl -s -X POST "$NODE_API/api/statistics" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $NODE_TOKEN" \
  -d '{"data": [[1,2,3],[4,5,6],[7,8,9]]}' | jq .
echo ""

# Test 8: Invalid Token (Should Fail)
echo -e "${YELLOW}Test 8: Invalid Token (Should Fail)${NC}"
echo "POST $GO_API/api/rotate (invalid token)"
RESPONSE=$(curl -s -X POST "$GO_API/api/rotate" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer invalid.token.here" \
  -d '{"data": [[1,2,3],[4,5,6],[7,8,9]]}')

if echo "$RESPONSE" | grep -q "error"; then
  echo -e "${GREEN}✓ Correctly rejected invalid token${NC}"
  echo "Response: $RESPONSE"
else
  echo -e "${RED}✗ Should have rejected invalid token${NC}"
fi
echo ""

# Test 9: Wrong Credentials (Should Fail)
echo -e "${YELLOW}Test 9: Wrong Credentials (Should Fail)${NC}"
echo "POST $GO_API/login (wrong password)"
RESPONSE=$(curl -s -X POST "$GO_API/login" \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"wrongpassword"}')

if echo "$RESPONSE" | grep -q "error"; then
  echo -e "${GREEN}✓ Correctly rejected wrong credentials${NC}"
  echo "Response: $RESPONSE"
else
  echo -e "${RED}✗ Should have rejected wrong credentials${NC}"
fi
echo ""

echo -e "${BLUE}================================${NC}"
echo -e "${GREEN}✓ All tests completed!${NC}"
echo -e "${BLUE}================================${NC}"
