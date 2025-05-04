package config

type ErrTrackConfig struct {
	Server   ServerConfig   `koanf:"server"`
	Database DatabaseConfig `koanf:"database"`
	Kafka    KafkaConfig    `koanf:"kafka"`
	Security SecurityConfig `koanf:"security"`
}

type DatabaseConfig struct {
	Type string
	Url  string
}

type KafkaConfig struct {
	Url string
}
