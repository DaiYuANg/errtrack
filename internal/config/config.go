package config

type ErrTrackConfig struct {
	Server   ServerConfig
	Database DatabaseConfig
	Kafka    KafkaConfig
}

type ServerConfig struct {
}

type DatabaseConfig struct{}

type KafkaConfig struct{}
