package polling

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPollAndPublish_Success(t *testing.T) {
	fetchRates := func() (map[string]float64, error) {
		return map[string]float64{"USD": 1.2, "EUR": 0.9}, nil
	}

	getStoredRates := func() (map[string]float64, error) {
		return map[string]float64{"USD": 1.1, "EUR": 0.9}, nil
	}

	sendToKafka := func(changes map[string]float64) error {
		assert.Equal(t, map[string]float64{"USD": 1.2}, changes)
		return nil
	}

	go func() {
		PollAndPublish(fetchRates, getStoredRates, sendToKafka)
	}()

	time.Sleep(2 * time.Second)
}

func TestPollAndPublish_FetchRatesError(t *testing.T) {
	fetchRates := func() (map[string]float64, error) {
		return nil, errors.New("fetch error")
	}

	getStoredRates := func() (map[string]float64, error) {
		return map[string]float64{"USD": 1.1, "EUR": 0.9}, nil
	}

	sendToKafka := func(changes map[string]float64) error {
		t.FailNow()
		return nil
	}

	go func() {
		PollAndPublish(fetchRates, getStoredRates, sendToKafka)
	}()

	time.Sleep(2 * time.Second)
}

func TestPollAndPublish_GetStoredRatesError(t *testing.T) {
	fetchRates := func() (map[string]float64, error) {
		return map[string]float64{"USD": 1.2, "EUR": 0.9}, nil
	}

	getStoredRates := func() (map[string]float64, error) {
		return nil, errors.New("get stored rates error")
	}

	sendToKafka := func(changes map[string]float64) error {
		t.FailNow()
		return nil
	}

	go func() {
		PollAndPublish(fetchRates, getStoredRates, sendToKafka)
	}()

	time.Sleep(2 * time.Second)
}

func TestPollAndPublish_SendToKafkaError(t *testing.T) {
	fetchRates := func() (map[string]float64, error) {
		return map[string]float64{"USD": 1.2, "EUR": 0.9}, nil
	}

	getStoredRates := func() (map[string]float64, error) {
		return map[string]float64{"USD": 1.1, "EUR": 0.9}, nil
	}

	sendToKafka := func(changes map[string]float64) error {
		return errors.New("send to kafka error")
	}

	go func() {
		PollAndPublish(fetchRates, getStoredRates, sendToKafka)
	}()

	time.Sleep(2 * time.Second)
}
