package service

import (
	"errors"
	"test-be-dbo/internal/auth/repository"
	"test-be-dbo/internal/models"

	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepo repository.Repository
}

func NewUserService(userRepo repository.Repository) Service {
	return &userService{userRepo: userRepo}
}

func GenerateHashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (u *userService) RegisterUser(user models.RegisterRequest) (models.UserResponse, error) {

	isEmailExists, err := u.userRepo.EmailExists(user.Email)
	if err != nil {
		return models.UserResponse{}, err
	}

	if isEmailExists {
		return models.UserResponse{}, errors.New("user already registered")
	}

	hashedPassword, err := GenerateHashPassword(user.Password)
	if err != nil {
		return models.UserResponse{}, err
	}

	userData := models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: hashedPassword,
	}

	result, err := u.userRepo.RegisterUser(userData)
	if err != nil {
		return models.UserResponse{}, err
	}

	response := models.UserResponse{
		ID:        result.ID.String(),
		Name:      result.Name,
		Email:     result.Email,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}

	return response, nil
}

func verifyPassword(hashPasswordDb, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashPasswordDb), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

func (u *userService) LoginUser(request models.LoginRequest) (models.UserResponse, error) {

	result, err := u.userRepo.GetLoginUser(request)
	if err != nil {
		return models.UserResponse{}, err
	}

	if err := verifyPassword(result.Password, request.Password); err != nil {
		return models.UserResponse{}, errors.New("Password is incorrect.")
	}

	response := models.UserResponse{
		ID:        result.ID.String(),
		Name:      result.Name,
		Email:     result.Email,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}

	return response, nil
}
