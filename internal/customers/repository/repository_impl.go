package repository

import (
	"test-be-dbo/internal/helpers"
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

func (c *customerRepository) GetAll(paginationParams helpers.PaginationParams, filters models.FilterCustomers) ([]models.Customer, int64, error) {
	var customers []models.Customer

	query := c.db.Model(&models.Customer{})

	if filters.Name != "" {
		query = query.Where("name LIKE ?", "%"+filters.Name+"%")
	}
	if filters.ID != "" {
		query = query.Where("id = ?", filters.ID)
	}
	if filters.Gender != "" {
		query = query.Where("gender = ?", filters.Gender)
	}

	var totalRecords int64
	if err := query.Count(&totalRecords).Error; err != nil {
		return nil, 0, err
	}

	offset := paginationParams.GetOffset()
	limit := paginationParams.PageSize

	if err := query.Offset(offset).Limit(limit).Find(&customers).Error; err != nil {
		return nil, 0, err
	}

	return customers, totalRecords, nil
}
