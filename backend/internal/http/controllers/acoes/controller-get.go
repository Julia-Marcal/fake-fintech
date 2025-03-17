package acoes_controller

import (
	"fmt"
	"net/http"

	cache "github.com/Julia-Marcal/fake-fintech/internal/cache/caching-func/acoes"
	queries "github.com/Julia-Marcal/fake-fintech/internal/schemas/acoes/queries"
	"github.com/gin-gonic/gin"
)

func GetAcoes(c *gin.Context) {
	AcaoId, exists := c.Params.Get("id_acao")

	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Investment allocation ID is required",
		})
		return
	}

	cachedAcoes, cacheErr := cache.GetCachedAcoes(AcaoId)
	if cacheErr != nil {
		fmt.Println("Cache retrieval error:", cacheErr)
	}

	if cachedAcoes.Id != "" {
		c.JSON(http.StatusOK, gin.H{
			"message": "Investment allocation retrieved from cache",
			"acoes":   cachedAcoes,
		})
		return
	}

	acao, err := queries.FindAcao(AcaoId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if acao == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Investment allocation not found",
		})
		return
	}

	cacheErr = cache.CacheAcoes(*acao)
	if cacheErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to cache Investment allocation",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "Investment allocation returned successfully",
		"cachedAcoes": acao,
	})
}
