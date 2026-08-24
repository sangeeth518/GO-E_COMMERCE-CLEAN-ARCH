package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/sangeeth518/go-Ecommerce/pkg/config"
)

func AdminAuthMiddleware(cfg config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {

		token, err := c.Cookie("Authorization")
		if err != nil {
			c.AbortWithStatus(401)
			return
		}

		token = strings.TrimPrefix(token, "Bearer")

		if token == "" {
			c.AbortWithStatus(401)
			return
		} else {
			token, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(cfg.JWTToken), nil
			})
			if err != nil {
				c.AbortWithStatus(401)
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok || !token.Valid {
				c.JSON(401, gin.H{"error": "Invalid token claims"})
				c.Abort()
				return
			}

			//Check

			role, ok := claims["role"].(string)
			if !ok || role != "admin" {
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Unauthorized access: admin role requiired"})
				return
			}
			id, exists := claims["id"]
			idfloat, ok := id.(float64)
			if !ok || !exists || int(idfloat) == 0 {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid admin ID in token"})
				return
			}

			// set user data to gin context
			c.Set("adminId", int(idfloat))
			c.Set("role", role)
			c.Set("email", claims["email"])
			c.Next()

		}

	}
}
