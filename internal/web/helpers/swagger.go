package helpers

import (
	"github.com/carlosclavijo/Pinterest-Services/internal/application/user/dto"
	"github.com/carlosclavijo/Pinterest-Services/internal/domain/shared"
)

type GetListUsersDTO struct {
	Success bool           `json:"success"`
	Length  *int           `json:"length,omitempty"`
	Data    []*dto.UserDTO `json:"data"`
	Error   *Error         `json:"error,omitempty"`
}

type GetUserDTO struct {
	Success bool           `json:"success"`
	Data    []*dto.UserDTO `json:"data"`
	Error   *Error         `json:"error,omitempty"`
}

type GetUserResponse struct {
	Success bool                `json:"success"`
	Data    []*dto.UserResponse `json:"data"`
	Error   *Error              `json:"error,omitempty"`
}

type LoginSuccessResponse struct {
	Success bool `json:"success"`
	Data    struct {
		User  *dto.UserDTO `json:"user"`
		Token string       `json:"token"`
	} `json:"data"`
	Error *Error `json:"error,omitempty"`
}

type LogoutSuccessResponse struct {
	Success bool   `json:"success"`
	Data    string `json:"data"`
}

type GetCountriesList struct {
	Success bool             `json:"success"`
	Data    []shared.Country `json:"data"`
}

type GetLanguagesList struct {
	Success bool              `json:"success"`
	Data    []shared.Language `json:"data"`
}
