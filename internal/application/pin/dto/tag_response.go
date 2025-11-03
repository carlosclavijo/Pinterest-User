package dto

import "time"

type TagResponse struct {
	*TagDTO
	CreatedAt time.Time  `json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
