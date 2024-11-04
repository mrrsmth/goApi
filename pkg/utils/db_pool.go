//go:build !test
// +build !test

package utils

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

// NewPgPool creates a database pool.
func NewPgPool(logger *Logger, config *Config) PgPool {
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		config.DbUser,
		config.DbPassword,
		config.DbHost,
		config.DbPort,
		config.DbName,
	)

	pool, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		log.Fatal(err)
	}

	// Ping the database to verify the connection
	if err := pool.Ping(context.Background()); err != nil {
		log.Fatal(err)
	}

	// Set baseline connection pool settings
	pool.Config().MaxConns = 10

	db := PgPool(pool)
	return db
}
