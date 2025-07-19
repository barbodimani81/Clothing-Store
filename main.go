package main

import (
	"store/config"
	"store/models"
	"store/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config.ConnectDB()
	config.DB.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.CartItem{},
		&models.Order{},
		&models.OrderItem{},
	)

	routes.AuthRoutes(r)
	routes.ProductRoutes(r)
	routes.CartRoutes(r)
	routes.OrderRoutes(r)

	r.Run(":8080")
}
