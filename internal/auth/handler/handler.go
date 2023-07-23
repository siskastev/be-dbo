package handler

import (
	"net/http"
	"test-be-dbo/internal/auth/service"
	"test-be-dbo/internal/helpers"
	"test-be-dbo/internal/models"

	"github.com/gin-gonic/gin"
)

type HandlerAuth struct {
	userService service.Service
}

func NewHandlerUser(userService service.Service) *HandlerAuth {
	return &HandlerAuth{userService: userService}
}

func (h *HandlerAuth) Register(c *gin.Context) {

	var request models.RegisterRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		validationErrors := helpers.HandleValidationErrors(err)
		c.JSON(http.StatusBadRequest, gin.H{"errors": validationErrors})
		return
	}

	userResponse, err := h.userService.RegisterUser(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
		return
	}

	// Generate JWT token
	token, err := helpers.GenerateJWT(userResponse)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
		return
	}

	response := models.UserResponseWithToken{
		UserResponse: userResponse,
		Token: models.UserSession{
			JWTToken: token,
		},
	}

	c.JSON(http.StatusCreated, gin.H{"data": response})
}

func (h *HandlerAuth) Login(c *gin.Context) {

	var request models.LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		validationErrors := helpers.HandleValidationErrors(err)
		c.JSON(http.StatusBadRequest, gin.H{"errors": validationErrors})
		return
	}

	userResponse, err := h.userService.LoginUser(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
		return
	}

	// Generate JWT token
	token, err := helpers.GenerateJWT(userResponse)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
		return
	}

	response := models.UserResponseWithToken{
		UserResponse: userResponse,
		Token: models.UserSession{
			JWTToken: token,
		},
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}
