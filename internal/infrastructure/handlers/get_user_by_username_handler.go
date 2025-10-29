package handlers

import (
	"context"
	"github.com/carlosclavijo/Pinterest-User/internal/application/dto"
	"github.com/carlosclavijo/Pinterest-User/internal/application/mappers"
	"github.com/carlosclavijo/Pinterest-User/internal/application/queries"
)

func (handler *UserHandler) HandleGetByUsername(context context.Context, query queries.GetUserByUsername) (*dto.UserDTO, error) {
	usr, err := handler.repository.GetByUsername(context, query.Username)
	if err != nil {
		return nil, err
	}

	userDTO := mappers.MapToUserDTO(usr)
	return userDTO, nil
}
