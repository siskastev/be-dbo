package database

import (
	"fmt"
	"math/rand"
	"test-be-dbo/internal/models"
	"time"

	"gorm.io/gorm"
)

func seeder(db *gorm.DB) {
	//migrate schema
	db.AutoMigrate(&models.User{}, &models.Customer{}, &models.Product{})

	products := make([]models.Product, 0)

	for i := 0; i < 15; i++ {
		product := models.Product{
			Name:      fmt.Sprintf("Product %d", i+1),
			Qty:       uint16(rand.Intn(100)),          // Random quantity between 0 and 99
			Price:     float64(rand.Intn(1000)) * 1000, // Random price between 0 and 99.9
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		products = append(products, product)
	}

	if err := db.Create(&products).Error; err != nil {
		fmt.Println("Error creating products:", err)
		return
	}

}
