package controllers

import (
	"net/http"
	"store/config"
	"store/models"
	"time"

	"github.com/gin-gonic/gin"
)

func Checkout(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	var cartItems []models.CartItem

	// Load user's cart with products
	config.DB.Preload("Product").Where("user_id = ?", user.ID).Find(&cartItems)
	if len(cartItems) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cart is empty"})
		return
	}

	// Prepare order items and calculate total
	var orderItems []models.OrderItem
	var total float64 = 0

	for _, item := range cartItems {
		subtotal := float64(item.Quantity) * item.Product.Price
		total += subtotal

		orderItems = append(orderItems, models.OrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     item.Product.Price,
		})
	}

	// Create order
	order := models.Order{
		UserID:    user.ID,
		Items:     orderItems,
		Total:     total,
		CreatedAt: time.Now().Unix(),
	}

	if err := config.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create order"})
		return
	}

	// Clear the cart
	config.DB.Where("user_id = ?", user.ID).Delete(&models.CartItem{})

	c.JSON(http.StatusOK, gin.H{"message": "Order placed successfully", "order_id": order.ID})
}

func GetOrders(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	var orders []models.Order
	err := config.DB.
		Preload("Items.Product").
		Where("user_id = ?", user.ID).
		Order("created_at desc").
		Find(&orders).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
		return
	}

	c.JSON(http.StatusOK, orders)
}
