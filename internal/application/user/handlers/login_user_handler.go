package handlers

import (
	"context"
	"github.com/carlosclavijo/Pinterest-Services/internal/application/user/commands"
	"github.com/carlosclavijo/Pinterest-Services/internal/application/user/dto"
	"github.com/carlosclavijo/Pinterest-Services/internal/application/user/mappers"
	"github.com/carlosclavijo/Pinterest-Services/internal/domain/shared"
	"github.com/carlosclavijo/Pinterest-Services/internal/domain/user"
	"golang.org/x/crypto/bcrypt"
)

func (h *UserHandler) HandleLogin(ctx context.Context, cmd commands.LoginUserCommand) (*dto.UserResponse, error) {
	email, err := shared.NewEmail(cmd.Email)
	if err != nil {
		return nil, err
	}

	exists, err := h.repository.ExistsByEmail(ctx, cmd.Email)
	if err != nil {
		return nil, err
	} else if !exists {
		return nil, users.ErrNotFoundUser
	}

	usr, err := h.repository.GetByEmail(ctx, email.String())
	if err != nil {
		return nil, err
	}

	password, err := shared.NewPassword(cmd.Password)
	if err != nil {
		return nil, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(usr.Password().String()), []byte(password.String())); err != nil {
		return nil, users.ErrInvalidCredentialsUser
	}

	usr.ChangeLastLoginAt()
	err = h.repository.Update(ctx, usr)
	if err != nil {
		return nil, err
	}

	userDTO := mappers.MapToUserDTO(usr)
	userResponse := mappers.MapToUserResponse(userDTO, usr.LastLoginAt(), usr.CreatedAt(), usr.UpdatedAt(), usr.DeletedAt())

	return userResponse, nil
}
