package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	RabbitMQConnString string
	AlphaVantageAPIKey string
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, reading from environment variables")
	}

	rabbitmqUser := getEnv("RABBITMQ_USER", "guest")
	rabbitmqPassword := getEnv("RABBITMQ_PASSWORD", "guest")
	rabbitmqHost := getEnv("RABBITMQ_HOST", "localhost")
	rabbitmqPort := getEnv("RABBITMQ_PORT", "5672")

	alphaVantageAPIKey := getEnv("ALPHAVANTAGE_API_KEY", "")
	if alphaVantageAPIKey == "" {
		return nil, fmt.Errorf("error: ALPHAVANTAGE_API_KEY is not set in the environment")
	}

	cfg := &Config{
		RabbitMQConnString: fmt.Sprintf("amqp://%s:%s@%s:%s/", rabbitmqUser, rabbitmqPassword, rabbitmqHost, rabbitmqPort),
		AlphaVantageAPIKey: alphaVantageAPIKey,
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
