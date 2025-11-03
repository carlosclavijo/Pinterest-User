package persistence

import (
	"github.com/carlosclavijo/Pinterest-User/internal/infrastructure/services"
	"log"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     string
	JWTSecret  string
}

func LoadConfig(v *services.VaultClient) *Config {
	secret, err := v.GetSecret("pinterest/vault")
	if err != nil {
		log.Fatalf("error fetching secrets: %v", err)
	}

	return &Config{
		DBUser:     secret["DB_USER"].(string),
		DBPassword: secret["DB_PASSWORD"].(string),
		DBName:     secret["DB_NAME"].(string),
		DBHost:     secret["DB_HOST"].(string),
		DBPort:     secret["DB_PORT"].(string),
		JWTSecret:  secret["JWT_SECRET"].(string),
	}
}
