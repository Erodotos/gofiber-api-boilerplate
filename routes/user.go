package routes

import (
	"github.com/erodotos/gofiber-api-boilerplate/handlers"
	"github.com/gofiber/fiber/v2"
)

func UserRouter(app *fiber.App) {
	bookGroup := app.Group("/api/")
	bookGroup.Get("/user/:id", handlers.ReadUser)
	// bookGroup.Get("/users", handlers.ReadAllUsers)
	bookGroup.Post("/user/", handlers.CreateUser)
	// bookGroup.Put("/user/:id", handlers.UpdateUser)
	// bookGroup.Delete("/user/:id", handlers.DeleteUser)
}
