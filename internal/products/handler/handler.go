package handler

import (
	"net/http"

	"test-be-dbo/internal/products/service"

	"github.com/gin-gonic/gin"
)

type HandlerProducts struct {
	productService service.Service
}

func NewHandlerProduct(productService service.Service) *HandlerProducts {
	return &HandlerProducts{productService: productService}
}

func (h *HandlerProducts) GetAll(c *gin.Context) {
	result, err := h.productService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}
