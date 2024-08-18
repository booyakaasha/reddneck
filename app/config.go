package app

import "github.com/booyakaasha/reddneck/internal/db"

type appConfig struct {
	db db.Config
}

type Option func(*appConfig)

func newAppConfig(opts ...Option) appConfig {
	config := appConfig{}

	for _, opt := range opts {
		opt(&config)
	}

	return config
}

func WithDBConfig(cfg db.Config) Option {
	return func(ac *appConfig) { ac.db = cfg }
}
