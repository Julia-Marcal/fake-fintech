package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	RabbitMQConnString string
	CoinCapAPIKey      string
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, reading from environment variables")
	}

	rabbitmqUser := getEnv("RABBITMQ_USER", "guest")
	rabbitmqPassword := getEnv("RABBITMQ_PASSWORD", "guest")
	rabbitmqHost := getEnv("RABBITMQ_HOST", "localhost")
	rabbitmqPort := getEnv("RABBITMQ_PORT", "5672")

	coinCapAPIKey := getEnv("COINCAP_API_KEY", "")
	if coinCapAPIKey == "" {
		return nil, fmt.Errorf("error: COINCAP_API_KEY is not set in the environment")
	}

	cfg := &Config{
		RabbitMQConnString: fmt.Sprintf("amqp://%s:%s@%s:%s/", rabbitmqUser, rabbitmqPassword, rabbitmqHost, rabbitmqPort),
		CoinCapAPIKey:      coinCapAPIKey,
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
