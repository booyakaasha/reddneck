package app

import (
	"context"
	"fmt"

	"github.com/booyakaasha/reddneck/internal/db"
)

type App struct {
	config appConfig

	db *db.DB
}

func New(_ context.Context, opts ...Option) (*App, error) {
	config := newAppConfig(opts...)

	app := &App{
		config: config,
	}

	var err error

	app.db, err = db.New(config.db)
	if err != nil {
		return nil, fmt.Errorf("db.New: %w", err)
	}

	return app, nil
}
