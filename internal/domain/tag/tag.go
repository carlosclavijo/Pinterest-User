package tag

import (
	"errors"
	"github.com/carlosclavijo/Pinterest-User/internal/domain/abstractions"
	"github.com/google/uuid"
	"time"
)

var ErrAlreadyDeletedTag = errors.New("already deleted tag")

type Tag struct {
	*abstractions.Entity
	name      string
	createdAt time.Time
	deletedAt *time.Time
}

func NewTag(name string) *Tag {
	return &Tag{
		Entity:    abstractions.NewEntity(uuid.New()),
		name:      name,
		createdAt: time.Now(),
	}
}

func (t *Tag) Id() uuid.UUID {
	return t.Entity.Id
}

func (t *Tag) Name() string {
	return t.name
}

func (t *Tag) CreatedAt() time.Time {
	return t.createdAt
}

func (t *Tag) DeletedAt() *time.Time {
	return t.deletedAt
}

func (t *Tag) Delete() error {
	if t.deletedAt != nil {
		return ErrAlreadyDeletedTag
	}

	now := time.Now()
	t.deletedAt = &now

	return nil
}
