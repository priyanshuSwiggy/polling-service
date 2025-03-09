package mocks

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type MockProducer struct {
	ProducedMessages []*kafka.Message
	ProduceError     error
}

func (m *MockProducer) Produce(msg *kafka.Message, deliveryChan chan kafka.Event) error {
	if m.ProduceError != nil {
		return m.ProduceError
	}
	m.ProducedMessages = append(m.ProducedMessages, msg)
	return nil
}

func (m *MockProducer) Flush(timeoutMs int) int {
	return 0
}

func (m *MockProducer) Close() {}

func NewMockProducer() *MockProducer {
	return &MockProducer{
		ProducedMessages: make([]*kafka.Message, 0),
	}
}
