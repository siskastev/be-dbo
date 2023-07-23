package service

import (
	"test-be-dbo/internal/models"

	"github.com/google/uuid"
)

type Service interface {
	ProductIDExists(request models.OrderRequest) []string
	CreateOrder(request models.OrderRequest) (models.OrderResponse, error)
	DeleteOrder(id uuid.UUID) error
	OrderIDExist(id uuid.UUID) (bool, error)
	UpdateOrder(request models.OrderRequest, id uuid.UUID) (models.OrderResponse, error)
	OrderHasPaid(id uuid.UUID) (bool, error)
	GetOrderByID(id uuid.UUID) (models.ManageOrderResponse, error)
}
