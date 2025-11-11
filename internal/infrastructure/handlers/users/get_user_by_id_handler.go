package users

import (
	"context"
	"github.com/carlosclavijo/Pinterest-Services/internal/application/user/dto"
	"github.com/carlosclavijo/Pinterest-Services/internal/application/user/mappers"
	"github.com/carlosclavijo/Pinterest-Services/internal/application/user/queries"
)

func (h *UserHandler) HandleGetById(context context.Context, query queries.GetUserByIdQuery) (*dto.UserDTO, error) {
	usr, err := h.repository.GetById(context, query.Id)
	if err != nil {
		return nil, err
	}

	userDto := mappers.MapToUserDTO(usr)

	return userDto, nil
}
