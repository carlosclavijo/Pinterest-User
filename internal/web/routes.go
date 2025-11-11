package web

import (
	"database/sql"
	"github.com/carlosclavijo/Pinterest-Services/internal/infrastructure/services"
	"github.com/carlosclavijo/Pinterest-Services/internal/web/controllers"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
	"path/filepath"
)

type Routes struct {
	UserController  *controllers.UserController
	BoardController *controllers.BoardController
}

func NewRoutes(db *sql.DB, jwt *services.JWTService, blr *services.TokenBlacklist, emService *services.EmailService) *Routes {
	return &Routes{
		UserController:  controllers.NewUserController(db, jwt, blr, emService),
		BoardController: controllers.NewBoardController(db),
	}
}

func (routes *Routes) Router() chi.Router {
	mux := chi.NewRouter()

	uploadsPath, err := filepath.Abs("../../uploads/profile_pics")
	if err != nil {
		panic(err)
	}

	mux.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir(uploadsPath))))
	mux.Get("/verify-email", routes.UserController.VerifyEmail)
	mux.Get("/swagger/*", httpSwagger.WrapHandler)

	mux.Route("/users", routes.UserController.RegisterRoutes)
	mux.Route("/boards", routes.BoardController.RegisterRoutes)

	return mux
}
