package handlers

import (
	"os"
	"sync"

	"github.com/franco102/intercors/go-api/internal/models"
	"github.com/franco102/intercors/go-api/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	jwtSecret       []byte
	userCredentials sync.Map
}

func NewAuthHandler() *AuthHandler {
	secret := []byte(os.Getenv("KEY_JWT"))
	return &AuthHandler{
		jwtSecret: secret,
	}
}
func (h *AuthHandler) StoreCredentials(username, password string) {
	h.userCredentials.Store(username, password)
}
func (h *AuthHandler) GetCredentials(username string) (string, bool) {
	if password, ok := h.userCredentials.Load(username); ok {
		return password.(string), true
	}
	return "", false
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var loginReq models.LoginRequest
	if err := c.BodyParser(&loginReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	// Simple authentication (in production, use proper authentication)
	if loginReq.Username == "admin" && loginReq.Password == "password" {
		token, err := utils.GenerateToken(loginReq.Username, h.jwtSecret)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Could not generate token",
			})
		}

		h.StoreCredentials(loginReq.Username, loginReq.Password)
		return c.JSON(models.TokenResponse{Token: token, User: loginReq.Username})

	}

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": "Invalid credentials",
	})
}
