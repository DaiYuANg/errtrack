package event_bus_module

import (
	goeventbus "github.com/stanipetrosyan/go-eventbus"
	"go.uber.org/fx"
)

var Module = fx.Module("event_bus_module", fx.Provide(newEventBus))

func newEventBus() goeventbus.EventBus {
	return goeventbus.NewEventBus()
}
