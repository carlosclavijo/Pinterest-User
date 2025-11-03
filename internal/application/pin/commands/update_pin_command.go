package commands

import "github.com/google/uuid"

type UpdatePinCommand struct {
	Id          uuid.UUID `json:"id"`
	Title       *string   `json:"title,omitempty"`
	Description *string   `json:"description,omitempty"`
	Visibility  *bool     `json:"visibility,omitempty"`
	Tags        *[]string `json:"tags,omitempty"`
}
