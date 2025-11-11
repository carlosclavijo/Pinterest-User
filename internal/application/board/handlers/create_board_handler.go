package handlers

import (
	"context"
	"github.com/carlosclavijo/Pinterest-Services/internal/application/board/commands"
	"github.com/carlosclavijo/Pinterest-Services/internal/application/board/dto"
	"github.com/carlosclavijo/Pinterest-Services/internal/application/board/mappers"
)

func (h *BoardHandler) HandleCreate(ctx context.Context, cmd commands.CreateBoardCommand) (*dto.BoardResponse, error) {
	boardFactory, err := h.factory.Create(cmd.UserId, cmd.Name, cmd.Description, cmd.Visibility)
	if err != nil {
		return nil, err
	}

	board, err := h.repository.Create(ctx, boardFactory)
	if err != nil {
		return nil, err
	}

	boardDto := mappers.MapToBoardDTO(board)
	boardResponse := mappers.MapToBoardResponse(boardDto, board.CreatedAt(), board.UpdatedAt(), board.DeletedAt())
	return boardResponse, nil
}
