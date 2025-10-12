package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	domain "github.com/Julia-Marcal/fake-fintech/assets-publisher/internal/domain/entity"
	"github.com/Julia-Marcal/fake-fintech/assets-publisher/internal/infrastructure/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	action := flag.String("action", "FETCH_PRICE", "action for the task")
	market := flag.String("market", "crypto", "market for the task")
	symbol := flag.String("symbol", "BTC", "symbol for the task")
	url := flag.String("url", os.Getenv("RABBITMQ_URL"), "RabbitMQ connection URL (env RABBITMQ_URL)")
	flag.Parse()

	if *url == "" {
		log.Fatal("RabbitMQ URL must be provided via --url or RABBITMQ_URL environment variable")
	}

	conn, err := amqp.Dial(*url)
	if err != nil {
		log.Fatalf("failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("failed to open channel: %v", err)
	}
	defer ch.Close()

	task := domain.AssetTask{
		Action: *action,
		Market: *market,
		Symbol: *symbol,
	}

	err = rabbitmq.PublishMessage(ch, "assets-tasks", task)

	if err != nil {
		log.Fatalf("failed to publish a message: %v", err)
		return
	}

	fmt.Printf("Published AssetTask to queue assets-tasks: %+v\n", task)
}
