package main

import (
	"log"
	"new-go-api/pkg/configs"
	"new-go-api/pkg/routes"
	"new-go-api/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Define Fiber config.
	config := configs.FiberConfig()

	// Define a new Fiber app with config.
	app := fiber.New(config)

	// Routes.
	routes.SwaggerRoute(app)  // Register a route for API Docs (Swagger).
	routes.PublicRoutes(app)  // Register a public routes for app.
	routes.PrivateRoutes(app) // Register a private routes for app.
	routes.NotFoundRoute(app) // Register route for 404 Error.

	// Start server (with or without graceful shutdown).

	utils.StartServer(app)

}
