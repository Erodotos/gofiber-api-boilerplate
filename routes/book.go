package routes

import (
	"github.com/erodotos/gofiber-api-boilerplate/handlers"
	"github.com/gofiber/fiber/v2"
)

func BookRouter(app *fiber.App) {
	bookGroup := app.Group("/api")
	bookGroup.Get("/book/:id", handlers.ReadBook)
	bookGroup.Post("/book/", handlers.CreateBook)
	bookGroup.Patch("/book/:id", handlers.UpdateBook)
	bookGroup.Delete("/book/:id", handlers.DeleteBook)
}