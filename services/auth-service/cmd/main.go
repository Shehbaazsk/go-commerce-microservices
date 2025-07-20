package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/shehbaazsl/go-commerce-microservices/auth-service/config"
	"github.com/shehbaazsl/go-commerce-microservices/auth-service/routes"
)

func main() {
	// Load .env file (relative to root or service folder)
	if err := godotenv.Load("./auth-service/.env"); err != nil {
		log.Fatalf("Failed to load env file: %v", err)
	}

	// Load configuration
	cfg := config.LoadConfig()

	// dbPool, err := db.ConnectDB(cfg.DBURL)
	// if err != nil {
	// 	log.Fatalf("Failed to connect to database: %v", err)
	// }

	// Set up routes
	r := gin.New()
	r.Use(gin.Recovery())

	routes.SetupRoutes(r)
	// Start the server
	if err := r.Run(":" + cfg.AppPort); err != nil {
		log.Fatal("Failed to start server", "error", err)
	}
}
