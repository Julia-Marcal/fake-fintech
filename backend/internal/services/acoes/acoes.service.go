package acoes_service

import (
	"fmt"
	"net/http"

	cache "github.com/Julia-Marcal/fake-fintech/internal/cache/caching-func/acoes"
	database_acoes "github.com/Julia-Marcal/fake-fintech/internal/schemas/acoes"
	queries "github.com/Julia-Marcal/fake-fintech/internal/schemas/acoes/queries"
	database_wallet_acoes "github.com/Julia-Marcal/fake-fintech/internal/schemas/walletAcoes"
	wallet_acoes_queries "github.com/Julia-Marcal/fake-fintech/internal/schemas/walletAcoes/queries"
	"github.com/gin-gonic/gin"
)

func GetAcoes(acaoId string) (map[string]interface{}, int) {
	cachedAcoes, cacheErr := cache.GetCachedAcoes(acaoId)
	if cacheErr != nil {
		fmt.Println("Cache retrieval error:", cacheErr)
	}

	if cachedAcoes.Id != "" {
		return gin.H{
			"message": "Investment allocation retrieved from cache",
			"acoes":   cachedAcoes,
		}, http.StatusOK
	}

	acao, err := queries.FindAcao(acaoId)
	if err != nil {
		return gin.H{"error": err.Error()}, http.StatusInternalServerError
	}

	if acao == nil {
		return gin.H{"error": "Investment allocation not found"}, http.StatusNotFound
	}

	cacheErr = cache.CacheAcoes(*acao)
	if cacheErr != nil {
		return gin.H{"error": "Failed to cache Investment allocation"}, http.StatusInternalServerError
	}

	return gin.H{
		"message":     "Investment allocation returned successfully",
		"cachedAcoes": acao,
	}, http.StatusOK
}

func CreateAcoes(walletAcao *database_wallet_acoes.WalletAcoes, acao database_acoes.Acoes) (map[string]interface{}, int) {
	err := wallet_acoes_queries.Create(walletAcao)
	if err != nil {
		return gin.H{"error": err.Error()}, http.StatusInternalServerError
	}

	cacheErr := cache.CacheAcoes(acao)
	if cacheErr != nil {
		fmt.Println("Failed to cache newly created acao:", cacheErr)
	}

	return gin.H{
		"message": "Investment allocation created successfully",
		"acoes":   acao,
	}, http.StatusCreated
}
