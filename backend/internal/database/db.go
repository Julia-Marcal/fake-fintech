package repository

import (
	"fmt"

	"sync"

	database "github.com/Julia-Marcal/reusable-api/internal/user"
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

		connectionStr := "user=postgres password=password dbname=api_db host=postgres port=5432 sslmode=disable"
		fmt.Println("about to connect to database")

		db, err = gorm.Open(postgres.Open(connectionStr), &gorm.Config{
			SkipDefaultTransaction: true,
		})
		if err != nil {
			panic(fmt.Sprintf("failed to connect to database: %v", err))
		}
		err = db.AutoMigrate(&database.User{})
		if err != nil {
			panic(fmt.Sprintf("failed to migrate database: %v", err))
		}
	})

	return db
}
