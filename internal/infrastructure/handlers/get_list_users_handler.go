package handlers

import (
	"context"
	"github.com/carlosclavijo/Pinterest-User/internal/application/dto"
	"github.com/carlosclavijo/Pinterest-User/internal/application/mappers"
	"github.com/carlosclavijo/Pinterest-User/internal/application/queries"
)

func (handler *UserHandler) HandleGetList(context context.Context, query queries.GetListUsers) ([]*dto.UserDTO, error) {
	users, err := handler.repository.GetList(context)

	if err != nil {
		return nil, err
	}

	var usersDTO []*dto.UserDTO
	for _, usr := range users {
		userDTO := mappers.MapToUserDTO(usr)
		usersDTO = append(usersDTO, userDTO)
	}

	return usersDTO, nil
}
