package queries_wallet

import (
	repository "github.com/Julia-Marcal/fake-fintech/internal/database"
	database "github.com/Julia-Marcal/fake-fintech/internal/schemas/wallet"
	_ "github.com/gin-gonic/gin"
)

func Create(wallet_info *database.Wallet) error {
	db := repository.NewPostgres()
	result := db.Create(wallet_info)
	return result.Error
}