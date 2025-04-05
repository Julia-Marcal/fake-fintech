package user_controller

import (
	user_service "github.com/Julia-Marcal/fake-fintech/internal/services/users"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	response := user_service.CreateUser(c)
	c.JSON(response.Status, response.Body)
}

func DeleteUser(c *gin.Context) {
	response := user_service.DeleteUser(c)
	c.JSON(response.Status, response.Body)
}

func GetUser(c *gin.Context) {
	response := user_service.GetUser(c)
	c.JSON(response.Status, response.Body)
}

func GetAllUsers(c *gin.Context) {
	response := user_service.GetAllUsers(c)
	c.JSON(response.Status, response.Body)
}

func UpdateUser(c *gin.Context) {
	response := user_service.UpdateUser(c)
	c.JSON(response.Status, response.Body)
}

func TotalAmount(c *gin.Context) {
	response := user_service.TotalAmount(c)
	c.JSON(response.Status, response.Body)
}
