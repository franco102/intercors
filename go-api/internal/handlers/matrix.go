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
	var request struct {
		Data [][]int `json:"data"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON format",
		})
	}
	authHeader := c.Get("Authorization")
	tokenString := authHeader[7:]
	return c.JSON(fiber.Map{
		"tokenString": tokenString,
	})
	// rotatedMatrix, statistics, err := h.matrixService.RotateMatrix(request.Data, tokenString)
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
