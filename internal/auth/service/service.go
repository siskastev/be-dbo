package service

import (
	"test-be-dbo/internal/models"
)

type Service interface {
	RegisterUser(user models.RegisterRequest) (models.UserResponse, error)
	LoginUser(user models.LoginRequest) (models.UserResponse, error)
}
