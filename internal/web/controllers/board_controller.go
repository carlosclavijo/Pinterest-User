package controllers

import (
	"database/sql"
	"encoding/json"
	"github.com/carlosclavijo/Pinterest-Services/internal/application/board/commands"
	"github.com/carlosclavijo/Pinterest-Services/internal/application/board/dto"
	command "github.com/carlosclavijo/Pinterest-Services/internal/application/board/handlers"
	boards "github.com/carlosclavijo/Pinterest-Services/internal/domain/board"
	"github.com/carlosclavijo/Pinterest-Services/internal/infrastructure/persistence/repositories"
	"github.com/carlosclavijo/Pinterest-Services/internal/web/helpers"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type BoardController struct {
	commandHandler command.BoardHandler
}

func NewBoardController(db *sql.DB) *BoardController {
	repository := repositories.NewBoardRepository(db)
	factory := boards.NewBoardFactory()
	commandHandler := command.NewBoardHandler(repository, factory)
	return &BoardController{
		commandHandler: *commandHandler,
	}
}

func (c *BoardController) CreateBoard(w http.ResponseWriter, r *http.Request) {
	var cmd commands.CreateBoardCommand
	if err := json.NewDecoder(r.Body).Decode(&cmd); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "INVALID_REQUEST_BODY",
				Message: ErrJSONFormat,
			},
		})
		return
	}

	board, err := c.commandHandler.HandleCreate(r.Context(), cmd)
	if err != nil {
		errStr := err.Error()
		helpers.WriteJSON(w, http.StatusInternalServerError, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "BOARD_CREATION_FAILED",
				Message: "Could not create board",
				Err:     &errStr,
			},
		})
		return
	}

	helpers.WriteJSON(w, http.StatusCreated, helpers.Response[*dto.BoardResponse]{
		Success: true,
		Data:    board,
	})
}

func (c *BoardController) RegisterRoutes(r chi.Router) {
	r.Post("/create", c.CreateBoard)
}
