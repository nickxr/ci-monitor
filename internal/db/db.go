package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

// NewDB creates a new connection pool to PostgreSQL using the provided DSN.
func NewDB(dsn string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(context.Background()); err != nil {
		return nil, err
	}

	log.Println("Successfully connected to PostgreSQL")
	return pool, nil
}
