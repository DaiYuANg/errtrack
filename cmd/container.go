package cmd

import (
	"errtrack/module/authentication"
	"errtrack/module/cache_module"
	"errtrack/module/config_module"
	"errtrack/module/controller_module"
	"errtrack/module/db_module"
	"errtrack/module/event_bus_module"
	"errtrack/module/http_module"
	"errtrack/module/id_generator_module"
	"errtrack/module/kafka_module"
	"errtrack/module/logger_module"
	"errtrack/module/mail_module"
	"errtrack/module/repository_module"
	"errtrack/module/schedule_mdoule"
	"errtrack/module/service_module"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func container() *fx.App {
	return fx.New(
		config_module.Module,
		id_generator_module.Module,
		http_module.Module,
		controller_module.Module,
		authentication.Module,
		db_module.Module,
		logger_module.Module,
		kafka_module.Module,
		mail_module.Module,
		service_module.Module,
		event_bus_module.Module,
		schedule_mdoule.Module,
		repository_module.Module,
		cache_module.Module,
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
	)
}
