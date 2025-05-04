package controller

import (
	"github.com/gofiber/fiber/v2"
)

type AuthenticationController struct{}

func (a *AuthenticationController) RegisterRoutes(app *fiber.App) {
	app.Get("/login", a.login)
}

func (a *AuthenticationController) login(ctx *fiber.Ctx) error {

	return nil
}

func newAuthenticationController() *AuthenticationController {
	return &AuthenticationController{}
}
