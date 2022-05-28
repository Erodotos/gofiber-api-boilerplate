package handlers

import (
	"github.com/erodotos/gofiber-api-boilerplate/database"
	"github.com/erodotos/gofiber-api-boilerplate/models"
	"github.com/gofiber/fiber/v2"
)

func CreateBook(c *fiber.Ctx) error {
	book := new(models.Book)

	if err := c.BodyParser(book); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	if err := database.InsertOne(book); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	c.Status(201)
	return c.JSON(book)
}

func ReadBook(c *fiber.Ctx) error {
	return c.SendString(c.BaseURL())
}

func UpdateBook(c *fiber.Ctx) error {
	return c.SendString(string(c.Body()))
}

func DeleteBook(c *fiber.Ctx) error {
	return c.SendString("Delete")
}
