package routes

import (
	"new-go-api/app/controllers"
	"new-go-api/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")
	route.Post("/user/create", utils.ValidateInput(&controllers.CreateUserStruct{}), controllers.CreateUser)
	route.Post("/user/sign-in", utils.ValidateInput(&controllers.SignInStruct{}), controllers.UserSignIn)
}
