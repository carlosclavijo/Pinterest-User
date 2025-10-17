package handlers

import (
	"context"
	"github.com/carlosclavijo/Pinterest-User/internal/applicationn/commands"
	"github.com/carlosclavijo/Pinterest-User/internal/applicationn/dto"
	"github.com/carlosclavijo/Pinterest-User/internal/applicationn/mappers"
	"github.com/carlosclavijo/Pinterest-User/internal/domain/shared"
	"github.com/carlosclavijo/Pinterest-User/internal/domain/user"
	"github.com/google/uuid"
)

func (h *UserHandler) HandleUpdate(ctx context.Context, cmd commands.UpdateUserCommand) (*dto.UserResponse, error) {
	var err error

	if cmd.Id == uuid.Nil {
		return nil, user.ErrIdNilUser
	}

	exist, err := h.repository.ExistById(ctx, cmd.Id)
	if err != nil {
		return nil, err
	} else if !exist {
		return nil, user.ErrNotFoundUser
	}

	username, err := shared.NewUsername(cmd.Email)
	if err != nil {
		return nil, err
	}

	email, err := shared.NewEmail(cmd.Email)
	if err != nil {
		return nil, err
	}

	password, err := shared.NewPassword(cmd.Password)
	if err != nil {
		return nil, err
	}

	gender, err := shared.ParseGender(cmd.Gender)
	if err != nil {
		return nil, err
	}

	birth, err := shared.NewBirthDate(cmd.Birth)
	if err != nil {
		return nil, err
	}

	country, err := shared.ParseCountry(cmd.Country)
	if err != nil {
		return nil, err
	}

	language, err := shared.ParseLanguage(cmd.Language)
	if err != nil {
		return nil, err
	}

	phone, err := shared.NewPhone(cmd.Phone)
	if err != nil {
		return nil, err
	}

	userFactory, err := h.factory.Create(cmd.FirstName, cmd.LastName, username, email, password, gender, birth, country, language, phone)
	if err != nil {
		return nil, err
	}

	usr, err := h.repository.Create(ctx, userFactory)
	if err != nil {
		return nil, err
	}

	userDto := mappers.MapToUserDTO(usr)
	userResponse := mappers.MapToUserResponse(userDto, usr.CreatedAt(), usr.UpdatedAt(), usr.DeletedAt())

	return userResponse, nil
}
