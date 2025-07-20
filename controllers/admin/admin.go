package admin

import (
	"net/http"
	"time"

	"store/config"
	"store/models"

	// "store/seeders"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}
	c.JSON(http.StatusOK, users)
}

func GetAllOrders(c *gin.Context) {
	var orders []models.Order
	if err := config.DB.Preload("Items").Preload("Items.Product").Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
		return
	}
	c.JSON(http.StatusOK, orders)
}

func DashboardStats(c *gin.Context) {
	var orderCount int64
	var totalSales float64

	config.DB.Model(&models.Order{}).Count(&orderCount)
	config.DB.Model(&models.Order{}).Select("SUM(total)").Scan(&totalSales)

	c.JSON(http.StatusOK, gin.H{
		"total_orders": orderCount,
		"total_sales":  totalSales,
		"timestamp":    time.Now(),
	})
}

func CreateProduct(c *gin.Context) {
	var input models.Product

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Product created", "product": input})
}

func UpdateProduct(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	var input models.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product.Name = input.Name
	product.Description = input.Description
	product.Price = input.Price

	if err := config.DB.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated", "product": product})
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	if err := config.DB.Delete(&models.Product{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}

// func SeedData(c *gin.Context) {
// 	if err := seeders.SeedDatabase(); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Seeding failed"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"message": "Database seeded successfully"})
// }
