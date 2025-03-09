package mocks

var FetchRatesFunc func() (map[string]float64, error)

func FetchRates() (map[string]float64, error) {
	return FetchRatesFunc()
}

var GetStoredRatesFunc func() (map[string]float64, error)

func GetStoredRates() (map[string]float64, error) {
	return GetStoredRatesFunc()
}

var SendToKafkaFunc func(changes map[string]float64) error

func SendToKafka(changes map[string]float64) error {
	return SendToKafkaFunc(changes)
}
