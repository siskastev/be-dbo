package handler

import (
	"net/http"
	"test-be-dbo/internal/customers/service"
	"test-be-dbo/internal/helpers"
	"test-be-dbo/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type HandlerCustomer struct {
	customerService service.Service
}

func NewHandlerCustomer(customerService service.Service) *HandlerCustomer {
	return &HandlerCustomer{customerService: customerService}
}

func (h *HandlerCustomer) CreateCustomer(c *gin.Context) {
	logger := c.MustGet("logger").(*logrus.Logger)

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
		logger.WithFields(logrus.Fields{
			"method":  c.Request.Method,
			"route":   c.Request.URL.Path,
			"error":   err.Error(),
			"payload": request,
		}).Error("Internal Server Error")
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
	logger := c.MustGet("logger").(*logrus.Logger)

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
		logger.WithFields(logrus.Fields{
			"method":  c.Request.Method,
			"route":   c.Request.URL.Path,
			"error":   err.Error(),
			"payload": request,
		}).Error("Internal Server Error")
		c.JSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": CustomerResponse})
}

func (h *HandlerCustomer) DeleteCustomer(c *gin.Context) {
	logger := c.MustGet("logger").(*logrus.Logger)

	customerID := c.Param("id")

	_, err := h.customerService.GetByIDCustomer(customerID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"errors": err.Error()})
		return
	}

	if err := h.customerService.DeleteCustomer(customerID); err != nil {
		logger.WithFields(logrus.Fields{
			"method": c.Request.Method,
			"route":  c.Request.URL.Path,
			"error":  err.Error(),
		}).Error("Internal Server Error")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success Delete ID" + customerID})
}

func (h *HandlerCustomer) GetAllCustomers(c *gin.Context) {

	logger := c.MustGet("logger").(*logrus.Logger)

	paginationParams := helpers.GetPaginationParams(c)

	var filter models.FilterCustomers
	if err := c.ShouldBindQuery(&filter); err != nil {
		logger.WithFields(logrus.Fields{
			"method": c.Request.Method,
			"route":  c.Request.URL.Path,
			"error":  err.Error(),
		}).Error("Invalid query parameters")
		c.JSON(http.StatusBadRequest, gin.H{"errors": "Invalid query parameters"})
		return
	}

	customers, meta, err := h.customerService.GetAll(
		paginationParams,
		filter,
	)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"method": c.Request.Method,
			"route":  c.Request.URL.Path,
			"error":  err.Error(),
		}).Error("Internal server error")
		c.JSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": customers,
		"meta": meta,
	})
}
