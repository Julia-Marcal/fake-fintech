package user_controller

import (
	"net/http"

	queries "github.com/Julia-Marcal/fake-fintech/internal/schemas/user/queries"
	"github.com/gin-gonic/gin"
)

func TotalAmount(c *gin.Context) {
	userId, exists := c.Params.Get("id_user")

	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User ID is required",
		})
		return
	}

	amount, err := queries.TotalAmountByUser(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Amount returned successfully",
		"amount":  amount,
	})
}
