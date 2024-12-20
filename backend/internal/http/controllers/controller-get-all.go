package controllers

import (
	"net/http"

	queries "github.com/Julia-Marcal/reusable-api/internal/user/queries"
	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	users, err := queries.FindUsers()

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
