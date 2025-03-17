package wallet_controller

import (
	"fmt"
	"net/http"

	cache "github.com/Julia-Marcal/fake-fintech/internal/cache/caching-func/wallet"
	queries "github.com/Julia-Marcal/fake-fintech/internal/schemas/wallet/queries"
	"github.com/gin-gonic/gin"
)

func GetWallet(c *gin.Context) {
	walletId, exists := c.Params.Get("id_wallet")

	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Wallet ID is required",
		})
		return
	}

	cachedWallet, cacheErr := cache.GetCachedWallet(walletId)
	if cacheErr != nil {
		fmt.Println("Cache retrieval error:", cacheErr)
	}

	if cachedWallet.Id != "" {
		c.JSON(http.StatusOK, gin.H{
			"message": "Wallet retrieved from cache",
			"wallet":  cachedWallet,
		})
		return
	}

	wallet, err := queries.FindWallet(walletId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if wallet == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Wallet not found",
		})
		return
	}

	cacheErr = cache.CacheWallet(*wallet)
	if cacheErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to cache wallet",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Wallet returned successfully",
		"wallet":  wallet,
	})
}
