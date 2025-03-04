package dao

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type ExchangeRate struct {
	Currency string  `json:"currency"`
	Rate     float64 `json:"rate"`
}

func GetStoredRates() (map[string]float64, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create session: %w", err)
	}

	svc := dynamodb.New(sess)
	input := &dynamodb.ScanInput{
		TableName: aws.String("conversion_rates"),
	}

	result, err := svc.Scan(input)
	if err != nil {
	}

	var rates []ExchangeRate
	if err := dynamodbattribute.UnmarshalListOfMaps(result.Items, &rates); err != nil {
		return nil, fmt.Errorf("failed to unmarshal result items: %w", err)
	}

	storedRates := make(map[string]float64)
	for _, rate := range rates {
		storedRates[rate.Currency] = rate.Rate
	}
	return storedRates, nil
}
