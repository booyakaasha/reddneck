package db

import (
	"database/sql"
	"fmt"

	"github.com/lib/pq"
)

type DB struct {
	db *sql.DB
}

// New создает и возвращает подключение к базе данных PostgreSQL с использованием sqlx.
func New(cfg Config) (*DB, error) {
	connector, err := pq.NewConnector(cfg.DSN)
	if err != nil {
		return nil, fmt.Errorf("pq.NewConnector: %w", err)
	}

	return &DB{
		db: sql.OpenDB(connector),
	}, nil
}
