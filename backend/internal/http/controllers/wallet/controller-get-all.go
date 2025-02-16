package wallet_controller

import (
	"fmt"
	"net/http"

	queries "github.com/Julia-Marcal/fake-fintech/internal/schemas/wallet/queries"
	"github.com/gin-gonic/gin"
)

func GetAllWallets(c *gin.Context) {
	userId, exists := c.Params.Get("id")
	fmt.Printf("%s", userId)

	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User ID is required",
		})
		return
	}

	wallets, err := queries.FindWallets(userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Wallets returned successfully",
		"users":   wallets,
	})
}
