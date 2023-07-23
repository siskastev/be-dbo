package repository

import (
	"test-be-dbo/internal/models"
)

type Repository interface {
	GetAll() ([]models.Product, error)
}
