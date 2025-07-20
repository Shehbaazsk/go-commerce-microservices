package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/shehbaazsk/go-commerce-microservices/user-service/config"
	"github.com/shehbaazsk/go-commerce-microservices/user-service/db"
	"github.com/shehbaazsk/go-commerce-microservices/user-service/routes"
)

func main() {
	// Load .env file (relative to root or service folder)
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load env file: %v", err)
	}

	// Load configuration
	cfg := config.LoadConfig()

	if cfg.Debug {
		gin.SetMode(gin.DebugMode)
		log.Println("Running in DEBUG mode")
	} else {
		gin.SetMode(gin.ReleaseMode)
		log.Println("Running in RELEASE mode")
	}

	dbPool, err := db.ConnectDB(cfg.DBURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Set up routes
	r := gin.New()
	r.Use(gin.Recovery())

	routes.SetupRoutes(r, dbPool)

	// Start API Gateway server
	addr := cfg.AppHost + ":" + cfg.AppPort
	log.Printf("API Gateway running at %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
