package id_generator

import (
	"github.com/sony/sonyflake"
	"go.uber.org/fx"
	"time"
)

var Module = fx.Module("id_generator", fx.Provide(newSnowflake))

func newSnowflake() (*sonyflake.Sonyflake, error) {
	return sonyflake.New(sonyflake.Settings{
		StartTime: time.Now(),
	})
}
