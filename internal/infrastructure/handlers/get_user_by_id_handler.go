package handlers

import (
	"context"
	"github.com/carlosclavijo/Pinterest-User/internal/application/dto"
	"github.com/carlosclavijo/Pinterest-User/internal/application/mappers"
	"github.com/carlosclavijo/Pinterest-User/internal/application/queries"
)

func (handler *UserHandler) HandleGetById(context context.Context, query queries.GetUserById) (*dto.UserDTO, error) {
	usr, err := handler.repository.GetById(context, query.Id)
	if err != nil {
		return nil, err
	}

	userDto := mappers.MapToUserDTO(usr)

	return userDto, nil
}
