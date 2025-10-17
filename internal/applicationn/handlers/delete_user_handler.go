package handlers

import (
	"context"
	"github.com/carlosclavijo/Pinterest-User/internal/applicationn/dto"
	"github.com/carlosclavijo/Pinterest-User/internal/applicationn/mappers"
	"github.com/carlosclavijo/Pinterest-User/internal/domain/user"
	"github.com/google/uuid"
)

func (h *UserHandler) HandleDelete(ctx context.Context, id uuid.UUID) (*dto.UserResponse, error) {
	if id == uuid.Nil {
		return nil, user.ErrIdNilUser
	}

	exist, err := h.repository.ExistById(ctx, id)
	if err != nil {
		return nil, err
	} else if !exist {
		return nil, user.ErrNotFoundUser
	}

	usr, err := h.repository.Delete(ctx, id)
	if err != nil {
		return nil, err
	}

	userDto := mappers.MapToUserDTO(usr)
	userResponse := mappers.MapToUserResponse(userDto, usr.CreatedAt(), usr.UpdatedAt(), usr.DeletedAt())

	return userResponse, nil
}
