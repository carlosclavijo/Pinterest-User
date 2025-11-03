package handlers

import (
	"context"
	"github.com/carlosclavijo/Pinterest-User/internal/application/user/dto"
	"github.com/carlosclavijo/Pinterest-User/internal/application/user/mappers"
	"github.com/carlosclavijo/Pinterest-User/internal/domain/shared"
	"github.com/carlosclavijo/Pinterest-User/internal/domain/user"
	"golang.org/x/crypto/bcrypt"
)

func (h *UserHandler) HandleCreate(ctx context.Context, cmd commands.CreateUserCommand) (*dto.UserResponse, error) {
	username, err := shared.NewUsername(cmd.Username)
	if err != nil {
		return nil, err
	}

	email, err := shared.NewEmail(cmd.Email)
	if err != nil {
		return nil, err
	}

	password, err := hashPassword(cmd.Password)
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

	if err = h.ensureUserDoesNotExist(ctx, cmd.Username, cmd.Email); err != nil {
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
	userResponse := mappers.MapToUserResponse(userDto, usr.LastLoginAt(), usr.CreatedAt(), usr.UpdatedAt(), usr.DeletedAt())
	return userResponse, nil
}

func (h *UserHandler) ensureUserDoesNotExist(ctx context.Context, username, email string) error {
	exist, err := h.repository.ExistsByUserName(ctx, username)
	if err != nil {
		return err
	}
	if exist {
		return users.ErrExistsUser
	}

	exist, err = h.repository.ExistsByEmail(ctx, email)
	if err != nil {
		return err
	}
	if exist {
		return users.ErrExistsUser
	}
	return nil
}

func hashPassword(raw string) (shared.Password, error) {
	if _, err := shared.NewPassword(raw); err != nil {
		return shared.Password{}, err
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(raw), bcrypt.DefaultCost)
	if err != nil {
		return shared.Password{}, err
	}

	return shared.NewHashedPassword(string(hashed))
}
