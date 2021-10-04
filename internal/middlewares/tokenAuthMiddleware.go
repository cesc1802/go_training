package middlewares

import (
	"net/http"
	"os"

	"github.com/cesc1802/go_training/internal/services"
	"github.com/cesc1802/go_training/internal/storages"
	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.String(http.StatusForbidden, "Empty Authorization header")
			c.Abort()
			return
		}

		jwtWrapper := services.JwtWrapper{
			SecretKey: os.Getenv("SECRET"),
		}

		userId, err := jwtWrapper.ValidateToken(token)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorize")
			c.Abort()
			return
		}

		var user storages.User
		storages.Get().Where("id = ?", userId).First(&user)
		if user.ID == "" {
			c.String(http.StatusForbidden, "User not found")
			c.Abort()
			return
		}

		c.Set("USER", user)
		c.Next()
	}
}
