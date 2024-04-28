package controllers

import (
	"fmt"
	"new-go-api/app/models"
	"new-go-api/pkg/utils"
	"new-go-api/platform/database"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type CreateUserStruct struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required"`
}

type SignInStruct struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func findUserByEmail(email string) (*models.User, error) {
	db, err := database.PostgreSQLConnection()
	if err != nil {
		return nil, err
	}
	var user models.User
	_ = db.First(&user, "email = ?", email)

	return &user, nil
}
func ErrorHandler(c *fiber.Ctx, status int, message string) error {
	return c.Status(status).JSON(fiber.Map{"error": message})
}
func UserSignIn(c *fiber.Ctx) error {
	var signInUser SignInStruct
	if err := c.BodyParser(&signInUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	user, err := findUserByEmail(signInUser.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if user == nil || user.ID == uuid.Nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "User not found"})
	}
	checkPassword := utils.ComparePasswords(user.Password, signInUser.Password)
	if !checkPassword {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Wrong password"})
	}

	tokenString, err := utils.GenerateToken(user.ID.String(), time.Now().Add(time.Hour*24))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate token"})
	}

	return c.JSON(fiber.Map{"user": user, "token": tokenString})
}

func CheckAuth(c *fiber.Ctx) error {
	db, err := database.PostgreSQLConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Find all users
	var users []models.User
	result := db.Find(&users)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}

	// Access the user data from the 'users' slice
	fmt.Printf("\nNumber of users: %d\n", result.RowsAffected)
	for _, user := range users {
		fmt.Printf("User ID: %d, Name: %s, Email: %s\n", user.ID, user.Name, user.Email)
	}

	return c.JSON(fiber.Map{"users": users})
}

func CreateUser(c *fiber.Ctx) error {
	var createUserStruct CreateUserStruct
	if err := c.BodyParser(&createUserStruct); err != nil {
		return ErrorHandler(c, fiber.StatusBadRequest, "Invalid request body")
	}

	db, err := database.PostgreSQLConnection()
	if err != nil {
		return ErrorHandler(c, fiber.StatusInternalServerError, err.Error())
	}

	var users []models.User
	result := db.Find(&users)

	user, err := findUserByEmail(createUserStruct.Email)
	if err != nil {
		return ErrorHandler(c, fiber.StatusInternalServerError, err.Error())
	}
	if user != nil && user.ID != uuid.Nil {
		return ErrorHandler(c, fiber.StatusBadRequest, "User already exists")
	}

	hashedPassword := utils.GeneratePassword(createUserStruct.Password)
	newUser := models.User{
		ID:       uuid.New(),
		Email:    createUserStruct.Email,
		Name:     createUserStruct.Name,
		Password: hashedPassword,
	}

	result = db.Create(&newUser)
	if result.Error != nil {
		return ErrorHandler(c, fiber.StatusBadRequest, result.Error.Error())
	}

	tokenString, err := utils.GenerateToken(newUser.ID.String(), time.Now().Add(time.Hour*24))
	if err != nil {
		return ErrorHandler(c, fiber.StatusInternalServerError, "Failed to generate token")
	}

	return c.JSON(fiber.Map{"user": newUser, "token": tokenString})
}
