package config

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitDB() (*pgxpool.Pool, error) {
	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		return nil, fmt.Errorf("DATABASE_URL is not set")
	}

	// Create a context with timeout for DB connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Connect to the PostgreSQL DB using a connection pool
	dbpool, err := pgxpool.New(ctx, databaseUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to DB: %w", err)
	}

	// Ping to verify connection
	err = dbpool.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to ping DB: %w", err)
	}

	return dbpool, nil
}
