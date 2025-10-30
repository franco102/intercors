package main

import (
	"log"
	"os"

	"github.com/franco102/intercors/go-api/internal/config"
	"github.com/franco102/intercors/go-api/internal/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {

	config.LoadEnv()

	app := fiber.New()

	routes.RegisterRoutes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Go API server starting on port %s", port)
	log.Fatal(app.Listen(":" + port))
}
