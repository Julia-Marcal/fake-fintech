package rabbitmq

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

func ConsumeMessages(connString string, queue_name string) (<-chan amqp.Delivery, *amqp.Connection, *amqp.Channel, error) {

	conn, err := amqp.Dial(connString)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, nil, nil, fmt.Errorf("failed to open a channel: %w", err)
	}

	err = ch.Qos(1, 0, false)
	if err != nil {
		ch.Close()
		conn.Close()
		return nil, nil, nil, fmt.Errorf("failed to set QoS: %w", err)
	}

	queue, err := ch.QueueDeclare(
		queue_name,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		ch.Close()
		conn.Close()
		return nil, nil, nil, fmt.Errorf("failed to declare a queue: %w", err)
	}

	messages, err := ch.Consume(
		queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		ch.Close()
		conn.Close()
		return nil, nil, nil, fmt.Errorf("failed to consume messages: %w", err)
	}

	return messages, conn, ch, nil
}
