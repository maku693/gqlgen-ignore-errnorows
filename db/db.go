package db

import (
	"context"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type DB interface {
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
}

func InitializeDB() (*sql.DB, error) {
	return sql.Open("sqlite3", "file:data.sqlite3")
}
