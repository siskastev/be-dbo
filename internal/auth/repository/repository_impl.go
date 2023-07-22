package repository

import (
	"test-be-dbo/internal/models"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) Repository {
	return &userRepository{db: db}
}

func (u *userRepository) EmailExists(email string) (bool, error) {
	var user models.User
	if err := u.db.Where(models.User{Email: email}).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return user.Email != "", nil
}

func (u *userRepository) RegisterUser(user models.User) (models.User, error) {
	if err := u.db.Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (u *userRepository) GetLoginUser(request models.LoginRequest) (models.User, error) {
	var user models.User
	if err := u.db.Where(models.User{Email: request.Email}).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return user, err
		}
	}
	return user, nil
}
