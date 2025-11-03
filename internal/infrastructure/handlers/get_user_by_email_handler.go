package handlers

import (
	"context"
	"github.com/carlosclavijo/Pinterest-User/internal/application/user/dto"
	"github.com/carlosclavijo/Pinterest-User/internal/application/user/mappers"
	"github.com/carlosclavijo/Pinterest-User/internal/application/user/queries"
)

func (handler *UserHandler) HandleGetByEmail(context context.Context, query queries.GetUserByEmail) (*dto.UserDTO, error) {
	usr, err := handler.repository.GetByEmail(context, query.Email)
	if err != nil {
		return nil, err
	}

	userDTO := mappers.MapToUserDTO(usr)
	return userDTO, nil
}
