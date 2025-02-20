package wallet_controller

import (
	"net/http"

	"github.com/Julia-Marcal/fake-fintech/helpers/validation"
	database_acoes "github.com/Julia-Marcal/fake-fintech/internal/schemas/acoes"
	acoes_queries "github.com/Julia-Marcal/fake-fintech/internal/schemas/acoes/queries"
	database_wallet_acoes "github.com/Julia-Marcal/fake-fintech/internal/schemas/walletAcoes"
	wallet_acoes_queries "github.com/Julia-Marcal/fake-fintech/internal/schemas/walletAcoes/queries"
	"github.com/gin-gonic/gin"
)

type CreateAcoesRequest struct {
	WalletId string               `json:"wallet_id" binding:"required,uuid4"`
	Acoes    database_acoes.Acoes `json:"acoes" binding:"required"`
}

func CreateAcoes(c *gin.Context) {
	var req CreateAcoesRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request payload",
		})
		return
	}

	if !validation.AcoesWalletValidator(req.Acoes) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Validation failed for Acoes",
		})
		return
	}

	err := acoes_queries.Create(&req.Acoes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create Acoes",
		})
		return
	}

	walletAcoes := database_wallet_acoes.WalletAcoes{
		WalletId: req.WalletId,
		AcoesId:  req.Acoes.Id,
	}

	err = wallet_acoes_queries.Create(&walletAcoes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to link Acoes to wallet",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Investment allocation created and linked to wallet successfully",
		"acoes":   req.Acoes,
	})
}
