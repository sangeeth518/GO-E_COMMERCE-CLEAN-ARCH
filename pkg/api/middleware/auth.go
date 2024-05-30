package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func AdminAuthMiddleware(c *gin.Context) {
	token, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(401)
	}

	token = strings.TrimPrefix(token, "Bearer")

	// token, err = jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
	// 	return []byte(config.Config.JWTToken), nil
	// })
}
