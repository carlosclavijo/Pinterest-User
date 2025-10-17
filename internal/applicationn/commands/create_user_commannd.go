package commands

import "time"

type CreateUserCommand struct {
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
