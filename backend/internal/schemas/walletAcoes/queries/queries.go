package queries_wallet_acoes

import (
	repository "github.com/Julia-Marcal/fake-fintech/internal/database"
	database "github.com/Julia-Marcal/fake-fintech/internal/schemas/walletAcoes"
	_ "github.com/gin-gonic/gin"
)

func Create(walletAcao *database.WalletAcoes) error {
	db := repository.NewPostgres()
	result := db.Create(walletAcao)
	return result.Error
}

func FindAcoesByWallet(walletId string) ([]database.WalletAcoes, error) {
	db := repository.NewPostgres()
	var walletAcoes []database.WalletAcoes

	result := db.Where("wallet_id = ?", walletId).Find(&walletAcoes)

	return walletAcoes, result.Error
}
