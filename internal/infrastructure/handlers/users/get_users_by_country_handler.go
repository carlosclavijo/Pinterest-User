package users

import (
	"context"
	"github.com/carlosclavijo/Pinterest-Services/internal/application/user/dto"
	"github.com/carlosclavijo/Pinterest-Services/internal/application/user/mappers"
	"github.com/carlosclavijo/Pinterest-Services/internal/application/user/queries"
)

func (h *UserHandler) HandleGetListByCountry(context context.Context, query queries.GetUsersByCountryQuery) ([]*dto.UserDTO, error) {
	users, err := h.repository.GetListByCountry(context, query.Country)

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
