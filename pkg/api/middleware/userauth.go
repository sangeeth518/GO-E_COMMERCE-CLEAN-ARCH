package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func UserAuth(c *gin.Context) {
	fmt.Println("in niddleware")
	token := c.GetHeader("UserAuhorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization token"})
		c.Abort()
		return
	} else {
		token, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			return []byte("usersecret"), nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
			c.Abort()
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
			c.Abort()
			return

		}
		fmt.Println("claims", claims)
		role, ok := claims["role"].(string)
		if !ok || role != "user" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized access"})
			c.Abort()
			return
		}

		id, ok := claims["id"].(float64)
		if !ok || id == 0 {
			c.JSON(http.StatusForbidden, gin.H{"error": "error in retrieving id"})
			c.Abort()
			return
		}

		c.Set("role", role)
		c.Set("id", int(id))

		c.Next()

	}
}
