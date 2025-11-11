package users

import (
	"context"
	"github.com/carlosclavijo/Pinterest-Services/internal/application/user/dto"
	"github.com/carlosclavijo/Pinterest-Services/internal/application/user/mappers"
	"github.com/carlosclavijo/Pinterest-Services/internal/application/user/queries"
)

func (h *UserHandler) HandleGetByEmail(context context.Context, query queries.GetUserByEmailQuery) (*dto.UserDTO, error) {
	usr, err := h.repository.GetByEmail(context, query.Email)
	if err != nil {
		return nil, err
	}

	userDTO := mappers.MapToUserDTO(usr)
	return userDTO, nil
}
