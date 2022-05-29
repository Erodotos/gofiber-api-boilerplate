package middleware

import "github.com/gofiber/fiber/v2"

func initPageNotFound(c *fiber.Ctx) error {
	return c.SendStatus(404)
}
