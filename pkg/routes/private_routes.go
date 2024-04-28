package routes

import (
	"new-go-api/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

// PrivateRoutes func for describe group of private routes.
func PrivateRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	// Add middleware for all routes.
	route.Use(middleware.JWTProtected())
	// route.Get("/user", controllers.UserSignUp)
}
