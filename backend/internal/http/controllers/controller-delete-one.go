package controllers

import (
	"net/http"

	queries "github.com/Julia-Marcal/reusable-api/internal/user/queries"
	"github.com/gin-gonic/gin"
)

func DeleteUser(c *gin.Context) {
	userId, exists := c.Params.Get("id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User ID is required",
		})
		return
	}

	result := queries.DeleteOne(userId)

	if result == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})
}
