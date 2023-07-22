package routes

import (
	"test-be-dbo/internal/config/database"
	"test-be-dbo/internal/customers/handler"
	"test-be-dbo/internal/customers/repository"
	"test-be-dbo/internal/customers/service"

	"github.com/gin-gonic/gin"
)

type CustomerRoute struct {
	handler *handler.HandlerCustomer
}

func NewCustomerRoute(h *handler.HandlerCustomer) *CustomerRoute {
	return &CustomerRoute{
		handler: h,
	}
}

func ProvideManageHandler() *handler.HandlerCustomer {
	repo := repository.NewCustomerRepository(database.DB)
	service := service.NewCustomerService(repo)
	return handler.NewHandlerCustomer(service)
}

func (a *CustomerRoute) RegisterRoute(r *gin.RouterGroup) {
	r.POST("/customers", a.handler.CreateCustomer)
	r.PUT("/customers/:id", a.handler.UpdateCustomer)
	r.DELETE("/customers/:id", a.handler.DeleteCustomer)
	r.GET("/customers/:id", a.handler.GetByIDCustomer)
}
