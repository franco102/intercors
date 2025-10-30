package handlers

import (
	"github.com/franco102/intercors/go-api/internal/services"

	"github.com/gofiber/fiber/v2"
)

type MatrixHandler struct {
	matrixService *services.MatrixService
}

func NewMatrixHandler(matrixService *services.MatrixService) *MatrixHandler {
	return &MatrixHandler{
		matrixService: matrixService,
	}
}

func (h *MatrixHandler) RotateMatrix(c *fiber.Ctx) error {
	// Get and validate the token from Authorization header
	authHeader := c.Get("Authorization")
	if authHeader == "" || len(authHeader) < 8 || authHeader[:7] != "Bearer " {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Missing or invalid Authorization header",
		})
	}
	tokenString := authHeader[7:]

	// Parse request body
	var request struct {
		Data [][]int `json:"data"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON format",
		})
	}

	// Process the matrix rotation with the token
	rotatedMatrix, statistics, err := h.matrixService.RotateMatrix(request.Data, tokenString)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"rotatedMatrix": rotatedMatrix,
		"statistics":    statistics,
	})
}
