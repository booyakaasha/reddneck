package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/booyakaasha/reddneck/app"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT)
	defer cancel()

	cfg := mustGetConfig("./config/local.yaml")

	_, err := app.New(
		ctx,
		app.WithDBConfig(cfg.dbConfig()),
	)
	if err != nil {
		exit(fmt.Errorf("app.New: %w", err))
	}
}

func exit(err error) {
	_, _ = fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
