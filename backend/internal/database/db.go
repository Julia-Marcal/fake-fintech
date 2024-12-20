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
		connectionStr := "user=postgres password=password dbname=api_db host=postgres port=5432 sslmode=disable"
		fmt.Println("about to connect to database")
		db, err := gorm.Open(postgres.Open(connectionStr), &gorm.Config{
			SkipDefaultTransaction: true,
			//disable operations inside transaction to ensure data consistency
		})
		if err != nil {
			panic("failed to connect to database")
		}
		db.AutoMigrate(&database.User{})
	})
	return db
}
