package publisher

import (
	"fmt"

	"encoding/json"

	domain "github.com/Julia-Marcal/assets-watcher/internal/domain"
	amqp "github.com/rabbitmq/amqp091-go"
)

func PublishMessage(ch *amqp.Channel, queueName string, body domain.PublishResponse) error {
	_, err := ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return fmt.Errorf("failed to declare a queue: %w", err)
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("failed to marshal body: %w", err)
	}

	err = ch.Publish(
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        jsonBody,
		},
	)

	if err != nil {
		return fmt.Errorf("failed to publish a message: %w", err)
	}

	return nil
}
