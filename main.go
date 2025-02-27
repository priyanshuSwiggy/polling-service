package main

import (
	"log"
	"polling-service/polling"
	"polling-service/util"
)

func main() {
	err := util.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	log.Println("Starting polling service...")
	polling.PollAndPublish()
}
