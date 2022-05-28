package routes

import (
	"github.com/erodotos/gofiber-api-boilerplate/handlers"
	"github.com/gofiber/fiber/v2"
)

func BookRouter(app *fiber.App) {
	bookGroup := app.Group("/api")
	bookGroup.Get("/book/:id", handlers.ReadBook)
	bookGroup.Get("/books", handlers.ReadAllBooks)
	bookGroup.Post("/book/", handlers.CreateBook)
	bookGroup.Put("/book/:id", handlers.UpdateBook)
	bookGroup.Delete("/book/:id", handlers.DeleteBook)
}
