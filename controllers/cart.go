package controllers

import (
	"net/http"
	"store/config"
	"store/models"

	"github.com/gin-gonic/gin"
)

func GetCart(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	var cart []models.CartItem

	config.DB.Preload("Product").Where("user_id = ?", user.ID).Find(&cart)
	c.JSON(http.StatusOK, cart)
}

func AddToCart(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	var input struct {
		ProductID uint `json:"product_id"`
		Quantity  int  `json:"quantity"`
	}

	if err := c.ShouldBindJSON(&input); err != nil || input.Quantity < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var product models.Product
	if err := config.DB.First(&product, input.ProductID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	var cartItem models.CartItem
	db := config.DB.Where("user_id = ? AND product_id = ?", user.ID, input.ProductID).First(&cartItem)

	if db.RowsAffected > 0 {
		cartItem.Quantity += input.Quantity
		config.DB.Save(&cartItem)
	} else {
		newItem := models.CartItem{
			UserID:    user.ID,
			ProductID: input.ProductID,
			Quantity:  input.Quantity,
		}
		config.DB.Create(&newItem)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product added to cart"})
}

func RemoveFromCart(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	id := c.Param("id")

	var item models.CartItem
	if err := config.DB.First(&item, id).Error; err != nil || item.UserID != user.ID {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	config.DB.Delete(&item)
	c.JSON(http.StatusOK, gin.H{"message": "Item removed from cart"})
}
