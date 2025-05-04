package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func registerCors(app *fiber.App) {
	app.Use(cors.New())
}
