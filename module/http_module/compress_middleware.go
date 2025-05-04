package http_module

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
)

func registerCompressMiddleware(app *fiber.App) {
	app.Use(compress.New())
}
