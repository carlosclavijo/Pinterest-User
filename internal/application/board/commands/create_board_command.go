package commands

import "github.com/google/uuid"

type CreateBoardCommand struct {
	UserId      uuid.UUID `json:"user_id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	Visibility  bool      `json:"visibility"`
}
