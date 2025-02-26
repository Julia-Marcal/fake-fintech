package handlers

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type JWTClaim struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func RoleBasedAccess(requiredRole string, requiredId string) gin.HandlerFunc {
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

		log.Printf("Claims: %+v\n", claims)
		log.Printf("Required Role: %s\n", requiredRole)
		log.Printf("Required ID: %s\n", requiredId)

		if claims.Role != requiredRole {
			if claims.Id != requiredId {
				c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to access this resource"})
				c.Abort()
				return
			} else {
				c.JSON(http.StatusForbidden, gin.H{"error": "You do not have the required role to access this resource"})
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
