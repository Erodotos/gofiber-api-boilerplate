package handlers

import (
	"errors"

	"github.com/erodotos/gofiber-api-boilerplate/common"
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
		return errors.New("error creating new book")
	}

	c.Status(201)
	return c.JSON(book)
}

func ReadBook(c *fiber.Ctx) error {
	book_id := c.Params("id")

	result_book := new(models.Book)

	if result := database.DB.First(&result_book, book_id); result.Error != nil {
		return errors.New("error finding book")
	}

	c.SendStatus(200)
	return c.JSON(result_book)
}

func UpdateBook(c *fiber.Ctx) error {
	book_id := c.Params("id")

	result_book := new(models.Book)

	// Fetch the queried book
	if result := database.DB.First(&result_book, book_id); result.Error != nil {
		return c.Status(404).JSON(common.Response{Error: result.Error.Error()})
	}

	// Update book fields according to the new data
	if err := c.BodyParser(&result_book); err != nil {
		return c.Status(400).JSON(common.Response{Error: err.Error()})
	}

	// Save the changes to the database
	if result := database.DB.Save(result_book); result.Error != nil {
		return c.Status(500).JSON(common.Response{Error: result.Error.Error()})
	}

	return c.Status(200).JSON(common.Response{Result: result_book})
}

func DeleteBook(c *fiber.Ctx) error {
	book_id := c.Params("id")

	if result := database.DB.Delete(&models.Book{}, book_id); result.Error != nil {
		return c.Status(500).JSON(common.Response{Error: result.Error.Error()})
	}

	return c.Status(200).JSON(common.Response{Result: "success"})
}
