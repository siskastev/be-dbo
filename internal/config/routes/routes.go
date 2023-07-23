package routes

import (
	authRoutes "test-be-dbo/internal/auth/routes"
	"test-be-dbo/internal/config/middleware"
	customerRoutes "test-be-dbo/internal/customers/routes"
	orderRoutes "test-be-dbo/internal/orders/routes"
	productRoutes "test-be-dbo/internal/products/routes"

	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {

	routeGroup := router.Group("/api")

	authRoutes.NewAuthRoute(
		authRoutes.ProvideManageHandler(),
	).RegisterRoute(routeGroup)

	productRoutes.NewProductRoute(
		productRoutes.ProvideManageHandler(),
	).RegisterRoute(routeGroup)

	routeGroup.Use(middleware.JWTMiddleware())
	{
		customerRoutes.NewCustomerRoute(
			customerRoutes.ProvideManageHandler(),
		).RegisterRoute(routeGroup)

		orderRoutes.NewOrderRoute(
			orderRoutes.ProvideManageHandler(),
		).RegisterRoute(routeGroup)
	}
}
