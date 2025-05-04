package config_module

import (
	"errtrack/internal/config"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/structs"
	"github.com/knadh/koanf/v2"
	"github.com/samber/lo"
	"go.uber.org/fx"
	"strings"
)

const EnvPrefix = "ERR_TRACK_"

var Module = fx.Module("config_module", fx.Provide(createKoanf, parseConfig))

func createKoanf() *koanf.Koanf {
	kc := koanf.Conf{
		Delim:       ".",
		StrictMerge: true,
	}
	return koanf.NewWithConf(kc)
}

type ParseConfigResult struct {
	fx.Out
	Config         *config.ErrTrackConfig
	ServerConfig   *config.ServerConfig
	DatabaseConfig *config.DatabaseConfig
	KafkaConfig    *config.KafkaConfig
	SecurityConfig *config.SecurityConfig
}

type ParseParams struct {
	fx.In
	K *koanf.Koanf
}

func parseConfig(params ParseParams) (ParseConfigResult, error) {
	k := params.K
	c := config.DefaultConfig()
	lo.Must0(k.Load(structs.Provider(c, "default"), nil))
	lo.Must0(k.Load(env.Provider(EnvPrefix, ".", func(s string) string {
		return strings.Replace(strings.ToLower(
			strings.TrimPrefix(s, EnvPrefix)), "_", ".", -1)
	}), nil))

	var out config.ErrTrackConfig
	lo.Must0(k.Unmarshal("", &out), "error unmarshalling config_module")
	configMapping := &out
	return ParseConfigResult{
		Config:         configMapping,
		DatabaseConfig: &configMapping.Database,
		ServerConfig:   &configMapping.Server,
		KafkaConfig:    &configMapping.Kafka,
		SecurityConfig: &configMapping.Security,
	}, nil
}
