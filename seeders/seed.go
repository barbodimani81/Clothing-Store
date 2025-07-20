package seeders

import (
	"fmt"
	"store/config"
	"store/models"

	"golang.org/x/crypto/bcrypt"
)

func Seed() {
	// Sample users
	password, _ := bcrypt.GenerateFromPassword([]byte("password123"), 14)

	users := []models.User{
		{Name: "Ali", Email: "ali@ali.com", Password: string(password), Role: "user"},
		{Name: "qasem", Email: "qasem@qasem.com", Password: string(password), Role: "user"},
	}

	for _, user := range users {
		config.DB.Where("email = ?", user.Email).FirstOrCreate(&user)
	}

	// Sample products
	products := []models.Product{
		{Name: "T-Shirt", Description: "Oversize White", Price: 500},
		{Name: "Hoodie", Description: "Warm and comfy", Price: 900},
		{Name: "Sneakers", Description: "Running shoes", Price: 1400},
	}

	for _, product := range products {
		config.DB.Where("name = ?", product.Name).FirstOrCreate(&product)
	}

	fmt.Println("✅ Seeding completed.")
}
