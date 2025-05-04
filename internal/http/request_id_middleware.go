package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func registerRequestID(app *fiber.App) {
	app.Use(requestid.New())
}
