package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/shehbaazsk/go-commerce-micorservices/api-gateway/config"
	"github.com/shehbaazsk/go-commerce-micorservices/api-gateway/logger"
	"github.com/shehbaazsk/go-commerce-micorservices/api-gateway/routes"
)

func main() {
	// Load .env file (relative to root or service folder)
	if err := godotenv.Load("./api-gateway/.env"); err != nil {
		log.Fatalf("Failed to load env file: %v", err)
	}

	// Load configuration
	cfg := config.LoadConfig()

	// Set up routes
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(logger.CustomLogger())

	routes.SetupRoutes(r)
	// Start the server
	if err := r.Run(cfg.AppHost + ":" + cfg.AppPort); err != nil {
		log.Fatal("Failed to start server", "error", err)
	}
}
