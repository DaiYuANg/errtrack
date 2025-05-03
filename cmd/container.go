package cmd

import (
	"errtrack/internal/cache"
	"errtrack/internal/config"
	"errtrack/internal/db"
	"errtrack/internal/http"
	"errtrack/internal/kafka"
	"errtrack/internal/logger"
	"errtrack/internal/mail"
	"errtrack/internal/repository"
	"errtrack/internal/schedule"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func container() *fx.App {
	return fx.New(
		config.Module,
		http.Module,
		db.Module,
		logger.Module,
		kafka.Module,
		mail.Module,
		schedule.Module,
		repository.Module,
		cache.Module,
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
	)
}
