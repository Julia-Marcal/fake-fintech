package acoes_controller

import (
	"net/http"

	database_acoes "github.com/Julia-Marcal/fake-fintech/internal/schemas/acoes"
	database_wallet_acoes "github.com/Julia-Marcal/fake-fintech/internal/schemas/walletAcoes"
	acoes_service "github.com/Julia-Marcal/fake-fintech/internal/services/acoes"
	"github.com/gin-gonic/gin"
)

type CreateAcoesRequest struct {
	WalletId string               `json:"wallet_id" binding:"required,uuid4"`
	Acoes    database_acoes.Acoes `json:"acoes" binding:"required"`
}

func GetAcoes(c *gin.Context) {
	AcaoId, exists := c.Params.Get("id_acao")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Investment allocation ID is required"})
		return
	}

	response, status := acoes_service.GetAcoes(AcaoId)
	c.JSON(status, response)
}

func CreateAcoes(c *gin.Context) {
	var req CreateAcoesRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	walletAcao := database_wallet_acoes.WalletAcoes{
		WalletId: req.WalletId,
		AcoesId:  req.Acoes.Id,
	}

	response, status := acoes_service.CreateAcoes(&walletAcao, req.Acoes)
	c.JSON(status, response)
}
