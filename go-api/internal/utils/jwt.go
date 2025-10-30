package utils

import (
	"time"

	"github.com/franco102/intercors/go-api/internal/models"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(username string, secret []byte) (string, error) {
	claims := &models.Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}
