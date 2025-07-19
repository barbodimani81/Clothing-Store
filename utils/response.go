package utils

import "github.com/gin-gonic/gin"

// Standard error response
func Error(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{
		"error": message,
	})
}

// Standard success response
func Success(c *gin.Context, status int, data any) {
	c.JSON(status, gin.H{
		"data": data,
	})
}
