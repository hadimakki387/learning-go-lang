package middleware

import (
	"new-go-api/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

// JWTProtected func for specify routes group with JWT authentication.
// See: https://github.com/gofiber/contrib/jwt

func JWTProtected() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		token := c.Get("Authorization")
		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "No token found"})
		}
		err := utils.ValidateToken(token, c)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}
		extractedToken, err := utils.ExtractTokenContent(token, c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to extract token content"})
		}
		c.Locals("user", extractedToken["user"])
		return c.Next()
	}
}
