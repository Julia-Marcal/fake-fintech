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

func StartConsumer(messages <-chan amqp.Delivery, api_key string) {
	for msg := range messages {
		var task task.AssetTask

		err := json.Unmarshal(msg.Body, &task)
		if err != nil {
			log.Println("Failed to unmarshal message:", err)
			msg.Nack(false, false)
			continue
		}

		log.Printf("Processing task: %+v\n", task)

		var api_response []byte
		var api_err error

		if task.Action == "FETCH_PRICE" {
			api_response, api_err = services.FetchPriceCoinCap(task, api_key)

			if api_err != nil {
				log.Println("Failed to fetch price:", api_err)
				msg.Nack(false, false)
				continue
			}

			log.Printf("API Response for %s: %s", task.Symbol, string(api_response))
		}

		msg.Ack(false)
	}
}

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	messages, conn, ch, err := rabbitmq.ConsumeMessages(cfg.RabbitMQConnString, "assets-tasks")
	if err != nil {
		log.Fatal("Failed to consume messages:", err)
	}
	defer conn.Close()
	defer ch.Close()

	forever := make(chan bool)

	go StartConsumer(messages, cfg.CoinCapAPIKey)

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
