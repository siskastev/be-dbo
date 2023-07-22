package repository

import "test-be-dbo/internal/models"

type Repository interface {
	EmailExists(email string) (bool, error)
	RegisterUser(user models.User) (models.User, error)
	GetLoginUser(request models.LoginRequest) (models.User, error)
}
