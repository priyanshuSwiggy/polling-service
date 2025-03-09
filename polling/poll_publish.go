package polling

import (
	"log"
	"time"
)

type FetchRatesFunc func() (map[string]float64, error)
type GetStoredRatesFunc func() (map[string]float64, error)
type SendToKafkaFunc func(changes map[string]float64) error

func detectChanges(newRates, storedRates map[string]float64) map[string]float64 {
	changes := make(map[string]float64)
	for currency, newRate := range newRates {
		if storedRate, exists := storedRates[currency]; !exists || storedRate != newRate {
			changes[currency] = newRate
		}
	}
	return changes
}

func PollAndPublish(fetchRates FetchRatesFunc, getStoredRates GetStoredRatesFunc, sendToKafka SendToKafkaFunc) {
	for {
		newRates, err := fetchRates()
		if err != nil {
			log.Println("Failed to fetch rates:", err)
			continue
		}
		storedRates, err := getStoredRates()
		if err != nil {
			log.Println("Failed to get stored rates:", err)
			continue
		}
		changes := detectChanges(newRates, storedRates)
		if len(changes) > 0 {
			if err := sendToKafka(changes); err != nil {
				log.Println("Failed to send to Kafka:", err)
			}
		}
		time.Sleep(1 * time.Hour)
	}
}
