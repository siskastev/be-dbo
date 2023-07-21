package main

import (
	"log"
	"test-be-dbo/internal/config/database"
	"test-be-dbo/internal/config/routes"
	"test-be-dbo/internal/config/server"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Initialize database
	database.Init()

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// Initialize routes from the "routes" package
	routes.Setup(router)

	if err := server.Start(router); err != nil {
		log.Fatalf("Server error: %v", err)
	}

}
