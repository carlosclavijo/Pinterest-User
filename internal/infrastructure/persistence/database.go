package persistence

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type DBConfig struct {
	DBUser     string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     string
}

func (c *DBConfig) NewPostgresDB() (*sql.DB, error) {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", c.DBUser, c.DBPassword, c.DBName, c.DBHost, c.DBPort)

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
