package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/shehbaazsk/go-commerce-micorservices/api-gateway/config"
	"github.com/shehbaazsk/go-commerce-micorservices/api-gateway/logger"
	"github.com/shehbaazsk/go-commerce-micorservices/api-gateway/middlewares"
	"github.com/shehbaazsk/go-commerce-micorservices/api-gateway/routes"
)

func main() {
	// Load .env file (relative to root or service folder)
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load env file: %v", err)
	}

	// Load configuration
	cfg := config.LoadConfig()

	// Set Gin mode based on DEBUG env
	if cfg.Debug {
		gin.SetMode(gin.DebugMode)
		log.Println("Running in DEBUG mode")
	} else {
		gin.SetMode(gin.ReleaseMode)
		log.Println("Running in RELEASE mode")
	}

	// Set up
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(logger.CustomLogger())
	r.Use(middlewares.CORSMiddleware(cfg.AllowedOrigins))

	// Set up routes
	routes.SetupRoutes(r, cfg)

	// Start API Gateway server
	addr := cfg.AppHost + ":" + cfg.AppPort
	log.Printf("API Gateway running at %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
