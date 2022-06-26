package handlers

import (
	"github.com/erodotos/gofiber-api-boilerplate/dal"
	"github.com/erodotos/gofiber-api-boilerplate/models"
	"github.com/gofiber/fiber/v2"
)

func CreateBook(c *fiber.Ctx) error {
	book := new(models.Book)

	if err := c.BodyParser(book); err != nil {
		return c.Status(400).JSON(fiber.Map{"data": nil, "error": err.Error()})
	}

	if err := dal.CreateBook(book); err != nil {
		return c.Status(500).JSON(fiber.Map{"data": nil, "error": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{"data": "success", "error": nil})
}

func ReadBook(c *fiber.Ctx) error {
	book_id := c.Params("id")

	err, result := dal.ReadBook(book_id)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"data": nil, "error": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"data": result, "error": nil})
}

func ReadAllBooks(c *fiber.Ctx) error {

	err, result := dal.ReadAllBooks()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"data": nil, "error": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"data": result, "error": nil})
}

func UpdateBook(c *fiber.Ctx) error {
	book_id := c.Params("id")

	result_book := new(models.Book)

	// Fetch the queried book
	err, _ := dal.ReadBook(book_id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"data": nil, "error": err.Error()})
	}

	// Update book fields according to the new data
	if err := c.BodyParser(&result_book); err != nil {
		return c.Status(400).JSON(fiber.Map{"data": nil, "error": err.Error()})
	}

	// Save the changes to the database
	if err := dal.UpdateBook(result_book); err != nil {
		return c.Status(500).JSON(fiber.Map{"data": nil, "error": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"data": result_book, "error": nil})
}

func DeleteBook(c *fiber.Ctx) error {
	book_id := c.Params("id")

	if err := dal.DeleteBook(book_id); err != nil {
		return c.Status(500).JSON(fiber.Map{"data": nil, "error": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"data": "success", "error": nil})
}
