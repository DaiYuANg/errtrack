package service_module

import (
	"errtrack/internal/service"
	"go.uber.org/fx"
)

var Module = fx.Module("service_module",
	fx.Provide(
		service.NewJWTService,
		service.NewAuthenticationService,
	),
)
