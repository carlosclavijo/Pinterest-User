package handlers

import (
	"context"
	"github.com/carlosclavijo/Pinterest-Services/internal/application/board/commands"
	"github.com/carlosclavijo/Pinterest-Services/internal/application/board/dto"
	"github.com/carlosclavijo/Pinterest-Services/internal/domain/board"
	"github.com/google/uuid"
)

func (h *BoardHandler) HandleUpdate(ctx context.Context, cmd commands.UpdateBoardCommand) (*dto.BoardResponse, error) {
	if cmd.Id == uuid.Nil {
		return nil, boards.ErrIdNilBoard
	}

	exist, err := h.repository.ExistById(ctx, cmd.Id)
	if err != nil {
		return nil, err
	} else if !exist {
		return nil, boards.ErrNotFoundBoard
	}

	return nil, nil
}
