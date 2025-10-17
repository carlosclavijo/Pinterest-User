package user

import (
	"context"
	"github.com/google/uuid"
)

type UserRepository interface {
	GetAll(ctx context.Context) ([]*User, error)
	GetList(ctx context.Context) ([]*User, error)
	GetById(ctx context.Context, id uuid.UUID) (*User, error)
	GetByCountry(ctx context.Context, country string) ([]*User, error)
	GetByLanguage(ctx context.Context, language string) ([]*User, error)
	GetByUsername(ctx context.Context, username string) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	ExistById(ctx context.Context, id uuid.UUID) (bool, error)

	Create(ctx context.Context, user *User) (*User, error)
	Update(ctx context.Context, user *User) (*User, error)
	Delete(ctx context.Context, id uuid.UUID) (*User, error)
	Restore(ctx context.Context, id uuid.UUID) (*User, error)

	GetCountries(ctx context.Context) ([]string, error)
	GetLanguages(ctx context.Context) ([]string, error)
	GetDialCodes(ctx context.Context) ([]string, error)
}
