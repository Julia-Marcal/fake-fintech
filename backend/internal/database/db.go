package database

import (
	"fmt"
	"os"
	"sync"

	acoes "github.com/Julia-Marcal/fake-fintech/internal/schemas/acoes"
	user "github.com/Julia-Marcal/fake-fintech/internal/schemas/user"
	wallet "github.com/Julia-Marcal/fake-fintech/internal/schemas/wallet"
	walletAcoes "github.com/Julia-Marcal/fake-fintech/internal/schemas/walletAcoes"
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
			panic(fmt.Sprintf("failed ]o connect to database: %v", err))
		}

		schemas := []interface{}{
			&user.User{},
			&wallet.Wallet{},
			&walletAcoes.WalletAcoes{},
			&acoes.Acoes{},
		}
		for _, schema := range schemas {
			if err := db.AutoMigrate(schema); err != nil {
				panic(fmt.Sprintf("failed to auto-migrate database: %v", err))

			}
		}

		if err != nil {
			panic(fmt.Sprintf("failed to migrate database: %v", err))
		}
	})

	return db
}
