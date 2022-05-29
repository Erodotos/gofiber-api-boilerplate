package handlers

import (
	"github.com/erodotos/gofiber-api-boilerplate/common"
	"github.com/erodotos/gofiber-api-boilerplate/database"
	// "github.com/erodotos/gofiber-api-boilerplate/database"
	"github.com/erodotos/gofiber-api-boilerplate/models"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(common.Response{Error: err.Error()})
	}

	if result := database.DB.Create(&user); result.Error != nil {
		return c.Status(500).JSON(common.Response{Error: result.Error.Error()})
	}

	return c.Status(201).JSON(common.Response{Message: user})
}

func ReadUser(c *fiber.Ctx) error {
	user_id := c.Params("id")

	user := new(models.User)

	if result := database.DB.First(&user, user_id); result.Error != nil {
		return c.Status(404).JSON(common.Response{Error: result.Error.Error()})
	}

	return c.Status(200).JSON(common.Response{Message: user})
}

func ReadAllUsers(c *fiber.Ctx) error {
	users := []models.User{}

	if result := database.DB.Find(&users); result.Error != nil {
		return c.Status(404).JSON(common.Response{Error: result.Error.Error()})
	}

	return c.Status(200).JSON(common.Response{Message: users})
}

func UpdateUser(c *fiber.Ctx) error {
	user_id := c.Params("id")

	result_user := new(models.User)

	// Fetch the queried book
	if result := database.DB.First(&result_user, user_id); result.Error != nil {
		return c.Status(404).JSON(common.Response{Error: result.Error.Error()})
	}

	// Update book fields according to the new data
	if err := c.BodyParser(&result_user); err != nil {
		return c.Status(400).JSON(common.Response{Error: err.Error()})
	}

	// Save the changes to the database
	if result := database.DB.Save(result_user); result.Error != nil {
		return c.Status(500).JSON(common.Response{Error: result.Error.Error()})
	}

	return c.Status(200).JSON(common.Response{Message: result_user})
}

func DeleteUser(c *fiber.Ctx) error {
	user_id := c.Params("id")

	if result := database.DB.Delete(&models.Book{}, user_id); result.Error != nil {
		return c.Status(500).JSON(common.Response{Error: result.Error.Error()})
	}

	return c.Status(200).JSON(common.Response{Message: "success"})
}
