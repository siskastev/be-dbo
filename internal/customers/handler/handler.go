package handler

import (
	"net/http"
	"test-be-dbo/internal/customers/service"
	"test-be-dbo/internal/helpers"
	"test-be-dbo/internal/models"

	"github.com/gin-gonic/gin"
)

type HandlerCustomer struct {
	customerService service.Service
}

func NewHandlerCustomer(customerService service.Service) *HandlerCustomer {
	return &HandlerCustomer{customerService: customerService}
}

func (h *HandlerCustomer) CreateCustomer(c *gin.Context) {
	var request models.CustomerRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		validationErrors := helpers.HandleValidationErrors(err)
		c.JSON(http.StatusBadRequest, gin.H{"errors": validationErrors})
		return
	}

	createdBy := c.Request.Context().Value("user").(*models.UserResponse)
	request.CreatedBy = createdBy.Email

	customerResponse, err := h.customerService.CreateCustomer(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": customerResponse})
}

func (h *HandlerCustomer) GetByIDCustomer(c *gin.Context) {

	customerID := c.Param("id")

	CustomerResponse, err := h.customerService.GetByIDCustomer(customerID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"errors": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": CustomerResponse})
}

func (h *HandlerCustomer) UpdateCustomer(c *gin.Context) {
	var request models.CustomerRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		validationErrors := helpers.HandleValidationErrors(err)
		c.JSON(http.StatusBadRequest, gin.H{"errors": validationErrors})
		return
	}

	customerID := c.Param("id")

	_, err := h.customerService.GetByIDCustomer(customerID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"errors": err.Error()})
		return
	}

	updatedBy := c.Request.Context().Value("user").(*models.UserResponse)
	request.UpdatedBy = updatedBy.Email

	CustomerResponse, err := h.customerService.UpdateCustomer(request, customerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": CustomerResponse})
}

func (h *HandlerCustomer) DeleteCustomer(c *gin.Context) {

	customerID := c.Param("id")

	_, err := h.customerService.GetByIDCustomer(customerID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"errors": err.Error()})
		return
	}

	if err := h.customerService.DeleteCustomer(customerID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success Delete ID" + customerID})
}

func (h *HandlerCustomer) GetAllCustomers(c *gin.Context) {

	paginationParams := helpers.GetPaginationParams(c)

	var filter models.FilterCustomers
	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": "Invalid query parameters"})
		return
	}

	customers, meta, err := h.customerService.GetAll(
		paginationParams,
		filter,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": customers,
		"meta": meta,
	})
}
