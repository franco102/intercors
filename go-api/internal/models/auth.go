package models

import "github.com/golang-jwt/jwt/v5"

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
