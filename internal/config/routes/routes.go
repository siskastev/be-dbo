package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {
	routeGroup := router.Group("/api")

	routeGroup.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello World"})
	})
}
