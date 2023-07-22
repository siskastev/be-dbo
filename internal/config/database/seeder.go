package database

import (
	"test-be-dbo/internal/models"

	"gorm.io/gorm"
)

func seeder(db *gorm.DB) {
	//migrate schema
	db.AutoMigrate(&models.User{})

}
