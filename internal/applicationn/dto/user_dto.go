package dto

import (
	"github.com/google/uuid"
	"time"
)

type UserDTO struct {
	Id        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Gender    string    `json:"gender"`
	Birth     time.Time `json:"birth"`
	Country   string    `json:"country"`
	Language  string    `json:"language"`
	Phone     *string   `json:"phone,omitempty"`
}
