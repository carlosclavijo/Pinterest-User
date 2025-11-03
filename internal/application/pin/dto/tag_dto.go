package dto

import "github.com/google/uuid"

type TagDTO struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
