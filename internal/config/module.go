package config

import (
	"github.com/knadh/koanf/v2"
	"go.uber.org/fx"
)

var Module = fx.Module("config", fx.Provide(createKoanf))

func createKoanf() *koanf.Koanf {
	return koanf.New(".")
}
