package user_controller

import (
	"net/http"

	validation "github.com/Julia-Marcal/fake-fintech/helpers/validation"
	cache "github.com/Julia-Marcal/fake-fintech/internal/cache/caching-func/user"
	database "github.com/Julia-Marcal/fake-fintech/internal/schemas/user"
	queries "github.com/Julia-Marcal/fake-fintech/internal/schemas/user/queries"
	"github.com/gin-gonic/gin"
)

func UpdateUser(c *gin.Context) {
	userId := c.Param("id_user")
	if userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User ID is required",
		})
		return
	}

	existingUser, err := queries.FindUserByEmail(userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}

	var updatedUser database.User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input data",
		})
		return
	}

	if !validation.UserValidator(updatedUser) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user data",
		})
		return
	}

	existingUser.Name = updatedUser.Name
	existingUser.LastName = updatedUser.LastName
	existingUser.Age = updatedUser.Age
	existingUser.Email = updatedUser.Email
	existingUser.Password = updatedUser.Password
	existingUser.Role = updatedUser.Role

	CacheErr := cache.CacheUser(*existingUser)
	if CacheErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to cache user",
		})
		return
	}

	err = queries.Update(existingUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully",
		"user":    existingUser,
	})
}
