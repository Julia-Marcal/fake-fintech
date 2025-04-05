package user_controller

import (
	"net/http"

	auth_service "github.com/Julia-Marcal/fake-fintech/internal/services/auth"
	"github.com/gin-gonic/gin"
)

type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GenerateToken(c *gin.Context) {
	var request TokenRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, status := auth_service.GenerateToken(request.Email, request.Password)
	c.JSON(status, response)
}
