package queries_wallet_acoes

import (
	repository "github.com/Julia-Marcal/fake-fintech/internal/database"
	database_acoes "github.com/Julia-Marcal/fake-fintech/internal/schemas/acoes"
	database "github.com/Julia-Marcal/fake-fintech/internal/schemas/walletAcoes"
	_ "github.com/gin-gonic/gin"
)

func Create(walletAcao *database.WalletAcoes) error {
	db := repository.NewPostgres()
	result := db.Create(walletAcao)
	return result.Error
}

func FindAcoesByWallet(walletId string) ([]database_acoes.Acoes, error) {
	db := repository.NewPostgres()
	var acoes []database_acoes.Acoes

	result := db.Table("wallet_acoes").
		Select("acoes.*").
		Joins("JOIN acoes ON wallet_acoes.acoes_id = acoes.id").
		Where("wallet_acoes.wallet_id = ?", walletId).
		Scan(&acoes)

	if result.Error != nil {
		return nil, result.Error
	}

	return acoes, nil
}
