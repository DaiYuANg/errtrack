package http_module

import "go.uber.org/fx"

var middlewareModule = fx.Module("middleware",
	fx.Invoke(
		registerMonitor,
		registerHealthCheck,
		registerCache,
		registerEarlydata,
		registerCompressMiddleware,
		registerEnvvars,
		registerLogger,
		registerCors,
		registerRequestID,
		RegisterEtagMiddleware,
	),
)
