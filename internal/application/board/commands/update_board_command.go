package commands

import "github.com/google/uuid"

type UpdateBoardCommand struct {
	Id          uuid.UUID `json:"id"`
	Name        *string   `json:"name"`
	Description *string   `json:"description"`
	Visibility  *bool     `json:"visibility"`
}
