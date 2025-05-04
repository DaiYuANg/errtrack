package config

func DefaultConfig() ErrTrackConfig {

	return ErrTrackConfig{
		Server: ServerConfig{
			Port:    19090,
			Prefork: false,
		},
		Database: DatabaseConfig{
			Type: "sqlite",
			Url:  "file::memory:?cache_module=shared",
		},
		Kafka: KafkaConfig{
			Url: "localhost:9092",
		},
	}
}
