package boards

import (
	"github.com/google/uuid"
)

type BoardFactory interface {
	Create(userId uuid.UUID, name string, description *string, visibility bool) (*Board, error)
}

type boardFactory struct{}

func (b boardFactory) Create(userId uuid.UUID, name string, description *string, visibility bool) (*Board, error) {
	if userId == uuid.Nil {
		return nil, ErrIdNilBoard
	}

	if name == "" {
		return nil, ErrEmptyNameBoard
	}

	if len(name) > 50 {
		return nil, ErrLongNameBoard
	}

	if description != nil {
		if len(*description) > 500 {
			return nil, ErrLongDescriptionBoard
		}
	}

	return NewBoard(userId, name, description, visibility), nil
}

func NewBoardFactory() BoardFactory {
	return &boardFactory{}
}
