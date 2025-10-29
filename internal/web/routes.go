package web

import (
	"database/sql"
	"github.com/carlosclavijo/Pinterest-User/internal/infrastructure/extensions"
	"github.com/carlosclavijo/Pinterest-User/internal/web/controllers"
	"github.com/go-chi/chi/v5"
)

type Routes struct {
	UserController *controllers.UserController
}

func NewRoutes(db *sql.DB, jwt extensions.JWTService) *Routes {
	return &Routes{
		UserController: controllers.NewUserController(db, jwt),
	}
}

func (routes *Routes) Router() chi.Router {
	mux := chi.NewRouter()

	mux.Route("/users", routes.UserController.RegisterRoutes)

	return mux
}
