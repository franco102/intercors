package handlers

import (
	"fmt"

	"github.com/franco102/intercors/go-api/internal/services"

	"github.com/gofiber/fiber/v2"
)

type MatrixHandler struct {
	matrixService *services.MatrixService
	authHandler   *services.AuthHandler
}

func NewMatrixHandler(matrixService *services.MatrixService) *MatrixHandler {
	return &MatrixHandler{
		matrixService: matrixService,
		authHandler:   services.GetAuthHandler(),
	}
}

func (h *MatrixHandler) RotateMatrix(c *fiber.Ctx) error {
	var request struct {
		Data [][]int `json:"data"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON format",
		})
	}
	// authHeader := c.Get("Authorization")
	// tokenString := authHeader[7:]
	username := c.Locals("username").(string)
	password, found := h.authHandler.GetCredentials(username)
	fmt.Println(password)
	fmt.Println(username)
	return c.JSON(fiber.Map{
		"username": username,
		"password": password,
	})
	// if !found {
	// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
	// 		"error": "Invalid credentials",
	// 	})
	// }
	// fmt.Println(username, password)
	// rotatedMatrix, statistics, err := h.matrixService.RotateMatrix(request.Data, username, password)
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"error": err.Error(),
	// 	})
	// }

	// return c.JSON(fiber.Map{
	// 	"rotatedMatrix": rotatedMatrix,
	// 	"statistics":    statistics,
	// })
}
