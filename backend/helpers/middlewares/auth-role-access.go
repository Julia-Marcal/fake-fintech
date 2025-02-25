package handlers

import (
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type JWTClaim struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func RoleBasedAccess(requiredRole string, requiredEmail string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		tokenString := strings.Split(authHeader, " ")[1]
		claims := &JWTClaim{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("jwtKey")), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		if requiredRole != "" {
			if claims.Role != requiredRole {
				c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to access this resource"})
				c.Abort()
				return
			}
		}

		if requiredEmail != "" {
			if claims.Role != requiredEmail {
				c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to access this resource"})
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
