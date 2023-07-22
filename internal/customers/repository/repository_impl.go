package repository

import (
	"test-be-dbo/internal/models"

	"gorm.io/gorm"
)

type customerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) Repository {
	return &customerRepository{db: db}
}

func (c *customerRepository) CreateCustomer(customer models.Customer) (models.Customer, error) {
	if err := c.db.Create(&customer).Error; err != nil {
		return customer, err
	}
	return customer, nil
}

func (c *customerRepository) EmailExists(email string) (bool, error) {
	var customer models.Customer
	if err := c.db.Where(models.User{Email: email}).First(&customer).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return customer.Email != "", nil
}

func (c *customerRepository) UpdateCustomer(customer models.Customer, id string) (models.Customer, error) {
	if err := c.db.Where("id = ?", id).Updates(&customer).Error; err != nil {
		return customer, err
	}
	return customer, nil
}

func (c *customerRepository) DeleteCustomer(id string) error {
	var customer models.Customer

	if err := c.db.Where("id = ?", id).Delete(&customer).Error; err != nil {
		return err
	}
	return nil
}

func (c *customerRepository) GetByIDCustomer(id string) (models.Customer, error) {
	var customer models.Customer

	if err := c.db.Where("id = ?", id).First(&customer).Error; err != nil {
		return customer, err
	}
	return customer, nil
}
