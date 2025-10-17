package dto

import "time"

type UserResponse struct {
	*UserDTO
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
