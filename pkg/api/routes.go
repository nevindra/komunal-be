// pkg/api/routes.go
package api

import (
	"komunal-be/pkg/api/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to Komunal API")
	})

	users := api.Group("/users")
	// USERS QUERIES
	users.Get("/", handlers.GetUsers) // TODO: Implement pagination,limit, and sorting
	users.Get("/:id", handlers.GetUserById)
	users.Get("username/:username", handlers.GetUserByUsername)
	users.Put("/:id", handlers.UpdateUser)
	users.Get("/:username/question", handlers.GetQuestionByUsername) // TODO: Implement pagination,limit, and sorting

	questions := users.Group("/questions")
	// QUESTIONS QUERIES
	questions.Get("/", handlers.GetQuestions)
	questions.Get("questions/:id", handlers.GetQuestionById)
	questions.Post("/", handlers.CreateQuestion)
	questions.Put("/:id", handlers.UpdateQuestion)
	questions.Delete("/:id", handlers.DeleteQuestion)

}
