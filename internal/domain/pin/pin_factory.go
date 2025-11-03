package pins

import (
	"github.com/carlosclavijo/Pinterest-User/internal/domain/tag"
	"github.com/google/uuid"
)

type PinFactory interface {
	Create(userId, boardId uuid.UUID, title string, description *string, tags []tag.Tag) (*Pin, error)
}

type pinFactory struct{}

func (p pinFactory) Create(userId, boardId uuid.UUID, title string, description *string, tags []tag.Tag) (*Pin, error) {
	if userId == uuid.Nil {
		return nil, ErrNilUserIdPin
	}

	if boardId == uuid.Nil {
		return nil, ErrNilBoardIdPin
	}

	if title == "" {
		return nil, ErrEmptyTitlePin
	}

	if len(title) > 100 {
		return nil, ErrLongTitlePin
	}

	if description != nil {
		if len(*description) > 500 {
			return nil, ErrLongDescriptionPin
		}
	}

	if len(tags) > 10 {
		return nil, ErrManyTagsPin
	}

	return NewPin(userId, boardId, title, description, tags), nil
}

func NewPinFactory() PinFactory {
	return &pinFactory{}
}
