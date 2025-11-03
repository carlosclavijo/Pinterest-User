package pins

import (
	"errors"
	"github.com/carlosclavijo/Pinterest-User/internal/domain/abstractions"
	"github.com/carlosclavijo/Pinterest-User/internal/domain/tag"
	"github.com/google/uuid"
	"slices"
	"time"
)

var (
	ErrNilUserIdPin       = errors.New("user id cannot be nil")
	ErrNilBoardIdPin      = errors.New("board id cannot be nil")
	ErrEmptyTitlePin      = errors.New("title can't be empty")
	ErrLongTitlePin       = errors.New("title can't be longer than 100 characters")
	ErrLongDescriptionPin = errors.New("title can't be longer than 100 characters")
	ErrManyTagsPin        = errors.New("a pin cannot have more than 10 tags")
)

type Pin struct {
	*abstractions.AggregateRoot
	userId       uuid.UUID
	boardId      uuid.UUID
	title        string
	description  *string
	image        *string
	saveCount    int
	likeCount    int
	commentCount int
	visibility   bool
	tags         []tag.Tag
	createdAt    time.Time
	updatedAt    time.Time
	deletedAt    *time.Time
}

func NewPin(userId, boardId uuid.UUID, title string, description *string, tags []tag.Tag) *Pin {
	return &Pin{
		AggregateRoot: abstractions.NewAggregateRoot(uuid.New()),
		userId:        userId,
		boardId:       boardId,
		title:         title,
		description:   description,
		saveCount:     0,
		likeCount:     0,
		commentCount:  0,
		visibility:    true,
		tags:          tags,
		createdAt:     time.Now(),
		updatedAt:     time.Now(),
	}
}

func (p *Pin) Id() uuid.UUID {
	return p.AggregateRoot.Entity.Id
}

func (p *Pin) UserId() uuid.UUID {
	return p.userId
}

func (p *Pin) BoardId() uuid.UUID {
	return p.boardId
}

func (p *Pin) Title() string {
	return p.title
}

func (p *Pin) Description() *string {
	return p.description
}

func (p *Pin) Image() string {
	return *p.image
}

func (p *Pin) SaveCount() int {
	return p.saveCount
}

func (p *Pin) LikeCount() int {
	return p.likeCount
}

func (p *Pin) CommentCount() int {
	return p.commentCount
}

func (p *Pin) Visibility() bool {
	return p.visibility
}

func (p *Pin) Tags() []tag.Tag {
	return p.tags
}

func (p *Pin) CreatedAt() time.Time {
	return p.createdAt
}

func (p *Pin) UpdatedAt() time.Time {
	return p.updatedAt
}

func (p *Pin) DeletedAt() *time.Time {
	return p.deletedAt
}

func (p *Pin) ChangeTitle(title string) error {
	if title == "" {
		return ErrEmptyTitlePin
	} else if len(title) > 100 {
		return ErrLongTitlePin
	}
	p.title = title
	return nil
}

func (p *Pin) ChangeDescription(description *string) error {
	if description != nil {
		if len(*description) > 500 {
			return ErrLongDescriptionPin
		}
	}
	p.description = description
	return nil
}

func (p *Pin) ChangeImage(image *string) {
	p.image = image
}

func (p *Pin) PlusSaveCount() {
	p.saveCount++
}

func (p *Pin) LessSaveCount() {
	p.saveCount--
}

func (p *Pin) PlusLikeCount() {
	p.likeCount++
}

func (p *Pin) LessLkeCount() {
	p.likeCount--
}

func (p *Pin) PlusCommentCount() {
	p.commentCount++
}

func (p *Pin) LessCommentCount() {
	p.commentCount--
}

func (p *Pin) AddTag(tag tag.Tag) {
	p.tags = append(p.tags, tag)
}

func (p *Pin) SubTag(tag tag.Tag) {
	i := slices.Index(p.tags, tag)
	p.tags = append(p.tags[:i], p.tags[i+1:]...)
}

func (p *Pin) Update() {
	p.updatedAt = time.Now()
}

func (p *Pin) Delete() {
	now := time.Now()
	p.deletedAt = &now
}

func (p *Pin) Restore() {
	p.deletedAt = nil
}
