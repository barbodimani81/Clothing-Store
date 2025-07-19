package routes

import (
	"store/controllers"
	"store/middlewares"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(r *gin.Engine) {
	order := r.Group("/order")
	order.Use(middlewares.AuthMiddleware())
	{
		order.POST("/checkout", controllers.Checkout)
		order.GET("/history", controllers.GetOrders)
	}
}
