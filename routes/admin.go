package routes

import (
	"store/controllers/admin"
	"store/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterAdminRoutes(r *gin.Engine) {
	adminRoutes := r.Group("/admin")
	adminRoutes.Use(middlewares.AuthMiddleware(), middlewares.IsAdmin())

	adminRoutes.GET("/users", admin.GetAllUsers)
	adminRoutes.GET("/orders", admin.GetAllOrders)
	adminRoutes.POST("/products", admin.CreateProduct)
	adminRoutes.PUT("/products/:id", admin.UpdateProduct)
	adminRoutes.DELETE("/products/:id", admin.DeleteProduct)
	adminRoutes.GET("/dashboard", admin.DashboardStats)
}
