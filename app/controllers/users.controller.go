package controllers

import (
	"new-go-api/app/models"
	reponsemodels "new-go-api/app/reponse-models"
	"new-go-api/app/validations"
	"new-go-api/pkg/utils"
	"new-go-api/platform/database"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func FindUserByEmail(email string) (*models.User, error) {

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
	var signInUser validations.SignInStruct
	if err := c.BodyParser(&signInUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	user, err := FindUserByEmail(signInUser.Email)
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

	response := reponsemodels.UserResponse{
		ID:           user.ID,
		Name:         user.Name,
		Email:        user.Email,
		Age:          user.Age,
		MemberNumber: user.MemberNumber,
		ActivatedAt:  user.ActivatedAt,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
		Token:        tokenString,
	}

	return c.JSON(response)
}

func CreateUser(c *fiber.Ctx) error {
	var createUserStruct validations.CreateUserStruct
	if err := c.BodyParser(&createUserStruct); err != nil {
		return ErrorHandler(c, fiber.StatusBadRequest, "Invalid request body")
	}

	db, err := database.PostgreSQLConnection()
	if err != nil {
		return ErrorHandler(c, fiber.StatusInternalServerError, err.Error())
	}

	var users []models.User
	result := db.Find(&users)

	user, err := FindUserByEmail(createUserStruct.Email)
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
	response := reponsemodels.UserResponse{
		ID:           newUser.ID,
		Name:         newUser.Name,
		Email:        newUser.Email,
		Age:          newUser.Age,
		MemberNumber: newUser.MemberNumber,
		ActivatedAt:  newUser.ActivatedAt,
		CreatedAt:    newUser.CreatedAt,
		UpdatedAt:    newUser.UpdatedAt,
		Token:        tokenString,
	}

	return c.JSON(response)
}
