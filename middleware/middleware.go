package middleware

import "github.com/gofiber/fiber/v2"

func MiddlewareInit(app *fiber.App) {
	app.Use(initPageNotFound)
}
