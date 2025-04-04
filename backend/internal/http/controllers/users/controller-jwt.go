package user_controller

import (
	"net/http"

	auth "github.com/Julia-Marcal/fake-fintech/helpers/auth"
	security "github.com/Julia-Marcal/fake-fintech/helpers/security"
	validation "github.com/Julia-Marcal/fake-fintech/helpers/validation"
	queries "github.com/Julia-Marcal/fake-fintech/internal/schemas/user/queries"
	"github.com/gin-gonic/gin"
)

type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GenerateToken(context *gin.Context) {
	var request TokenRequest

	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validated := validation.EmailPassValidator(validation.EmailPassStruct{
		Email:    request.Email,
		Password: request.Password,
	})

	if validated {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		context.Abort()
		return
	}

	pass, err := queries.CheckPassword(request.Email)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	errPass := security.LoginSystem(request.Password, pass)
	if errPass != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	userInfo, _ := queries.FindUserByEmail(request.Email)
	tokenString, err := auth.GenerateJWT(userInfo.Id, userInfo.Name, userInfo.Role, userInfo.Email)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, gin.H{"token": tokenString})
}
