package main

import (
	"krushi-api/internal/handler"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Gin engine
	r := gin.Default()

	// API v1 group
	v1 := r.Group("/api/v1")
	{
		v1.GET("/health", handler.HealthCheck)
		v1.GET("/ping", handler.Ping)
	}

	// Start server on port 8080
	log.Println("Starting server on :8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
