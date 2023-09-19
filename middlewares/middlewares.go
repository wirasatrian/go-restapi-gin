package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wirasatrian/go-restapi-gin/utils/token"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}
		c.Next()
	}
}
