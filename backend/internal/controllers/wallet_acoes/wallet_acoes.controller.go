package wallet_acoes_controller

import (
	"net/http"

	queries "github.com/Julia-Marcal/fake-fintech/internal/schemas/walletAcoes/queries"
	"github.com/gin-gonic/gin"
)

func GetWalletAcoes(c *gin.Context) {
	walletId, exists := c.Params.Get("id_wallet")

	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Wallet ID is required",
		})
		return
	}

	walletAcoes, err := queries.FindAcoesByWallet(walletId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if walletAcoes == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Wallet not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Wallet with investment allocation was returned successfully",
		"acoes":   walletAcoes,
	})
}
