package routes

import (
	"store/controllers"
	"store/middlewares"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.Engine) {
	r.GET("/products", controllers.GetProducts)
	r.GET("/products/:id", controllers.GetProduct)

	admin := r.Group("/products")
	admin.Use(middlewares.AuthMiddleware(), middlewares.AdminOnly())
	{
		admin.POST("", controllers.CreateProduct)
	}
}
