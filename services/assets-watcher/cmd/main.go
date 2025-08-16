package main

import (
	"encoding/json"
	"log"

	"github.com/Julia-Marcal/assets-watcher/internal/config"
	rabbitmq "github.com/Julia-Marcal/assets-watcher/internal/consumer"
	task "github.com/Julia-Marcal/assets-watcher/internal/domain"
	services "github.com/Julia-Marcal/assets-watcher/internal/services"
	amqp "github.com/rabbitmq/amqp091-go"
)

func StartConsumer(messages <-chan amqp.Delivery, apiKey string) {
	for msg := range messages {
		var task task.AssetTask
		err := json.Unmarshal(msg.Body, &task)
		if err != nil {
			log.Println("Failed to unmarshal message:", err)
			msg.Nack(false, false)
			continue
		}

		log.Printf("Processing task: %+v\n", task)

		var apiResponse []byte
		var apiErr error

		if task.Action == "FETCH_PRICE" {
			apiResponse, apiErr = services.FetchPriceAlphaVantage(task, apiKey)
			if apiErr != nil {
				log.Println("Failed to fetch price:", apiErr)
				msg.Nack(false, false)
				continue
			}
			log.Printf("API Response for %s: %s", task.Symbol, string(apiResponse))
		}

		msg.Ack(false)
	}
}

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	messages, conn, ch, err := rabbitmq.ConsumeMessages(cfg.RabbitMQConnString, "assets-watcher")
	if err != nil {
		log.Fatal("Failed to consume messages:", err)
	}
	defer conn.Close()
	defer ch.Close()

	forever := make(chan bool)

	go StartConsumer(messages, cfg.AlphaVantageAPIKey)

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
