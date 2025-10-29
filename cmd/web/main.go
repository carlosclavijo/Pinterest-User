package main

import (
	services "github.com/carlosclavijo/Pinterest-User/internal/infrastructure/extensions"
	"github.com/carlosclavijo/Pinterest-User/internal/infrastructure/persistence"
	"github.com/carlosclavijo/Pinterest-User/internal/web"
	"log"
	"net/http"
	"os"
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

	// 4. Inicializar servicios
	jwtService := services.NewJWTService(os.Getenv("JWT_SECRET"), time.Hour*24)

	// 5. Iniciar servidor HTTP
	routes := web.NewRoutes(db, *jwtService)

	log.Println("[main] Servidor iniciado en :8080")
	if err := http.ListenAndServe(connection, routes.Router()); err != nil {
		log.Fatalf("[main] error servidor: %v", err)
	}
}

//go test ./... -coverprofile=coverage.out
