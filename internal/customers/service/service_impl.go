package service

import (
	"errors"
	"fmt"
	"test-be-dbo/internal/customers/repository"
	"test-be-dbo/internal/helpers"
	"test-be-dbo/internal/models"
)

type customerService struct {
	customerRepo repository.Repository
}

func NewCustomerService(customerRepo repository.Repository) Service {
	return &customerService{customerRepo: customerRepo}
}

func (c *customerService) CreateCustomer(request models.CustomerRequest) (models.CustomerResponse, error) {
	isEmailExists, err := c.customerRepo.EmailExists(request.Email)
	if err != nil {
		return models.CustomerResponse{}, err
	}

	if isEmailExists {
		return models.CustomerResponse{}, errors.New("customer already registered")
	}

	customerData := models.Customer{
		Name:      request.Name,
		Address:   request.Address,
		Email:     request.Email,
		Phone:     request.Phone,
		Gender:    request.Gender,
		CreatedBy: request.CreatedBy,
	}

	result, err := c.customerRepo.CreateCustomer(customerData)
	if err != nil {
		return models.CustomerResponse{}, err
	}

	response := models.CustomerResponse{
		ID:        result.ID.String(),
		Name:      result.Name,
		Email:     result.Email,
		Address:   result.Address,
		Phone:     result.Phone,
		Gender:    result.Gender,
		CreatedBy: result.CreatedBy,
		CreatedAt: &result.CreatedAt,
		UpdatedAt: &result.UpdatedAt,
	}

	return response, nil
}

func (c *customerService) UpdateCustomer(request models.CustomerRequest, id string) (models.CustomerResponse, error) {

	customerData := models.Customer{
		Name:      request.Name,
		Address:   request.Address,
		Email:     request.Email,
		Phone:     request.Phone,
		Gender:    request.Gender,
		UpdatedBy: request.UpdatedBy,
	}

	result, err := c.customerRepo.UpdateCustomer(customerData, id)

	fmt.Println(&result)

	if err != nil {
		return models.CustomerResponse{}, err
	}

	response := models.CustomerResponse{
		ID:        id,
		Name:      result.Name,
		Email:     result.Email,
		Address:   result.Address,
		Phone:     result.Phone,
		Gender:    result.Gender,
		UpdatedAt: &result.UpdatedAt,
		UpdatedBy: result.UpdatedBy,
	}

	return response, nil
}

func (c *customerService) DeleteCustomer(id string) error {
	result, err := c.customerRepo.GetByIDCustomer(id)

	if err != nil {
		return err
	}

	if err := c.customerRepo.DeleteCustomer(result.ID.String()); err != nil {
		return err
	}
	return nil

}

func (c *customerService) GetByIDCustomer(id string) (models.CustomerResponse, error) {
	result, err := c.customerRepo.GetByIDCustomer(id)
	if err != nil {
		return models.CustomerResponse{}, err
	}

	response := models.CustomerResponse{
		ID:        result.ID.String(),
		Name:      result.Name,
		Email:     result.Email,
		Address:   result.Address,
		Phone:     result.Phone,
		Gender:    result.Gender,
		CreatedBy: result.CreatedBy,
		CreatedAt: &result.CreatedAt,
		UpdatedAt: &result.UpdatedAt,
	}

	return response, nil
}

func (c *customerService) GetAll(paginationParams helpers.PaginationParams, filters models.FilterCustomers) ([]models.CustomerResponse, helpers.MetaData, error) {

	result, totalRecords, err := c.customerRepo.GetAll(paginationParams, filters)
	if err != nil {
		return []models.CustomerResponse{}, helpers.MetaData{}, err
	}

	metaData := helpers.MetaData{
		TotalRecords: totalRecords,
		Page:         paginationParams.Page,
		Offset:       paginationParams.GetOffset(),
		Limit:        paginationParams.PageSize,
	}
	metaData.TotalPages = metaData.CalculateTotalPage()

	var response []models.CustomerResponse
	for _, customer := range result {
		customerResponse := models.CustomerResponse{
			ID:        customer.ID.String(),
			Name:      customer.Name,
			Email:     customer.Email,
			Address:   customer.Address,
			Phone:     customer.Phone,
			Gender:    customer.Gender,
			CreatedBy: customer.CreatedBy,
			CreatedAt: &customer.CreatedAt,
			UpdatedAt: &customer.UpdatedAt,
			UpdatedBy: customer.UpdatedBy,
		}
		response = append(response, customerResponse)
	}

	return response, metaData, nil
}
