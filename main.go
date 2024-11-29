package main

import (
	"context"
	"verve/handler"
	"verve/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var (
	redisClient *redis.Client
	ctx         = context.Background()
)

func main() {
	// Initialize Redis client

	// Create a Gin router
	r := gin.Default()

	// Define the endpoint
	r.GET("/api/verve/accept", handler.HandleRequest)

	// Start the logging and resetting ticker
	go utils.StartTicker()

	// Run the server
	r.Run(":8080")
}
