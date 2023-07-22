package repository

import (
	"test-be-dbo/internal/helpers"
	"test-be-dbo/internal/models"
)

type Repository interface {
	CreateCustomer(customer models.Customer) (models.Customer, error)
	EmailExists(email string) (bool, error)
	UpdateCustomer(customer models.Customer, id string) (models.Customer, error)
	DeleteCustomer(id string) error
	GetByIDCustomer(id string) (models.Customer, error)
	GetAll(paginationParams helpers.PaginationParams, filters models.FilterCustomers) ([]models.Customer, int64, error)
}
