package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func Start(router *gin.Engine) error {
	// Create a server with the router
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("APP_PORT")),
		Handler: router,
	}

	// Create a context with a cancel function
	ctx, cancel := context.WithCancel(context.Background())

	// Listen for OS signals
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	// Start the server in a separate goroutine
	go func() {
		fmt.Printf("Server running on port %s\n", os.Getenv("APP_PORT"))
		log.Fatal(server.ListenAndServe())
	}()

	// Wait for OS signal or context cancellation
	select {
	case signal := <-signalChan:
		fmt.Printf("Received OS signal: %v\n", signal)
	case <-ctx.Done():
		fmt.Println("Context canceled")
	}

	//shutdown the server with a timeout
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Printf("Server shutdown error: %v", err)
	}

	// Trigger context cancellation to signal the server to stop
	cancel()

	fmt.Println("Server shutdown complete")

	return nil
}
