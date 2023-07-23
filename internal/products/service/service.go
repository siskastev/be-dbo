package service

import (
	"test-be-dbo/internal/models"
)

type Service interface {
	GetAll() ([]models.ProductResponse, error)
}
