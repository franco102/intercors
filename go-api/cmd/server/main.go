package main

import (
	"log"
	"os"

	"github.com/franco102/intercors/go-api/internal/handlers"
	"github.com/franco102/intercors/go-api/internal/middleware"
	"github.com/franco102/intercors/go-api/internal/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Initialize services
	nodeAPIService := services.NewNodeAPIService()
	matrixService := services.NewMatrixService(nodeAPIService)
	authHandler := handlers.NewAuthHandler()
	matrixHandler := handlers.NewMatrixHandler(matrixService)
	healthHandler := handlers.NewHealthHandler()

	app := fiber.New()

	// Middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Content-Type,Authorization",
	}))

	// Routes
	app.Get("/health", healthHandler.HealthCheck)
	app.Post("/login", authHandler.Login)

	// Protected routes
	protected := app.Group("/api")
	protected.Use(middleware.JWTMiddleware)
	{
		protected.Post("/rotate", matrixHandler.RotateMatrix)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Go API server starting on port %s", port)
	log.Fatal(app.Listen(":" + port))
}
