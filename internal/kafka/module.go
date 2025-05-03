package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module("kafka", fx.Provide(kafkaConnection), fx.Invoke())

func kafkaConnection(log *zap.SugaredLogger) (*kafka.Conn, error) {
	topic := "my-topic"
	partition := 0
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	return conn, nil
}
