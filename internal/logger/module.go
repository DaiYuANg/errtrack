package logger

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module("logger", fx.Provide(newLogger))

func newLogger() *zap.Logger {
	logger := zap.NewExample()
	defer logger.Sync() // flushes buffer, if any
	return logger
}
