package repository

import "test-be-dbo/internal/models"

type Repository interface {
	CreateCustomer(customer models.Customer) (models.Customer, error)
	EmailExists(email string) (bool, error)
	UpdateCustomer(customer models.Customer, id string) (models.Customer, error)
	DeleteCustomer(id string) error
	GetByIDCustomer(id string) (models.Customer, error)
}
