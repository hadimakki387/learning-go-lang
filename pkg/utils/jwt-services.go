package utils

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func ValidateToken(token string, c *fiber.Ctx) error {
	// Validate the generated token
	validatedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Return the secret key used for signing
		return []byte("secret_key"), nil
	})

	fmt.Println("validatedToken:")
	fmt.Println(validatedToken)

	fmt.Println("err:")
	fmt.Println(err)

	// Check for errors during token validation
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Check if the token is valid
	if !validatedToken.Valid {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Generated token is not valid"})
	}

	return nil
}

func GenerateToken(user string, exp time.Time) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user"] = user
	claims["exp"] = exp.Unix() // Ensure exp is a Unix timestamp

	tokenString, err := token.SignedString([]byte("secret_key"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ExtractTokenContent(tokenString string, c *fiber.Ctx) (jwt.MapClaims, error) {
	if err := ValidateToken(tokenString, c); err != nil {
		// Handle validation error
		return nil, err
	}
	// Create a new parser
	parser := jwt.Parser{}

	// Parse the token without validation
	token, _, err := parser.ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return nil, err
	}

	// Check if the token claims can be converted to MapClaims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("failed to parse token claims")
	}

	return claims, nil
}
