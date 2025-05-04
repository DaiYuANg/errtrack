package http_module

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
)

func registerCache(app *fiber.App) {
	app.Use(cache.New())
}
