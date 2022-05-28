package handlers

import (
	"errors"

	"github.com/erodotos/gofiber-api-boilerplate/database"
	"github.com/erodotos/gofiber-api-boilerplate/models"
	"github.com/gofiber/fiber/v2"
)

func CreateBook(c *fiber.Ctx) error {
	book := new(models.Book)

	if err := c.BodyParser(book); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	if result := database.DB.Create(&book); result.Error != nil {
		return errors.New("Database Insertion Error")
	}

	c.Status(201)
	return c.JSON(book)
}

func ReadBook(c *fiber.Ctx) error {
	book_id := c.Params("id")

	result_book := new(models.Book)

	if result := database.DB.First(&result_book, book_id); result.Error != nil {
		return errors.New("Database Selection Error")
	}

	c.SendStatus(200)
	return c.JSON(result_book)
}

func UpdateBook(c *fiber.Ctx) error {
	return c.SendString(string(c.Body()))
}

func DeleteBook(c *fiber.Ctx) error {
	return c.SendString("Delete")
}
