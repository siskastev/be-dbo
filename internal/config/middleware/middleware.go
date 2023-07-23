package middleware

import (
	"context"
	"net/http"
	"strings"
	"test-be-dbo/internal/helpers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Token format is usually "Bearer <token>"
		tokenString := (strings.Split(authHeader, " "))[1]

		user, err := helpers.VerifyAndExtractUserFromJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Attach the user to the request context so it can be accessed in the subsequent handler.
		ctx := context.WithValue(c.Request.Context(), "user", user)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func CORS() gin.HandlerFunc {
	// Enable CORS middleware
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Authorization", "Content-Type"}

	return cors.New(config)
}
