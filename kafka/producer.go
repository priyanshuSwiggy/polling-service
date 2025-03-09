package kafka

import (
	"encoding/json"
	"log"
	"polling-service/util"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Producer interface {
	Produce(msg *kafka.Message, deliveryChan chan kafka.Event) error
	Flush(timeoutMs int) int
	Close()
}

type ExchangeRate struct {
	Currency string  `json:"currency"`
	Rate     float64 `json:"rate"`
}

func SendToKafka(producer Producer, changes map[string]float64) error {
	for currency, rate := range changes {
		message, _ := json.Marshal(ExchangeRate{Currency: currency, Rate: rate})
		log.Printf("Sending message to Kafka: %s", message)
		err := producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &util.AppConfig.Kafka.Topic, Partition: kafka.PartitionAny},
			Value:          message,
		}, nil)
		if err != nil {
			return err
		}
	}
	producer.Flush(15 * 1000)
	producer.Close()
	return nil
}

func NewKafkaProducer() (Producer, error) {
	return kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": util.AppConfig.Kafka.Brokers})
}
