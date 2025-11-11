package handlers

import (
	"context"
	"github.com/carlosclavijo/Pinterest-Services/internal/application/user/dto"
	"github.com/carlosclavijo/Pinterest-Services/internal/application/user/mappers"
	"github.com/carlosclavijo/Pinterest-Services/internal/domain/user"
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

	if err = h.repository.Delete(ctx, usr); err != nil {
		return nil, err
	}

	userDto := mappers.MapToUserDTO(usr)
	userResponse := mappers.MapToUserResponse(userDto, usr.LastLoginAt(), usr.CreatedAt(), usr.UpdatedAt(), usr.DeletedAt())

	return userResponse, nil
}
