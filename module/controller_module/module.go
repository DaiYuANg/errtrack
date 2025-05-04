package controller_module

import (
	"errtrack/internal/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
	"go.uber.org/fx"
)

var Module = fx.Module("controller_module",
	fx.Provide(
		fx.Annotate(
			controller.NewAuthenticationController,
			fx.As(new(Controller)),
			fx.ResultTags(`group:"controllers"`),
		),
		fx.Annotate(
			controller.NewRBACController,
			fx.As(new(Controller)),
			fx.ResultTags(`group:"controllers"`),
		),
	),
	fx.Invoke(bindingController))

type BindingParams struct {
	fx.In
	App        *fiber.App
	Controller []Controller `group:"controllers"`
}

func bindingController(params BindingParams) {
	lo.ForEach(params.Controller, func(item Controller, index int) {
		item.RegisterRoutes(params.App)
	})
}
