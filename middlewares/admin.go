package middlewares

import (
	"net/http"
	"store/models"

	"github.com/gin-gonic/gin"
)

func IsAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, ok := c.MustGet("user").(models.User)
		if !ok || user.Role != "admin" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Admin access only"})
			return
		}
		c.Next()
	}
}
