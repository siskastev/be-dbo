package routes

import (
	"test-be-dbo/internal/auth/handler"
	"test-be-dbo/internal/auth/repository"
	"test-be-dbo/internal/auth/service"
	"test-be-dbo/internal/config/database"

	"github.com/gin-gonic/gin"
)

type AuthRoute struct {
	handler *handler.HandlerAuth
}

func NewAuthRoute(h *handler.HandlerAuth) *AuthRoute {
	return &AuthRoute{
		handler: h,
	}
}

func ProvideManageHandler() *handler.HandlerAuth {
	repo := repository.NewUserRepository(database.DB)
	service := service.NewUserService(repo)
	return handler.NewHandlerUser(service)
}

func (a *AuthRoute) RegisterRoute(r *gin.RouterGroup) {
	r.POST("/register", a.handler.Register)
	r.POST("/login", a.handler.Login)
}
