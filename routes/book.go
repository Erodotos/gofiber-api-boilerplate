package routes

import (
	"github.com/erodotos/gofiber-api-boilerplate/handlers"
	"github.com/gofiber/fiber/v2"
)

func BookRouter(api fiber.Router) {
	bookGroup := api.Group("/book")
	bookGroup.Get("/all", handlers.ReadAllBooks)
	bookGroup.Get("/:id", handlers.ReadBook)
	bookGroup.Post("/", handlers.CreateBook)
	bookGroup.Put("/:id", handlers.UpdateBook)
	bookGroup.Delete("/:id", handlers.DeleteBook)
}
