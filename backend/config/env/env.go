package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}
}

func setEnv() {
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	database := os.Getenv("POSTGRES_DATABASE")

	if host == "" || port == "" || user == "" || password == "" || database == "" {
		panic("Missing required environment variables for PostgreSQL connection")
	}
}

func GetPostgresConnectionString() string {
	setEnv()
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	database := os.Getenv("POSTGRES_DATABASE")

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, database)
}

func GetDatabase() string {
	setEnv()
	return os.Getenv("POSTGRES_DATABASE")
}

func SetSalt() []byte {
	loadEnv()
	return []byte(os.Getenv("MY_SALT"))
}

func GetJwtKey() string {
	loadEnv()
	return os.Getenv("JWT_KEY")
}
