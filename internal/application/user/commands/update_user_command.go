package commands

import (
	"github.com/google/uuid"
	"time"
)

type UpdateUserCommand struct {
	Id          uuid.UUID  `json:"id"`
	FirstName   *string    `json:"first_name,omitempty"`
	LastName    *string    `json:"last_name,omitempty"`
	UserName    *string    `json:"user_name,omitempty"`
	Email       *string    `json:"email,omitempty"`
	Password    *string    `json:"password,omitempty"`
	Gender      *string    `json:"gender,omitempty"`
	Birth       *time.Time `json:"birth,omitempty"`
	Country     *string    `json:"country,omitempty"`
	Language    *string    `json:"language,omitempty"`
	Phone       *string    `json:"phone,omitempty"`
	Information *string    `json:"information,omitempty"`
	ProfilePic  *string    `json:"profile_pic,omitempty"`
	Website     *string    `json:"web_site,omitempty"`
	Visibility  *bool      `json:"visibility,omitempty"`
}
