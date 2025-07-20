package controllers

import (
	"net/http"
	"store/config"
	"store/models"

	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	search := c.Query("search")
	minPrice, _ := strconv.ParseFloat(c.DefaultQuery("min", "0"), 64)
	maxPrice, _ := strconv.ParseFloat(c.DefaultQuery("max", "1000000"), 64)

	var products []models.Product
	var total int64

	offset := (page - 1) * limit
	query := config.DB.Model(&models.Product{})

	if search != "" {
		query = query.Where("name ILIKE ? OR description ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	query = query.Where("price >= ? AND price <= ?", minPrice, maxPrice)

	query.Count(&total)

	err := query.Offset(offset).Limit(limit).Find(&products).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"products":    products,
		"total":       total,
		"page":        page,
		"per_page":    limit,
		"total_pages": (total + int64(limit) - 1) / int64(limit),
	})
}

func CreateProduct(c *gin.Context) {
	var input models.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := models.Product{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
	}

	config.DB.Create(&product)
	c.JSON(http.StatusCreated, product)
}

func UploadProductImage(c *gin.Context) {
	id := c.Param("id")

	var product models.Product
	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No image uploaded"})
		return
	}

	// Save the image to disk (you can change this path)
	filePath := fmt.Sprintf("uploads/%d_%s", product.ID, file.Filename)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
		return
	}

	// Update product's ImageURL
	product.ImageURL = "/" + filePath
	config.DB.Save(&product)

	c.JSON(http.StatusOK, gin.H{"message": "Image uploaded successfully", "image_url": product.ImageURL})
}
