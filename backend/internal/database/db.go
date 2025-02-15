package database

import (
	"fmt"
	"os"
	"sync"

	user "github.com/Julia-Marcal/fake-fintech/internal/schemas/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	once = sync.Once{}
	db   *gorm.DB
)

func NewPostgres() *gorm.DB {
	once.Do(func() {
		var err error

		connectionStr := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_PORT"),
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_DATABASE"),
		)
		fmt.Println("about to connect to database with connection string:", connectionStr)

		db, err = gorm.Open(postgres.Open(connectionStr), &gorm.Config{
			SkipDefaultTransaction: true,
		})
		if err != nil {
			panic(fmt.Sprintf("failed to connect to database: %v", err))
		}
		err = db.AutoMigrate(&user.User{})
		if err != nil {
			panic(fmt.Sprintf("failed to migrate database: %v", err))
		}
	})

	return db
}
