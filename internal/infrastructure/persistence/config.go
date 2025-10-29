package persistence

import (
	"github.com/carlosclavijo/Pinterest-User/internal/infrastructure/extensions"
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

// LoadConfig obtiene las variables desde Vault
func LoadConfig(v *extensions.VaultClient) *Config {
	dbSecret, err := v.GetSecret("pinterest_user/env")
	if err != nil {
		log.Fatalf("error obteniendo DB secrets: %v", err)
	}

	// Si quisieras guardar también el JWT en Vault:
	// jwtSecret, _ := v.GetSecret("pinterest_user/jwt")

	return &Config{
		DBUser:     dbSecret["DB_USER"].(string),
		DBPassword: dbSecret["DB_PASSWORD"].(string),
		DBName:     dbSecret["DB_NAME"].(string),
		DBHost:     dbSecret["DB_HOST"].(string),
		DBPort:     dbSecret["DB_PORT"].(string),
		JWTSecret:  "my-fallback-secret", // O léelo también desde Vault
	}
}
