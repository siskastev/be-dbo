package repository

import "test-be-dbo/internal/models"

type Repository interface {
	EmailExists(email string) (bool, error)
	RegisterUser(user models.User) (models.User, error)
	GetLoginUser(user models.LoginRequest) (models.User, error)
}
