package controllers

import (
	"new-go-api/app/models"
	reponsemodels "new-go-api/app/reponse-models"
	"new-go-api/app/validations"
	"new-go-api/platform/database"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreatePost(c *fiber.Ctx) error {
	db, err := database.PostgreSQLConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	userId := c.Locals("user")
	var user models.User
	userCheck := db.First(&user, "id = ?", userId)
	if userCheck.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "User not found"})
	}

	var createPostStruct validations.CreatePostStruct
	if err := c.BodyParser(&createPostStruct); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}
	newPost := models.Post{
		ID:      uuid.New(),
		UserID:  user.ID,
		Title:   createPostStruct.Title,
		Content: createPostStruct.Content,
	}
	postCreated := db.Create(&newPost)
	if postCreated.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": postCreated.Error.Error()})
	}
	response := reponsemodels.PostResponse{
		ID:      newPost.ID,
		UserID:  newPost.UserID,
		Title:   newPost.Title,
		Content: newPost.Content,
	}
	return c.JSON(response)
}

func GetUserWithPost(c *fiber.Ctx) error {
	userID := c.Locals("user")

	db, err := database.PostgreSQLConnection() // Get the database connection
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	// var user models.User

	// Preload the Post relationship
	// if err := db.Preload("Posts").First(&user, "id = ?", userID).Error; err != nil {
	// 	if err == gorm.ErrRecordNotFound {
	// 		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	// 	}
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	// }

	posts := []models.Post{}
	if err := db.Where("user_id = ?", userID).Find(&posts).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(posts)
}

func UpdatePost(c *fiber.Ctx) error {
	db, err := database.PostgreSQLConnection() // Assuming GetDB retrieves a singleton DB connection
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	postID := c.Params("id")
	var post models.Post
	if err := db.First(&post, "id = ?", postID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Post not found"})
	}

	var updatePostStruct validations.UpdatePostStruct
	if err := c.BodyParser(&updatePostStruct); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	// Update the fields if provided
	if updatePostStruct.Title != "" {
		post.Title = updatePostStruct.Title
	}
	if updatePostStruct.Content != "" {
		post.Content = updatePostStruct.Content
	}

	// Save the updated post
	if err := db.Save(&post).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update post"})
	}

	// Prepare response
	response := reponsemodels.PostResponse{
		ID:      post.ID,
		UserID:  post.UserID,
		Title:   post.Title,
		Content: post.Content,
	}

	return c.JSON(response)
}

func DeletePost(c *fiber.Ctx) error {
	db, err := database.PostgreSQLConnection() // Assuming GetDB retrieves a singleton DB connection
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	postID := c.Params("id")
	var post models.Post
	if err := db.First(&post, "id = ?", postID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Post not found"})
	}

	// Delete the post
	if err := db.Delete(&post).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete post"})
	}

	return c.JSON(fiber.Map{"message": "Post deleted successfully"})
}
