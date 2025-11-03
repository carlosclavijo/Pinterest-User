package main

import (
	"github.com/carlosclavijo/Pinterest-User/internal/infrastructure/persistence"
	"github.com/carlosclavijo/Pinterest-User/internal/infrastructure/services"
	"github.com/carlosclavijo/Pinterest-User/internal/web"
	"log"
	"net/http"
	"time"
)

const (
	connection      = ":8080"
	vaultConnection = "http://127.0.0.1:8200"
)

func main() {
	vaultClient := services.NewVaultClient(vaultConnection, "root")

	cfg := persistence.LoadConfig(vaultClient)

	db, err := persistence.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("[main] error inicializando DB: %v", err)
	}

	jwtService := services.NewJWTService(cfg.JWTSecret, time.Hour*24)
	routes := web.NewRoutes(db, *jwtService)

	log.Println("[main] Servidor iniciado en :8080")
	if err := http.ListenAndServe(connection, routes.Router()); err != nil {
		log.Fatalf("[main] error servidor: %v", err)
	}
}

//go test ./... -coverprofile=coverage.out
