package main

import "github.com/gofiber/fiber/v2"

func MiddlewareInit(app *fiber.App) {
	app.Use(initPageNotFound)
}

func initPageNotFound(c *fiber.Ctx) error {
	return c.SendStatus(404)
}
