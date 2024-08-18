package main

import (
	"fmt"
	"os"

	"github.com/booyakaasha/reddneck/internal/db"
	"gopkg.in/yaml.v3"
)

type config struct {
	Postgres postgresConfig `yaml:"postgres"`
}

type postgresConfig struct {
	DSN string `yaml:"dsn"`
}

func mustGetConfig(filepath string) config {
	cfg := config{}

	f, err := os.OpenFile(filepath, os.O_RDONLY|os.O_SYNC, 0)
	if err != nil {
		exit(fmt.Errorf("os.OpenFile: %w", err))
	}

	defer func() {
		_ = f.Close()
	}()

	if err = yaml.NewDecoder(f).Decode(&cfg); err != nil {
		exit(fmt.Errorf("yaml.Decode: %w", err))
	}

	return cfg
}

func (c config) dbConfig() db.Config {
	return db.Config{
		DSN: c.Postgres.DSN,
	}
}
