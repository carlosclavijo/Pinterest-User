package persistence

import (
	"database/sql"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestNewPostgresDB_OpenError(t *testing.T) {
	openDB = func(driver, dsn string) (*sql.DB, error) {
		return nil, errors.New("mock open error")
	}
	defer func() { openDB = sql.Open }()

	_ = os.Setenv("DB_USER", "u")
	_ = os.Setenv("DB_PASSWORD", "p")
	_ = os.Setenv("DB_NAME", "n")
	_ = os.Setenv("DB_HOST", "h")
	_ = os.Setenv("DB_PORT", "5432")

	db, err := NewPostgresDB()

	require.Nil(t, db)
	require.Error(t, err)
	assert.EqualError(t, err, "mock open error")
}
