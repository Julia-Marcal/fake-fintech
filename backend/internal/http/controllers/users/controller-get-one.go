package user_controller

import (
	"net/http"

	cache "github.com/Julia-Marcal/fake-fintech/internal/cache/caching-func/user"
	queries "github.com/Julia-Marcal/fake-fintech/internal/schemas/user/queries"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	userId := c.Param("id_user")

	cachedUser, _ := cache.GetCachedUser(userId)

	if cachedUser.Id != "" {
		c.JSON(http.StatusOK, gin.H{
			"message": "User retrieved from cache",
			"user":    cachedUser,
		})
		return
	}

	user, err := queries.FindUserById(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}

	cacheErr := cache.CacheUser(*user)
	if cacheErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to cache user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User returned successfully",
		"user":    user,
	})
}
