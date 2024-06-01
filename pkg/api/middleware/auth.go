package middleware

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AdminAuthMiddleware(c *gin.Context) {
	token, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(401)
	}

	token = strings.TrimPrefix(token, "Bearer")

	if token == "" {
		c.Abort()
	} else {
		token, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte("abc1234"), nil
		})
		if err != nil {
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if id, exists := claims["ID"]; exists {
				if idfloat, ok := id.(float64); ok {
					userId := int(idfloat)
					if userId == 0 {
						c.Abort()
						return
					}
					c.Set("id", userId)
				}
			}

		}

	}

}
