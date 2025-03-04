package polling

import (
	"log"
	"polling-service/api"
	"polling-service/dao"
	"polling-service/kafka"
	"time"
)

func detectChanges(newRates, storedRates map[string]float64) map[string]float64 {
	changes := make(map[string]float64)
	for currency, newRate := range newRates {
		if storedRate, exists := storedRates[currency]; !exists || storedRate != newRate {
			changes[currency] = newRate
		}
	}
	return changes
}

func PollAndPublish() {
	for {
		newRates, err := api.FetchRates()
		if err != nil {
			log.Println("Failed to fetch rates:", err)
			continue
		}
		storedRates, err := dao.GetStoredRates()
		if err != nil {
			log.Println("Failed to get stored rates:", err)
			continue
		}
		changes := detectChanges(newRates, storedRates)
		if len(changes) > 0 {
			if err := kafka.SendToKafka(changes); err != nil {
				log.Println("Failed to send to Kafka:", err)
			}
		}
		time.Sleep(1 * time.Hour)
	}
}
