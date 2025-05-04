package http_module

import (
	"context"
	"errtrack/internal/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module("http_module",
	fx.Provide(newFiberApp),
	middlewareModule,
	fx.Invoke(startFiber),
)

func newFiberApp(serverConfiog *config.ServerConfig) *fiber.App {
	return fiber.New(fiber.Config{EnablePrintRoutes: true, Prefork: serverConfiog.Prefork})
}

func startFiber(lc fx.Lifecycle, app *fiber.App, serverConfiog *config.ServerConfig, logger *zap.SugaredLogger) {
	lc.Append(fx.Hook{OnStart: func(ctx context.Context) error {
		go func() {
			err := app.Listen(serverConfiog.ListenAddress())
			if err != nil {
				log.Fatal(err)
			}
		}()
		return nil
	}})
}
