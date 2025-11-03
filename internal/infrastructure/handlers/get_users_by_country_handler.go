package handlers

import (
	"context"
	"github.com/carlosclavijo/Pinterest-User/internal/application/user/dto"
	"github.com/carlosclavijo/Pinterest-User/internal/application/user/mappers"
	"github.com/carlosclavijo/Pinterest-User/internal/application/user/queries"
)

func (handler *UserHandler) HandleGetListByCountry(context context.Context, query queries.GetUsersByCountry) ([]*dto.UserDTO, error) {
	users, err := handler.repository.GetListByCountry(context, query.Country)

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
