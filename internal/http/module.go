package http

import (
	"context"
	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/earlydata"
	"github.com/gofiber/fiber/v2/middleware/envvar"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module("http", fx.Provide(newFiberApp), fx.Invoke(startFiber, setupMiddleware))

func newFiberApp() *fiber.App {
	return fiber.New(fiber.Config{EnablePrintRoutes: true, Prefork: false})
}

func startFiber(lc fx.Lifecycle, app *fiber.App) {
	lc.Append(fx.Hook{OnStart: func(ctx context.Context) error {
		go app.Listen(":3000")
		return nil
	}})
}

func setupMiddleware(app *fiber.App, zapLogger *zap.Logger) {
	app.Get("/metrics", monitor.New())
	app.Use(healthcheck.New(healthcheck.Config{
		LivenessProbe: func(c *fiber.Ctx) bool {
			return true
		},
		LivenessEndpoint:  "/live",
		ReadinessEndpoint: "/ready",
	}))
	app.Use(logger.New())
	app.Use(requestid.New())
	app.Use(etag.New())
	app.Use("/expose/envvars", envvar.New())
	app.Use(earlydata.New())
	app.Use(cors.New())
	app.Use(compress.New())
	app.Use(cache.New())
	app.Use(fiberzap.New(fiberzap.Config{
		Logger: zapLogger,
	}))
}
