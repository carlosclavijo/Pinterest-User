package commands

import (
	"github.com/google/uuid"
	"time"
)

type UpdateUserCommand struct {
	Id        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Gender    string    `json:"gender"`
	Birth     time.Time `json:"birth"`
	Country   string    `json:"country"`
	Language  string    `json:"language"`
	Phone     *string   `json:"phone,omitempty"`
}
