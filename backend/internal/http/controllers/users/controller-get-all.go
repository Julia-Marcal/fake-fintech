package user_controller

import (
	"net/http"
	"strconv"

	queries "github.com/Julia-Marcal/fake-fintech/internal/schemas/user/queries"
	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	limit_str := c.DefaultQuery("limit", "10")

	limit, err := strconv.Atoi(limit_str)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid limit value",
		})
		return
	}

	if limit > 100 {
		limit = 100
	}

	users, err := queries.FindUsers(limit)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "users returned successfully",
		"users":   users,
	})
}
