package handler

import (
	"net/http"
	custService "test-be-dbo/internal/customers/service"
	"test-be-dbo/internal/helpers"
	"test-be-dbo/internal/models"
	ordService "test-be-dbo/internal/orders/service"

	"github.com/gin-gonic/gin"
)

type HandlerOrder struct {
	orderService    ordService.Service
	customerService custService.Service
}

func NewHandlerUser(orderService ordService.Service, customerService custService.Service) *HandlerOrder {
	return &HandlerOrder{orderService: orderService, customerService: customerService}
}

func (h *HandlerOrder) CreateOrder(c *gin.Context) {
	var request models.OrderRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		validationErrors := helpers.HandleValidationErrors(err)
		c.JSON(http.StatusBadRequest, gin.H{"errors": validationErrors})
		return
	}

	customerIDExist, _ := h.customerService.CustomerIDExists(request.CustomerID)

	errors := make(map[string][]string)
	if !customerIDExist {
		errors["customer_id"] = append(errors["customer_id"], "Customer not found")
	}

	if len(request.Products) < 1 {
		errors["products"] = append(errors["products"], "Products Empty")
	}

	// Check if each product ID exists using the orderService.ProductIDExists method
	invalidProductIDs := h.orderService.ProductIDExists(request)
	for _, errMessage := range invalidProductIDs {
		errors["products"] = append(errors["products"], errMessage)
	}

	if len(errors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
		return
	}

	createdBy := c.Request.Context().Value("user").(*models.UserResponse)
	request.CreatedBy = createdBy.Email

	result, err := h.orderService.CreateOrder(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": result})
}

func (h *HandlerOrder) DeleteHandler(c *gin.Context) {

	orderID := c.Param("id")
	id := helpers.ParseUUID(orderID)

	_, err := h.orderService.OrderIDExist(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"errors": err.Error()})
		return
	}

	if err := h.orderService.DeleteOrder(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success Delete OrderID " + orderID})
}

func (h *HandlerOrder) GetOrderByID(c *gin.Context) {

	orderID := c.Param("id")
	id := helpers.ParseUUID(orderID)

	_, err := h.orderService.OrderIDExist(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"errors": err.Error()})
		return
	}

	result, err := h.orderService.GetOrderByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}

func (h *HandlerOrder) UpdateOrder(c *gin.Context) {
	var request models.OrderRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		validationErrors := helpers.HandleValidationErrors(err)
		c.JSON(http.StatusBadRequest, gin.H{"errors": validationErrors})
		return
	}

	orderID := c.Param("id")
	id := helpers.ParseUUID(orderID)

	_, err := h.orderService.OrderIDExist(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"errors": err.Error()})
		return
	}

	isPaidOrder, _ := h.orderService.OrderHasPaid(id)
	if isPaidOrder {
		c.JSON(http.StatusBadRequest, gin.H{"errors": "Order Already PAID"})
		return
	}

	customerIDExist, _ := h.customerService.CustomerIDExists(request.CustomerID)

	errors := make(map[string][]string)

	if !customerIDExist {
		errors["customer_id"] = append(errors["customer_id"], "Customer not found")
	}

	if len(request.Products) < 1 {
		errors["products"] = append(errors["products"], "Products Empty")
	}

	// Check if each product ID exists using the orderService.ProductIDExists method
	invalidProductIDs := h.orderService.ProductIDExists(request)
	for _, errMessage := range invalidProductIDs {
		errors["products"] = append(errors["products"], errMessage)
	}

	if len(errors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
		return
	}

	updatedBy := c.Request.Context().Value("user").(*models.UserResponse)
	request.UpdatedBy = updatedBy.Email

	result, err := h.orderService.UpdateOrder(request, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}

func (h *HandlerOrder) GetAll(c *gin.Context) {

	paginationParams := helpers.GetPaginationParams(c)

	var filter models.FilterOrders
	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": "Invalid query parameters"})
		return
	}

	customers, meta, err := h.orderService.GetAll(
		paginationParams,
		filter,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": customers,
		"meta": meta,
	})
}
