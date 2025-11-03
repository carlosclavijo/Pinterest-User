package commands

import "github.com/google/uuid"

type CreatePinCommand struct {
	UserId      uuid.UUID `json:"user_id"`
	BoardId     uuid.UUID `json:"board_id"`
	Title       string    `json:"title"`
	Description *string   `json:"description,omitempty"`
	Tags        []string  `json:"tags"`
}
