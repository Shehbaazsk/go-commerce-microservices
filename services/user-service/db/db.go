package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectDB(DBURL string) (*pgxpool.Pool, error) {
	// Connect with context
	ctx := context.Background()

	dbPool, err := pgxpool.New(context.Background(), DBURL)
	if err != nil {
		log.Fatal("Failed to connect to DB: ", err)
	}

	// Test connection
	if err := dbPool.Ping(ctx); err != nil {
		log.Fatalf("failed to ping database: %s", err)
	}

	log.Println("Connected to the database successfully")
	return dbPool, nil
}
