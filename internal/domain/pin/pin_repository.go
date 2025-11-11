package pins

import (
	"context"
	"github.com/google/uuid"
)

type PinRepository interface {
	GetAll(ctx context.Context) ([]*Pin, error)
	GetList(ctx context.Context) ([]*Pin, error)
	GetListByUserId(ctx context.Context, id uuid.UUID) ([]*Pin, error)
	GetListByName(ctx context.Context, id uuid.UUID) ([]*Pin, error)
	GetListByTag(ctx context.Context, tag string) ([]*Pin, error)
	GetById(ctx context.Context, id uuid.UUID) ([]*Pin, error)

	ExistById(ctx context.Context, id uuid.UUID) (bool, error)

	Create(ctx context.Context, pin *Pin) (*Pin, error)
	Update(ctx context.Context, pin *Pin) error
	Delete(ctx context.Context, pin *Pin) error
}
