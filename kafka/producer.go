package kafka

import (
	"encoding/json"
	"polling-service/util"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type ExchangeRate struct {
	Currency string  `json:"currency"`
	Rate     float64 `json:"rate"`
}

func SendToKafka(changes map[string]float64) error {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": util.AppConfig.Kafka.Brokers})
	if err != nil {
		return err
	}
	defer func() {
		producer.Flush(15 * 1000)
		producer.Close()
	}()

	for currency, rate := range changes {
		message, _ := json.Marshal(ExchangeRate{Currency: currency, Rate: rate})
		producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &util.AppConfig.Kafka.Topic, Partition: kafka.PartitionAny},
			Value:          message,
		}, nil)
	}
	return nil
}
