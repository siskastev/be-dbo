package repository

import (
	"test-be-dbo/internal/models"

	"github.com/google/uuid"
)

type Repository interface {
	ProductIDExists(id string) (bool, error)
	GetProductByID(id string) (models.Product, error)
	CreateOrder(order models.Order) (models.Order, error)
	DeleteOrder(id uuid.UUID) error
	// GetOrderByID(id string) (models.Order, error)
	OrderIDExist(id uuid.UUID) (bool, error)
	UpdateOrder(order models.Order, id uuid.UUID) (models.Order, error)
	OrderHasPaid(id uuid.UUID) (bool, error)
}
