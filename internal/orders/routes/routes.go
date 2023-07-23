package routes

import (
	"test-be-dbo/internal/config/database"
	customerRepo "test-be-dbo/internal/customers/repository"
	customerService "test-be-dbo/internal/customers/service"
	"test-be-dbo/internal/orders/handler"
	orderRepo "test-be-dbo/internal/orders/repository"
	orderService "test-be-dbo/internal/orders/service"

	"github.com/gin-gonic/gin"
)

type OrderRoute struct {
	handler *handler.HandlerOrder
}

func NewOrderRoute(h *handler.HandlerOrder) *OrderRoute {
	return &OrderRoute{
		handler: h,
	}
}

func ProvideManageHandler() *handler.HandlerOrder {
	repoOrder := orderRepo.NewOrderRepository(database.DB)
	serviceOrder := orderService.NewOrderService(repoOrder)

	repoCustomer := customerRepo.NewCustomerRepository(database.DB)
	serviceCustomer := customerService.NewCustomerService(repoCustomer)
	return handler.NewHandlerUser(serviceOrder, serviceCustomer)
}

func (a *OrderRoute) RegisterRoute(r *gin.RouterGroup) {
	r.POST("/orders", a.handler.CreateOrder)
	r.PUT("/orders/:id", a.handler.UpdateOrder)
	r.DELETE("/orders/:id", a.handler.DeleteHandler)
	r.GET("/orders/:id", a.handler.GetOrderByID)
	r.GET("/orders", a.handler.GetAll)
}
