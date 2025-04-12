package user_service

import (
	"net/http"
	"strconv"

	validation "github.com/Julia-Marcal/fake-fintech/helpers/validation"
	cache "github.com/Julia-Marcal/fake-fintech/internal/cache/caching-func/user"
	database "github.com/Julia-Marcal/fake-fintech/internal/schemas/user"
	queries "github.com/Julia-Marcal/fake-fintech/internal/schemas/user/queries"
	"github.com/gin-gonic/gin"
)

type ServiceResponse struct {
	Status int
	Body   gin.H
}

func CreateUser(c *gin.Context) ServiceResponse {
	var user database.User

	if err := c.ShouldBindJSON(&user); err != nil || !validation.UserValidator(user) {
		return ServiceResponse{http.StatusBadRequest, gin.H{"error": "Invalid input data"}}
	}

	if user.Role == "" {
		user.Role = "user"
	}

	if err := cache.CacheUser(user); err != nil {
		return ServiceResponse{http.StatusInternalServerError, gin.H{"error": "Failed to cache user"}}
	}

	if err := queries.Create(&user); err != nil {
		return ServiceResponse{http.StatusInternalServerError, gin.H{"error": "Failed to create user"}}
	}

	return ServiceResponse{http.StatusOK, gin.H{"message": "User created successfully", "user": user}}
}

func DeleteUser(c *gin.Context) ServiceResponse {
	userId := c.Param("id")
	if userId == "" {
		return ServiceResponse{http.StatusBadRequest, gin.H{"error": "User ID is required"}}
	}

	result := queries.DeleteOne(userId)
	if result == nil {
		return ServiceResponse{http.StatusNotFound, gin.H{"error": "User not found"}}
	}

	return ServiceResponse{http.StatusOK, gin.H{"message": "User deleted successfully"}}
}

func GetUser(c *gin.Context) ServiceResponse {
	userId := c.Param("id_user")
	cachedUser, _ := cache.GetCachedUser(userId)

	if cachedUser.Id != "" {
		return ServiceResponse{http.StatusOK, gin.H{"message": "User retrieved from cache", "user": cachedUser}}
	}

	user, err := queries.FindUserById(userId)
	if err != nil {
		return ServiceResponse{http.StatusInternalServerError, gin.H{"error": err.Error()}}
	}

	if user == nil {
		return ServiceResponse{http.StatusNotFound, gin.H{"error": "User not found"}}
	}

	if err := cache.CacheUser(*user); err != nil {
		return ServiceResponse{http.StatusInternalServerError, gin.H{"error": "Failed to cache user"}}
	}

	return ServiceResponse{http.StatusOK, gin.H{"message": "User returned successfully", "user": user}}
}

func GetAllUsers(c *gin.Context) ServiceResponse {
	limitStr := c.DefaultQuery("limit", "10")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		return ServiceResponse{http.StatusBadRequest, gin.H{"error": "invalid limit value"}}
	}

	if limit > 100 {
		limit = 100
	}

	users, err := queries.FindUsers(limit)
	if err != nil {
		return ServiceResponse{http.StatusInternalServerError, gin.H{"error": err.Error()}}
	}

	return ServiceResponse{http.StatusOK, gin.H{"message": "users returned successfully", "users": users}}
}

func UpdateUser(c *gin.Context) ServiceResponse {
	userId := c.Param("id_user")
	if userId == "" {
		return ServiceResponse{http.StatusBadRequest, gin.H{"error": "User ID is required"}}
	}

	existingUser, err := queries.FindUserByEmail(userId)
	if err != nil || existingUser == nil {
		return ServiceResponse{http.StatusNotFound, gin.H{"error": "User not found"}}
	}

	var updatedUser database.User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		return ServiceResponse{http.StatusBadRequest, gin.H{"error": "Invalid input data"}}
	}

	if !validation.UserValidator(updatedUser) {
		return ServiceResponse{http.StatusBadRequest, gin.H{"error": "Invalid user data"}}
	}

	existingUser.Name = updatedUser.Name
	existingUser.LastName = updatedUser.LastName
	existingUser.Age = updatedUser.Age
	existingUser.Email = updatedUser.Email
	existingUser.Password = updatedUser.Password
	existingUser.Role = updatedUser.Role

	if err := cache.CacheUser(*existingUser); err != nil {
		return ServiceResponse{http.StatusInternalServerError, gin.H{"error": "Failed to cache user"}}
	}

	if err := queries.Update(existingUser); err != nil {
		return ServiceResponse{http.StatusInternalServerError, gin.H{"error": "Failed to update user"}}
	}

	return ServiceResponse{http.StatusOK, gin.H{"message": "User updated successfully", "user": existingUser}}
}

func TotalAmount(c *gin.Context) ServiceResponse {
	userId := c.Param("id_user")
	if userId == "" {
		return ServiceResponse{http.StatusBadRequest, gin.H{"error": "User ID is required"}}
	}

	amount, err := queries.TotalAmountByUser(userId)
	if err != nil {
		return ServiceResponse{http.StatusInternalServerError, gin.H{"error": err.Error()}}
	}

	return ServiceResponse{http.StatusOK, gin.H{"message": "Amount returned successfully", "amount": amount}}
}
