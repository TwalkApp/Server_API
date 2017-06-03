package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/twalkapp/server/misc/jwt"
	"strings"
)

func CheckAuth() gin.HandlerFunc {
	return func(c* gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if len(token) == 0{
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			c.Abort()
		}
		if !strings.HasPrefix(token, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token type"})
			c.Abort()
		}
		status, _, _ := jwt.ValidateToken(strings.Replace(token, "Bearer ", "", -1))
		if status {
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
		}
	}
}