package wallet_controller

import (
	"net/http"

	"github.com/Julia-Marcal/fake-fintech/helpers/validation"
	database "github.com/Julia-Marcal/fake-fintech/internal/schemas/wallet"
	queries "github.com/Julia-Marcal/fake-fintech/internal/schemas/wallet/queries"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var wallet database.Wallet

	validated := validation.WalletValidator(wallet)

	if err := c.ShouldBindJSON(&wallet); err != nil || !validated {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input data",
		})
		return
	}

	err := queries.Create(&wallet)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create wallet",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "wallet created successfully",
		"user":    wallet,
	})
}
