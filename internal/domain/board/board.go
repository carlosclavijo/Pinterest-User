package boards

import (
	"errors"
	"github.com/carlosclavijo/Pinterest-User/internal/domain/abstractions"
	"github.com/google/uuid"
	"time"
)

var (
	ErrNilUserIdBoard = errors.New("user id cannot be nil")
	ErrEmptyNameBoard       = errors.New("name can't be empty")
	ErrLongNameBoard        = errors.New("name can't be more than 50 characters long")
	ErrLongDescriptionBoard = errors.New("description can't be more than 50 characters long")
)

type Board struct {
	*abstractions.AggregateRoot
	userId      uuid.UUID
	name        string
	description *string
	visibility  bool
	pinCount    int
	portrait    *string
	createdAt   time.Time
	updatedAt   time.Time
	deletedAt   *time.Time
}

func NewBoard(userId uuid.UUID, name string, description *string, visibility bool) *Board {
	return &Board{
		AggregateRoot: abstractions.NewAggregateRoot(uuid.New()),
		userId:        userId,
		name:          name,
		description:   description,
		visibility:    visibility,
		pinCount:      0,
		createdAt:     time.Now(),
		updatedAt:     time.Now(),
		deletedAt:     nil,
	}
}

func (b *Board) Id() uuid.UUID {
	return b.AggregateRoot.Entity.Id
}

func (b *Board) UserId() uuid.UUID {
	return b.userId
}

func (b *Board) Name() string {
	return b.name
}

func (b *Board) Description() *string {
	return b.description
}

func (b *Board) Visibility() bool {
	return b.visibility
}

func (b *Board) PinCount() int {
	return b.pinCount
}

func (b *Board) Portrait() *string {
	return b.portrait
}

func (b *Board) CreatedAt() time.Time {
	return b.createdAt
}

func (b *Board) UpdatedAt() time.Time {
	return b.updatedAt
}

func (b *Board) DeletedAt() *time.Time {
	return b.deletedAt
}

func (b *Board) ChangeName(name string) error {
	if name == "" {
		return ErrEmptyNameBoard
	} else if len(name) > 50 {
		return ErrLongNameBoard
	}
	b.name = name
	return nil
}

func (b *Board) ChangeDescription(description *string) error {
	if description != nil {
		if len(*description) > 500 {
			return ErrLongDescriptionBoard
		}
	}
	b.description = description
	return nil
}

func (b *Board) ChangeVisibility(visibility bool) {
	b.visibility = visibility
}

func (b *Board) PlusPinCount() {
	b.pinCount++
}

func (b *Board) MinusPinCount() {
	b.pinCount--
}

func (b *Board) ChangePortrait(portrait *string) {
	b.portrait = portrait
}

func (b *Board) Update() {
	b.updatedAt = time.Now()
}

func (b *Board) Delete() {
	now := time.Now()
	b.deletedAt = &now
}

func (b *Board) Restore() {
	b.deletedAt = nil
}

func NewBoardFromDB(id, userId uuid.UUID, name string, description *string, visibility bool, pinCount int, portrait *string, createdAt, updatedAt time.Time, deletedAt *time.Time) *Board {
	return &Board{
		AggregateRoot: abstractions.NewAggregateRoot(id),
		userId:        userId,
		name:          name,
		description:   description,
		visibility:    visibility,
		pinCount:      pinCount,
		portrait:      portrait,
		createdAt:     createdAt,
		updatedAt:     updatedAt,
		deletedAt:     deletedAt,
	}
}
