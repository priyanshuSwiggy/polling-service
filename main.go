package main

import (
	"log"
	"os"
	"os/signal"
	"polling-service/api"
	"polling-service/dao"
	"polling-service/kafka"
	"polling-service/polling"
	"polling-service/util"
	"syscall"
)

func main() {
	err := util.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	log.Println("Starting polling service...")

	producer, err := kafka.NewKafkaProducer()
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %v", err)
	}
	defer producer.Close()

	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
		<-sigChan
		log.Println("Received termination signal. Shutting down...")
		os.Exit(0)
	}()

	polling.PollAndPublish(
		api.FetchRates,
		dao.GetStoredRates,
		func(changes map[string]float64) error {
			return kafka.SendToKafka(producer, changes)
		},
	)
}
