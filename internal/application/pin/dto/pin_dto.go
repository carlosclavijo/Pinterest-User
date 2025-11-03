package dto

import "github.com/google/uuid"

type PinDTO struct {
	Id           uuid.UUID `json:"id"`
	UserId       uuid.UUID `json:"user_id"`
	BoardId      uuid.UUID `json:"board_id"`
	Title        string    `json:"title"`
	Description  *string   `json:"description,omitempty"`
	Image        string    `json:"image"`
	SaveCount    int       `json:"save_count"`
	LikeCount    int       `json:"like_count"`
	CommentCount int       `json:"comment_count"`
	Visibility   bool      `json:"visibility"`
	Tags         []*TagDTO `json:"tags"`
}
