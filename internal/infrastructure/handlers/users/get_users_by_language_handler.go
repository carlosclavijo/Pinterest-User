package users

import (
	"context"
	"github.com/carlosclavijo/Pinterest-Services/internal/application/user/dto"
	"github.com/carlosclavijo/Pinterest-Services/internal/application/user/mappers"
	"github.com/carlosclavijo/Pinterest-Services/internal/application/user/queries"
)

func (h *UserHandler) HandleGetListByLanguage(context context.Context, query queries.GetUsersByLanguageQuery) ([]*dto.UserDTO, error) {
	users, err := h.repository.GetListByLanguage(context, query.Language)

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
