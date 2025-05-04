package controller_module

import "github.com/gofiber/fiber/v2"

type Controller interface {
	RegisterRoutes(app *fiber.App)
}
