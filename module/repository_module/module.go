package repository_module

import "go.uber.org/fx"

var Module = fx.Module("repository_module", fx.Provide(newUserRepository))
