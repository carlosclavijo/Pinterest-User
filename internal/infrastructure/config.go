package infrastructure

import (
	"github.com/carlosclavijo/Pinterest-Services/internal/infrastructure/persistence"
	"github.com/carlosclavijo/Pinterest-Services/internal/infrastructure/services"
	"log"
)

type Config struct {
	DBConfig     persistence.DBConfig
	JWTSecret    string
	EmailService services.EmailService
}

func LoadConfig(v *services.VaultClient) *Config {
	secret, err := v.GetSecret("pinterest/vault")
	if err != nil {
		log.Fatalf("error fetching secrets: %v", err)
	}

	dbConfig := persistence.DBConfig{
		DBUser:     secret["DB_USER"].(string),
		DBPassword: secret["DB_PASSWORD"].(string),
		DBName:     secret["DB_NAME"].(string),
		DBHost:     secret["DB_HOST"].(string),
		DBPort:     secret["DB_PORT"].(string),
	}

	emailConfig := services.EmailService{
		Host:     secret["SMTP_HOST"].(string),
		Port:     secret["SMTP_PORT"].(string),
		Username: secret["SMTP_USER"].(string),
		Password: secret["SMTP_PASS"].(string),
		From:     secret["SMTP_FROM"].(string),
		AppUrl:   secret["APP_URL"].(string),
	}

	return &Config{
		DBConfig:     dbConfig,
		JWTSecret:    secret["JWT_SECRET"].(string),
		EmailService: emailConfig,
	}
}
