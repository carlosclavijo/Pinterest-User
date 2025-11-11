package web

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/carlosclavijo/Pinterest-Services/internal/infrastructure/services"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewRoutes(t *testing.T) {
	db, _, _ := sqlmock.New()
	defer db.Close()

	routes := NewRoutes(db, services.JWTService{})
	require.NotNil(t, routes)
	require.NotNil(t, routes.UserController)
}

func TestRoutes_Router(t *testing.T) {
	db, _, _ := sqlmock.New()
	defer db.Close()

	routes := NewRoutes(db, services.JWTService{})
	router := routes.Router()

	require.NotNil(t, router)

}
