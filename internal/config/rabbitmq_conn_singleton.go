package config

import (
	"fmt"
	"sync"

	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	conn     *amqp.Connection
	connOnce sync.Once
	connErr  error
)

func GetConnection(connString string) (*amqp.Connection, error) {
	connOnce.Do(func() {
		c, err := amqp.Dial(connString)
		if err != nil {
			connErr = fmt.Errorf("failed to connect to RabbitMQ: %w", err)
			return
		}
		conn = c
	})
	return conn, connErr
}
