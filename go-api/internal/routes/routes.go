package routes

import (
	"github.com/franco102/intercors/go-api/internal/handlers"
	"github.com/franco102/intercors/go-api/internal/middlewares"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	app.Get("/health", handlers.HealthHandler)
	app.Post("/login", handlers.LoginHandler)

	api := app.Group("/api", middlewares.JWTMiddleware)
	api.Post("/rotate", handlers.RotateHandler)
}
