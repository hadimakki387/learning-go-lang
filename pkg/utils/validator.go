package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// NewValidator func for create a new validator for model fields.
func NewValidator() *validator.Validate {
	// Create a new validator for a Book model.
	validate := validator.New()

	// Custom validation for uuid.UUID fields.
	_ = validate.RegisterValidation("uuid", func(fl validator.FieldLevel) bool {
		field := fl.Field().String()
		if _, err := uuid.Parse(field); err != nil {
			return true
		}
		return false
	})

	return validate
}

// ValidatorErrors func for show validation errors for each invalid fields.
func ValidatorErrors(err error) map[string]string {
	// Define fields map.
	fields := map[string]string{}

	// Make error message for each invalid field.
	for _, err := range err.(validator.ValidationErrors) {
		fields[err.Field()] = err.Error()
	}

	return fields
}

func ValidateInput(schema interface{}) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if err := c.BodyParser(schema); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": true,
				"msg":   "Error decoding request body: " + err.Error(),
			})
		}

		// Create a new validator.
		validate := NewValidator()

		// Validate the data against the provided schema.
		if err := validate.Struct(schema); err != nil {
			// Return a validation error with details.
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": true,
				"msg":   ValidatorErrors(err),
			})
		}

		// Proceed to the next middleware or handler if validation passes.
		return c.Next()
	}
}
