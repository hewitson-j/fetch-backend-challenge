package main

import (
	"context"
	"fetch-backend-challenge/routes"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	// Setup router
	r := routes.SetupRouter()

	// Define the server
	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	// Channel to listen for OS interrupt signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	// Run server in a separate goroutine so it doesn’t block
	go func() {
		log.Println("Server is running on port 8080...")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Block until we receive a shutdown signal
	<-quit
	log.Println("Received shutdown signal, shutting down gracefully...")

	// Create a timeout context (5 seconds) to allow ongoing requests to finish
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt to gracefully shut down the server
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited cleanly.")
}
