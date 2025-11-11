package boards

import (
	"context"
	"github.com/google/uuid"
)

type BoardRepository interface {
	GetAll(ctx context.Context) ([]*Board, error)
	GetList(ctx context.Context) ([]*Board, error)
	GetListByUserId(ctx context.Context, id uuid.UUID) ([]*Board, error)
	GetListByName(ctx context.Context, name string) ([]*Board, error)
	GetById(ctx context.Context, id uuid.UUID) (*Board, error)

	ExistById(ctx context.Context, id uuid.UUID) (bool, error)

	Create(ctx context.Context, b *Board) (*Board, error)
	Update(ctx context.Context, b *Board) error
	Delete(ctx context.Context, b *Board) error
}
