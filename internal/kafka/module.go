package kafka

import (
	"context"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module("kafka", fx.Provide(kafkaConnection), fx.Invoke(listen))

func kafkaConnection() *kafka.Producer {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		panic(err)
	}

	defer p.Close()

	return p
}

func listen(lc fx.Lifecycle, kafkaConnection *kafka.Producer, logger *zap.Logger) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				for e := range kafkaConnection.Events() {
					switch ev := e.(type) {
					case *kafka.Message:
						if ev.TopicPartition.Error != nil {
							fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
						} else {
							fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
						}
					}
				}
			}()
			return nil
		},
	})
}
