package dto

import "github.com/google/uuid"

type BoardDTO struct {
	Id          uuid.UUID `json:"id"`
	UserId      uuid.UUID `json:"user_id"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	Visibility  bool      `json:"visibility"`
	PinCount    int       `json:"pin_count"`
	Portrait    *string   `json:"portrait,omitempty"`
}
