package persistence

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func NewPostgresDB(cfg *Config) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBHost, cfg.DBPort,
	)

	log.Printf("[db] DSN: %s", dsn)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Printf("[db] error abriendo conexión: %v", err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		log.Printf("[db] ping fallido: %v", err)
		return nil, err
	}

	return db, nil
}
