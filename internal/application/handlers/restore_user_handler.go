package handlers

import (
	"context"
	"github.com/carlosclavijo/Pinterest-User/internal/application/dto"
	"github.com/carlosclavijo/Pinterest-User/internal/application/mappers"
	"github.com/carlosclavijo/Pinterest-User/internal/domain/user"
	"github.com/google/uuid"
)

func (h *UserHandler) HandleRestore(ctx context.Context, id uuid.UUID) (*dto.UserResponse, error) {
	if id == uuid.Nil {
		return nil, users.ErrIdNilUser
	}

	usr, err := h.repository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	err = usr.Restore()
	if err != nil {
		return nil, err
	}

	if usr, err = h.repository.Delete(ctx, usr); err != nil {
		return nil, err
	}

	userDto := mappers.MapToUserDTO(usr)
	userResponse := mappers.MapToUserResponse(userDto, usr.LastLoginAt(), usr.CreatedAt(), usr.UpdatedAt(), usr.DeletedAt())

	return userResponse, nil
}
