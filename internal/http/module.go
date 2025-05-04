package http

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

var Module = fx.Module("http",
	fx.Provide(newFiberApp),
	middlewareModule,
	fx.Invoke(startFiber),
)

func newFiberApp() *fiber.App {
	return fiber.New(fiber.Config{EnablePrintRoutes: true, Prefork: true})
}

func startFiber(lc fx.Lifecycle, app *fiber.App) {
	lc.Append(fx.Hook{OnStart: func(ctx context.Context) error {
		go app.Listen(":3000")
		return nil
	}})
}
