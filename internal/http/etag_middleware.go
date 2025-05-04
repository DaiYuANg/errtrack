package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/etag"
)

func RegisterEtagMiddleware(app *fiber.App) {
	app.Use(etag.New())
}
