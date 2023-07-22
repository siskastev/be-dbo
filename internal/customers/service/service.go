package service

import "test-be-dbo/internal/models"

type Service interface {
	CreateCustomer(request models.CustomerRequest) (models.CustomerResponse, error)
	UpdateCustomer(request models.CustomerRequest, id string) (models.CustomerResponse, error)
	DeleteCustomer(id string) error
	GetByIDCustomer(id string) (models.CustomerResponse, error)
}
