package wallet_controller

import (
	"net/http"

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

	c.JSON(http.StatusOK, gin.H{
		"message": "Wallet returned successfully",
		"user":    wallet,
	})
}
