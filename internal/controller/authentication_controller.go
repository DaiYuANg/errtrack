package controller

import (
	"errtrack/internal/service"
	"github.com/gofiber/fiber/v2"
)

type AuthenticationController struct {
	AuthenticationService *service.AuthenticationService
}

func (a *AuthenticationController) RegisterRoutes(app *fiber.App) {
	app.Get("/login", a.login)
}

func (a *AuthenticationController) login(ctx *fiber.Ctx) error {
	return nil
}

func NewAuthenticationController(authenticationService *service.AuthenticationService) *AuthenticationController {
	return &AuthenticationController{
		AuthenticationService: authenticationService,
	}
}
