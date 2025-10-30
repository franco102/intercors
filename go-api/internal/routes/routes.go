package routes

import (
	"github.com/franco102/intercors/go-api/internal/handlers"
	"github.com/franco102/intercors/go-api/internal/middlewares"
	"github.com/franco102/intercors/go-api/internal/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func RegisterRoutes(app *fiber.App) {

	nodeAPIService := services.NewNodeAPIService()
	matrixService := services.NewMatrixService(nodeAPIService)
	authHandler := handlers.NewAuthHandler()
	matrixHandler := handlers.NewMatrixHandler(matrixService)
	healthHandler := handlers.NewHealthHandler()

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
	protected.Use(middlewares.JWTMiddleware)
	{
		protected.Post("/rotate", matrixHandler.RotateMatrix)
	}

}
