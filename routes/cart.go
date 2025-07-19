package routes

import (
	"store/controllers"
	"store/middlewares"

	"github.com/gin-gonic/gin"
)

func CartRoutes(r *gin.Engine) {
	cart := r.Group("/cart")
	cart.Use(middlewares.AuthMiddleware())
	{
		cart.GET("", controllers.GetCart)
		cart.POST("/add", controllers.AddToCart)
		cart.DELETE("/remove/:id", controllers.RemoveFromCart)
	}
}
