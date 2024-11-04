package utils

import (
	"context"
	"github.com/jackc/pgx/v5"
)

type PgPool interface {
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Close()
}
