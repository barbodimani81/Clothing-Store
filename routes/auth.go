package routes

import (
	"store/controllers"
	"store/middlewares"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	protected := r.Group("/")
	protected.Use(middlewares.AuthMiddleware())
	{
		protected.GET("/me", func(c *gin.Context) {
			user, _ := c.Get("user")
			c.JSON(200, user)
		})
	}
}
