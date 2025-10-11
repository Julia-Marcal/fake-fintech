package main

import (
	"log"

	services "github.com/Julia-Marcal/fake-fintech/internal/assets/services"
	"github.com/Julia-Marcal/fake-fintech/internal/config"
	amqp "github.com/rabbitmq/amqp091-go"
)

var conn *amqp.Connection

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	conn, err := config.GetConnection(cfg.RabbitMQConnString)
	if err != nil {
		log.Fatal(err)
	}

	services.MessageConsumer(conn, cfg)
}
