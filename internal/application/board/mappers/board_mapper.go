package mappers

import (
	"github.com/carlosclavijo/Pinterest-User/internal/application/board/dto"
	"github.com/carlosclavijo/Pinterest-User/internal/domain/board"
	"time"
)

func MapToBoardDTO(board *boards.Board) *dto.BoardDTO {
	return &dto.BoardDTO{
		Id:          board.Id(),
		UserId:      board.UserId(),
		Name:        board.Name(),
		Description: board.Description(),
		Visibility:  board.Visibility(),
		PinCount:    board.PinCount(),
		Portrait:    board.Portrait(),
	}
}

func MapToBoardResponse(board *dto.BoardDTO, createdAt, updatedAt time.Time, deletedAt *time.Time) *dto.BoardResponse {
	return &dto.BoardResponse{
		BoardDTO:  board,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		DeletedAt: deletedAt,
	}
}
