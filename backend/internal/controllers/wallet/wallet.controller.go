package controller

import (
	"net/http"

	"github.com/Julia-Marcal/fake-fintech/internal/schemas/wallet"
	service "github.com/Julia-Marcal/fake-fintech/internal/services/wallet"
	"github.com/gin-gonic/gin"
)

func GetAllWallets(c *gin.Context) {
	userId, exists := c.Params.Get("id_user")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	wallets, err := service.GetAllWalletsService(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Wallets returned successfully",
		"users":   wallets,
	})
}

func CreateWallet(c *gin.Context) {
	var wallet wallet.Wallet

	if err := c.ShouldBindJSON(&wallet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	createdWallet, err := service.CreateWalletService(wallet)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "wallet created successfully",
		"user":    createdWallet,
	})
}

func GetWallet(c *gin.Context) {
	walletId, exists := c.Params.Get("id_wallet")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wallet ID is required"})
		return
	}

	result, err := service.GetWalletService(walletId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if result == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Wallet not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Wallet returned successfully",
		"wallet":  result,
	})
}
