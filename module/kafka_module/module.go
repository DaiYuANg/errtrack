package kafka_module

import (
	"context"
	"errtrack/internal/config"
	"github.com/segmentio/kafka-go"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module("kafka_module", fx.Provide(kafkaConnection), fx.Invoke())

func kafkaConnection(log *zap.SugaredLogger, config *config.KafkaConfig) (*kafka.Conn, error) {
	topic := "my-topic"
	partition := 0
	conn, err := kafka.DialLeader(
		context.Background(),
		"tcp",
		config.Url,
		topic,
		partition,
	)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	return conn, nil
}
