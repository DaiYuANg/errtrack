package http_module

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/earlydata"
)

func registerEarlydata(app *fiber.App) {
	app.Use(earlydata.New())
}
