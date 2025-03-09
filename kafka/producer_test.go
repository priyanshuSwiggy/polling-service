package kafka

import (
	"encoding/json"
	"errors"
	"polling-service/kafka/mocks"
	"polling-service/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendToKafka_Success(t *testing.T) {
	mockProducer := mocks.NewMockProducer()

	util.AppConfig.Kafka.Topic = "mock-topic"

	changes := map[string]float64{"USD": 1.2, "EUR": 0.9}
	err := SendToKafka(mockProducer, changes)
	assert.NoError(t, err)

	assert.Len(t, mockProducer.ProducedMessages, 2)

	for _, msg := range mockProducer.ProducedMessages {
		var rate ExchangeRate
		err := json.Unmarshal(msg.Value, &rate)
		assert.NoError(t, err)
		assert.Contains(t, changes, rate.Currency)
		assert.Equal(t, changes[rate.Currency], rate.Rate)
	}
}

func TestSendToKafka_ProducerError(t *testing.T) {
	mockProducer := mocks.NewMockProducer()
	mockProducer.ProduceError = errors.New("mock produce error")

	util.AppConfig.Kafka.Topic = "mock-topic"

	changes := map[string]float64{"USD": 1.2, "EUR": 0.9}
	err := SendToKafka(mockProducer, changes)
	assert.Error(t, err)
	assert.Equal(t, "mock produce error", err.Error())
}
