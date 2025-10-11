package services

import (
	"encoding/json"
	"log"

	"github.com/Julia-Marcal/fake-fintech/internal/config"
	domain "github.com/Julia-Marcal/fake-fintech/internal/domain/entity"
	"github.com/Julia-Marcal/fake-fintech/internal/infrastructure/rabbitmq"
	"github.com/Julia-Marcal/fake-fintech/internal/service"
	amqp "github.com/rabbitmq/amqp091-go"
)

func MessageConsumer(conn *amqp.Connection, config *config.Config) {
	messages, conn, ch, err := rabbitmq.ConsumeMessages(conn, "assets-tasks")
	if err != nil {
		log.Fatal("Failed to consume messages:", err)
	}
	defer conn.Close()
	defer ch.Close()

	forever := make(chan bool)

	go ProcessAssetTasks(messages, config.CoinCapAPIKey, ch)

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func ProcessAssetTasks(messages <-chan amqp.Delivery, apiKey string, ch *amqp.Channel) {
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

		var apiResponse []byte
		var apiErr error

		if task.Action == "FETCH_PRICE" {
			if task.Market == "crypto" {
				apiResponse, apiErr = service.FetchPriceCoinCap(task, apiKey)
			}

			if apiErr != nil {
				log.Println("Failed to fetch price:", apiErr)
				msg.Nack(false, false)
				continue
			}

			response = domain.PublishResponse{
				Action:   task.Action,
				Market:   task.Market,
				Response: string(apiResponse),
			}

			log.Printf("API Response for %s: %s", task.Symbol, string(apiResponse))
			rabbitmq.PublishMessage(ch, "price-responses", response)
		}

		msg.Ack(false)
	}
}
