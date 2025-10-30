// internal/services/auth_service.go
package services

import "github.com/franco102/intercors/go-api/internal/handlers"

var authHandler *handlers.AuthHandler

func SetAuthHandler(handler *handlers.AuthHandler) {
	authHandler = handler
}

func GetAuthHandler() *handlers.AuthHandler {
	return authHandler
}
