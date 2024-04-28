package routes

import (
	"new-go-api/app/controllers"
	"new-go-api/app/validations"
	"new-go-api/pkg/middleware"
	"new-go-api/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

// PrivateRoutes func for describe group of private routes.
func PrivateRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")
	route.Post("/post", middleware.JWTProtected(), utils.ValidateInput(&validations.CreatePostStruct{}), controllers.CreatePost)
	route.Get("/post/get-by-user", middleware.JWTProtected(), controllers.GetUserWithPost)
	route.Patch("/post/update/:id", middleware.JWTProtected(), utils.ValidateInput(&validations.UpdatePostStruct{}), controllers.UpdatePost)
	route.Delete("/post/delete/:id", middleware.JWTProtected(), controllers.DeletePost)
}
