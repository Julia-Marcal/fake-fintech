package queries_acoes

import (
	repository "github.com/Julia-Marcal/fake-fintech/internal/database"
	database "github.com/Julia-Marcal/fake-fintech/internal/schemas/acoes"
	_ "github.com/gin-gonic/gin"
)

func Create(acao *database.Acoes) error {
	db := repository.NewPostgres()
	result := db.Create(acao)
	return result.Error
}

func FindAcao(id_acao string) (*database.Acoes, error) {
	db := repository.NewPostgres()
	acao := &database.Acoes{}
	result := db.First(acao, "id = ?", id_acao).Limit(1)
	return acao, result.Error
}
