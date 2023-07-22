package service

import (
	"test-be-dbo/internal/models"
)

type Service interface {
	RegisterUser(request models.RegisterRequest) (models.UserResponse, error)
	LoginUser(request models.LoginRequest) (models.UserResponse, error)
}
