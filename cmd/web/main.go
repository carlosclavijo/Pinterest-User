package main

import (
	"github.com/carlosclavijo/Pinterest-Services/internal/infrastructure"
	"github.com/carlosclavijo/Pinterest-Services/internal/infrastructure/services"
	"github.com/carlosclavijo/Pinterest-Services/internal/web"
	"go.uber.org/zap"
	"net/http"
	"time"
)

const (
	connection      = ":8080"
	vaultConnection = "http://127.0.0.1:8200"
	token           = "root"
	environment     = "development"
)

func main() {
	// Initialize global logger
	services.InitLogger(environment)
	
	log := services.Logger()
	defer log.Sync()

	// Vault client
	vaultClient := services.NewVaultClient(vaultConnection, token)

	// Load configuration (from Vault or env)
	cfg := infrastructure.LoadConfig(vaultClient)

	// Database
	db, err := cfg.DBConfig.NewPostgresDB()
	if err != nil {
		log.Fatal("error initializing DB", zap.Error(err))
	}

	// Redis + JWT + Routes
	rdb := services.NewRedisClient()
	blacklistRepo := services.NewTokenBlacklistRepository(rdb)
	jwtService := services.NewJWTService(cfg.JWTSecret, time.Hour*24)
	routes := web.NewRoutes(db, jwtService, blacklistRepo, &cfg.EmailService)

	// Start server
	log.Info("Server starting", zap.String("connection", connection), zap.String("environment", environment))

	if err := http.ListenAndServe(connection, routes.Router()); err != nil {
		log.Fatal("server error", zap.Error(err))
	}
}

//go test ./... -coverprofile=coverage.out
