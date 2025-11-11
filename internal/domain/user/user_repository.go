package users

import (
	"context"
	"github.com/google/uuid"
)

type UserRepository interface {
	GetAll(ctx context.Context) ([]*User, error)
	GetList(ctx context.Context) ([]*User, error)
	GetById(ctx context.Context, id uuid.UUID) (*User, error)
	GetByUsername(ctx context.Context, username string) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	GetListByCountry(ctx context.Context, country string) ([]*User, error)
	GetListByLanguage(ctx context.Context, language string) ([]*User, error)
	GetListLikeUsername(ctx context.Context, name string) ([]*User, error)

	ExistsById(ctx context.Context, id uuid.UUID) (bool, error)
	ExistsByUserName(ctx context.Context, username string) (bool, error)
	ExistsByEmail(ctx context.Context, email string) (bool, error)

	Create(ctx context.Context, u *User) (*User, error)
	Update(ctx context.Context, u *User) error
	Delete(ctx context.Context, u *User) error
}
