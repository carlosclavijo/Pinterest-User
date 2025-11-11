package users

import (
	"context"
	"github.com/carlosclavijo/Pinterest-Services/internal/application/user/dto"
	"github.com/carlosclavijo/Pinterest-Services/internal/application/user/mappers"
	"github.com/carlosclavijo/Pinterest-Services/internal/application/user/queries"
)

func (h *UserHandler) HandleGetByUsername(context context.Context, query queries.GetUserByUsernameQuery) (*dto.UserDTO, error) {
	usr, err := h.repository.GetByUsername(context, query.Username)
	if err != nil {
		return nil, err
	}

	userDTO := mappers.MapToUserDTO(usr)
	return userDTO, nil
}
