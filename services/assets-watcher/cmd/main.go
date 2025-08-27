package main

import (
	"encoding/json"
	"log"

	"github.com/Julia-Marcal/assets-watcher/internal/config"
	consumer "github.com/Julia-Marcal/assets-watcher/internal/consumer"
	domain "github.com/Julia-Marcal/assets-watcher/internal/domain"
	publisher "github.com/Julia-Marcal/assets-watcher/internal/publisher"
	services "github.com/Julia-Marcal/assets-watcher/internal/services"
	amqp "github.com/rabbitmq/amqp091-go"
)

var conn *amqp.Connection

func ProcessAssetTasks(messages <-chan amqp.Delivery, api_key string, ch *amqp.Channel) {
	for msg := range messages {
		var task domain.AssetTask
		var response domain.PublishResponse

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
			if task.Market == "crypto" {
				api_response, api_err = services.FetchPriceCoinCap(task, api_key)
			}

			if api_err != nil {
				log.Println("Failed to fetch price:", api_err)
				msg.Nack(false, false)
				continue
			}

			response = domain.PublishResponse{
				Action:   task.Action,
				Market:   task.Market,
				Response: string(api_response),
			}

			log.Printf("API Response for %s: %s", task.Symbol, string(api_response))
			publisher.PublishMessage(ch, "price-responses", response)
		}

		msg.Ack(false)
	}
}

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	conn, err := config.GetConnection(cfg.RabbitMQConnString)
	if err != nil {
		log.Fatal(err)
	}

	messages, conn, ch, err := consumer.ConsumeMessages(conn, "assets-tasks")
	if err != nil {
		log.Fatal("Failed to consume messages:", err)
	}
	defer conn.Close()
	defer ch.Close()

	forever := make(chan bool)

	go ProcessAssetTasks(messages, cfg.CoinCapAPIKey, ch)

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
