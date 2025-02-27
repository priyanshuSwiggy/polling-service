package db

import (
	"polling-service/util"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ExchangeRate struct {
	Currency string  `json:"currency"`
	Rate     float64 `json:"rate"`
}

func GetStoredRates() (map[string]float64, error) {
	db, err := gorm.Open(postgres.Open(util.AppConfig.DB.Conn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	var rates []ExchangeRate
	storedRates := make(map[string]float64)
	if err := db.Table(util.AppConfig.DB.Table).Find(&rates).Error; err != nil {
		return nil, err
	}
	for _, rate := range rates {
		storedRates[rate.Currency] = rate.Rate
	}
	return storedRates, nil
}
