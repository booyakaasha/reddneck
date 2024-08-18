package test

import (
	"os"
	"testing"

	"github.com/booyakaasha/reddneck/internal/db"
	"github.com/stretchr/testify/require"
)

func NewDB(t *testing.T) *db.DB {
	dsn := os.Getenv("TEST_PG_DSN")
	require.NotEmpty(t, dsn)

	db, err := db.New(db.Config{
		DSN: dsn,
	})
	require.NoError(t, err)

	return db
}
