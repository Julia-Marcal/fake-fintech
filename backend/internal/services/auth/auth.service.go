package auth_service

import (
	"net/http"

	auth "github.com/Julia-Marcal/fake-fintech/helpers/auth"
	security "github.com/Julia-Marcal/fake-fintech/helpers/security"
	validation "github.com/Julia-Marcal/fake-fintech/helpers/validation"
	queries "github.com/Julia-Marcal/fake-fintech/internal/schemas/user/queries"
	"github.com/gin-gonic/gin"
)

func GenerateToken(email, password string) (map[string]interface{}, int) {
	validated := validation.EmailPassValidator(validation.EmailPassStruct{
		Email:    email,
		Password: password,
	})

	if validated {
		return gin.H{"error": "Invalid request body"}, http.StatusBadRequest
	}

	storedPassword, err := queries.CheckPassword(email)
	if err != nil {
		return gin.H{"error": err.Error()}, http.StatusInternalServerError
	}

	errPass := security.LoginSystem(password, storedPassword)
	if errPass != nil {
		return gin.H{"error": "Invalid credentials"}, http.StatusUnauthorized
	}

	userInfo, err := queries.FindUserByEmail(email)
	if err != nil {
		return gin.H{"error": err.Error()}, http.StatusInternalServerError
	}

	tokenString, err := auth.GenerateJWT(userInfo.Id, userInfo.Name, userInfo.Role, userInfo.Email)
	if err != nil {
		return gin.H{"error": err.Error()}, http.StatusInternalServerError
	}

	return gin.H{"token": tokenString}, http.StatusOK
}
