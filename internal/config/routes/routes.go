package routes

import (
	authRoutes "test-be-dbo/internal/auth/routes"
	"test-be-dbo/internal/config/middleware"

	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {

	routeGroup := router.Group("/api")

	authRoutes.NewAuthRoute(
		authRoutes.ProvideManageHandler(),
	).RegisterRoute(routeGroup)

	routeGroup.Use(middleware.JWTMiddleware())
	{
		//
	}
}
