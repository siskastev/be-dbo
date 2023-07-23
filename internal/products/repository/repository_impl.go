package repository

import (
	"test-be-dbo/internal/models"

	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) Repository {
	return &productRepository{db: db}
}

func (p *productRepository) GetAll() ([]models.Product, error) {
	var products []models.Product
	if err := p.db.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}
