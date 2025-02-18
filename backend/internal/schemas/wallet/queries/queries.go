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

func FindWallet(id_wallet string) (*database.Wallet, error) {
	db := repository.NewPostgres()
	wallet := &database.Wallet{}
	result := db.First(wallet, "id = ?", id_wallet).Limit(1)
	return wallet, result.Error
}

func FindWallets(id_user string) ([]database.Wallet, error) {
	db := repository.NewPostgres()
	var wallets []database.Wallet

	result := db.Where("user_id = ?", id_user).Find(&wallets)

	return wallets, result.Error
}
