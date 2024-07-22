package handlers

import (
	"komunal-be/pkg/api/db"
	"komunal-be/pkg/api/models"

	"github.com/gofiber/fiber/v2"
)

func GetQuestions(c *fiber.Ctx) error {
	var questions []models.Question
	result := db.DB.Find(&questions)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch questions",
		})
	}
	return c.JSON(questions)
}

func GetQuestionByUsername(c *fiber.Ctx) error {
	username := c.Params("username")
	var user models.User
	var questions []struct {
		models.Question
		LikeCount    int `json:"like_count"`
		CommentCount int `json:"comment_count"`
	}

	// Find the user by username
	if err := db.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	// Find questions associated with the user, including like and comment counts
	result := db.DB.Model(&models.Question{}).
		Select("questions.*, COUNT(DISTINCT likes.like_id) as like_count, COUNT(DISTINCT comments.comment_id) as comment_count").
		Joins("LEFT JOIN likes ON likes.question_id = questions.question_id").
		Joins("LEFT JOIN comments ON comments.question_id = questions.question_id").
		Where("questions.user_id = ?", user.ID).
		Group("questions.question_id").
		Find(&questions)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch questions",
		})
	}

	return c.JSON(questions)
}

func GetQuestionById(c *fiber.Ctx) error {
	id := c.Params("id")
	var question models.Question
	result := db.DB.Where("question_id = ?", id).First(&question)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Question not found",
		})
	}
	return c.JSON(question)
}

func CreateQuestion(c *fiber.Ctx) error {
	var question models.Question

	// Parse the create data
	if err := c.BodyParser(&question); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid create data",
		})
	}

	// Create the question
	if err := db.DB.Create(&question).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create question",
		})
	}

	return c.JSON(question)
}

func UpdateQuestion(c *fiber.Ctx) error {
	id := c.Params("id")
	var question models.Question

	// Find the question by ID
	if err := db.DB.First(&question, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Question not found",
		})
	}

	// Parse the update data
	var updateData models.Question
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid update data",
		})
	}

	// Update the question
	if err := db.DB.Model(&question).Updates(updateData).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update question",
		})
	}

	return c.JSON(question)
}

func DeleteQuestion(c *fiber.Ctx) error {
	id := c.Params("id")
	var question models.Question

	// Find the question by ID
	if err := db.DB.First(&question, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Question not found",
		})
	}

	// Delete the question
	if err := db.DB.Delete(&question).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete question",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
