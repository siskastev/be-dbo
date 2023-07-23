package routes

import (
	"test-be-dbo/internal/config/database"
	"test-be-dbo/internal/products/handler"
	"test-be-dbo/internal/products/repository"
	"test-be-dbo/internal/products/service"

	"github.com/gin-gonic/gin"
)

type ProductRoute struct {
	handler *handler.HandlerProducts
}

func NewProductRoute(h *handler.HandlerProducts) *ProductRoute {
	return &ProductRoute{
		handler: h,
	}
}

func ProvideManageHandler() *handler.HandlerProducts {
	repo := repository.NewProductRepository(database.DB)
	service := service.NewProductService(repo)
	return handler.NewHandlerProduct(service)
}

func (a *ProductRoute) RegisterRoute(r *gin.RouterGroup) {
	r.GET("/products", a.handler.GetAll)
}
