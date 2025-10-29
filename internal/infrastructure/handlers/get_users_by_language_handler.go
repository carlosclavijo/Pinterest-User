package handlers

import (
	"context"
	"github.com/carlosclavijo/Pinterest-User/internal/application/dto"
	"github.com/carlosclavijo/Pinterest-User/internal/application/mappers"
	"github.com/carlosclavijo/Pinterest-User/internal/application/queries"
)

func (handler *UserHandler) HandleGetListByLanguage(context context.Context, query queries.GetUsersByLanguage) ([]*dto.UserDTO, error) {
	users, err := handler.repository.GetListByLanguage(context, query.Language)

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
