package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/golang-jwt/jwt/v5"
)

// Matrix represents a 2D array of integers
type Matrix struct {
	Data [][]int `json:"data"`
}

// RotatedMatrix represents the rotated matrix to send to Node.js API
type RotatedMatrix struct {
	Data              [][]int `json:"data"`
	OriginalDiagonal bool    `json:"originalDiagonal"`
}

// LoginRequest represents login credentials
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// TokenResponse represents the JWT token response
type TokenResponse struct {
	Token string `json:"token"`
}

// Claims represents JWT claims
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// JWT Secret key
var jwtSecret = []byte("your-secret-key-change-in-production")

// CheckDiagonal checks if a matrix is diagonal
func CheckDiagonal(matrix [][]int) bool {
	if len(matrix) == 0 {
		return false
	}

	rows := len(matrix)
	cols := len(matrix[0])

	// A diagonal matrix must be square
	if rows != cols {
		return false
	}

	// Check if all non-diagonal elements are 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if i != j && matrix[i][j] != 0 {
				return false
			}
		}
	}

	return true
}

// RotateMatrix rotates a matrix 90 degrees clockwise
func RotateMatrix(matrix [][]int) [][]int {
	if len(matrix) == 0 {
		return matrix
	}

	rows := len(matrix)
	cols := len(matrix[0])

	// Create rotated matrix with swapped dimensions
	rotated := make([][]int, cols)
	for i := range rotated {
		rotated[i] = make([]int, rows)
	}

	// Perform rotation: element at [i][j] goes to [j][rows-1-i]
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			rotated[j][rows-1-i] = matrix[i][j]
		}
	}

	return rotated
}

// SendToNodeAPI sends the rotated matrix to the Node.js API
func SendToNodeAPI(rotatedMatrix [][]int, originalDiagonal bool) (map[string]interface{}, error) {
	nodeAPIURL := os.Getenv("NODE_API_URL")
	if nodeAPIURL == "" {
		nodeAPIURL = "http://localhost:3000"
	}

	payload := RotatedMatrix{
		Data:              rotatedMatrix,
		OriginalDiagonal: originalDiagonal,
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("error marshaling JSON: %v", err)
	}

	resp, err := http.Post(
		nodeAPIURL+"/api/statistics",
		"application/json",
		bytes.NewBuffer(jsonPayload),
	)
	if err != nil {
		return nil, fmt.Errorf("error sending request to Node.js API: %v", err)
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return result, nil
}

// RotateHandler handles the POST request to rotate a matrix
func RotateHandler(c *fiber.Ctx) error {
	var matrix Matrix
	err := c.BodyParser(&matrix)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON",
		})
	}

	// Validate matrix
	if len(matrix.Data) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Matrix cannot be empty",
		})
	}

	// Check if original matrix is diagonal
	isOriginalDiagonal := CheckDiagonal(matrix.Data)

	// Rotate the matrix
	rotatedMatrix := RotateMatrix(matrix.Data)

	// Send to Node.js API
	statistics, err := SendToNodeAPI(rotatedMatrix, isOriginalDiagonal)
	if err != nil {
		log.Printf("Error sending to Node.js API: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error processing matrix",
		})
	}

	// Return response with rotated matrix and statistics
	response := fiber.Map{
		"rotatedMatrix": rotatedMatrix,
		"statistics":    statistics,
	}

	return c.JSON(response)
}

// GenerateToken generates a JWT token
func GenerateToken(username string) (string, error) {
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// LoginHandler handles login requests
func LoginHandler(c *fiber.Ctx) error {
	var loginReq LoginRequest
	if err := c.BodyParser(&loginReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	// Simple authentication (in production, use proper password hashing)
	if loginReq.Username == "admin" && loginReq.Password == "password" {
		token, err := GenerateToken(loginReq.Username)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Could not generate token",
			})
		}

		return c.JSON(TokenResponse{Token: token})
	}

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": "Invalid credentials",
	})
}

// JWTMiddleware validates JWT tokens
func JWTMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Missing authorization header",
		})
	}

	// Extract token from "Bearer <token>"
	tokenString := authHeader[7:] // Remove "Bearer "
	if len(authHeader) < 7 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid authorization header format",
		})
	}

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid or expired token",
		})
	}

	// Store username in context
	c.Locals("username", claims.Username)
	return c.Next()
}

// HealthHandler handles health check requests
func HealthHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": "healthy",
	})
}

func main() {
	app := fiber.New()

	// Enable CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Content-Type,Authorization",
	}))

	// Public routes
	app.Get("/health", HealthHandler)
	app.Post("/login", LoginHandler)

	// Protected routes (require JWT)
	app.Post("/api/rotate", JWTMiddleware, RotateHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Go API server starting on port %s", port)
	log.Fatal(app.Listen(":" + port))
}
