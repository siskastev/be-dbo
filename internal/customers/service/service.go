package service

import (
	"test-be-dbo/internal/helpers"
	"test-be-dbo/internal/models"
)

type Service interface {
	CreateCustomer(request models.CustomerRequest) (models.CustomerResponse, error)
	UpdateCustomer(request models.CustomerRequest, id string) (models.CustomerResponse, error)
	DeleteCustomer(id string) error
	GetByIDCustomer(id string) (models.CustomerResponse, error)
	GetAll(paginationParams helpers.PaginationParams, filter models.FilterCustomers) ([]models.CustomerResponse, helpers.MetaData, error)
}
