package handlers

import (
	"context"
	"encoding/base64"
	"github.com/carlosclavijo/Pinterest-Services/internal/application/user/commands"
	"github.com/carlosclavijo/Pinterest-Services/internal/application/user/dto"
	"github.com/carlosclavijo/Pinterest-Services/internal/application/user/mappers"
	"github.com/carlosclavijo/Pinterest-Services/internal/domain/email"
	"github.com/carlosclavijo/Pinterest-Services/internal/domain/shared"
	"github.com/carlosclavijo/Pinterest-Services/internal/domain/user"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

func (h *UserHandler) HandleCreate(ctx context.Context, cmd commands.CreateUserCommand) (*dto.UserResponse, error) {
	username, err := shared.NewUsername(cmd.Username)
	if err != nil {
		h.logger.Error("Invalid username", zap.String("username", cmd.Username), zap.Error(err))
		return nil, err
	}

	em, err := shared.NewEmail(cmd.Email)
	if err != nil {
		h.logger.Error("Invalid email", zap.String("email", cmd.Email), zap.Error(err))
		return nil, err
	}

	password, err := hashPassword(cmd.Password)
	if err != nil {
		h.logger.Error("Invalid password", zap.String("password", cmd.Password), zap.Error(err))
		return nil, err
	}

	gender, err := shared.ParseGender(cmd.Gender)
	if err != nil {
		h.logger.Error("Invalid gender", zap.String("gender", cmd.Gender), zap.Error(err))
		return nil, err
	}

	birth, err := shared.NewBirthDate(cmd.Birth)
	if err != nil {
		h.logger.Error("Invalid birth", zap.Time("birth", cmd.Birth), zap.Error(err))
		return nil, err
	}

	country, err := shared.ParseCountry(cmd.Country)
	if err != nil {
		h.logger.Error("Invalid country", zap.String("country", cmd.Country), zap.Error(err))
		return nil, err
	}

	language, err := shared.ParseLanguage(cmd.Language)
	if err != nil {
		h.logger.Error("Invalid country", zap.String("country", cmd.Country), zap.Error(err))
		return nil, err
	}

	phone, err := shared.NewPhone(cmd.Phone)
	if err != nil {
		return nil, err
	}

	if err = h.ensureUserDoesNotExist(ctx, cmd.Username, cmd.Email); err != nil {
		return nil, err
	}

	userFactory, err := h.factory.Create(cmd.FirstName, cmd.LastName, username, em, password, gender, birth, country, language, phone)
	if err != nil {
		return nil, err
	}

	usr, err := h.repository.Create(ctx, userFactory)
	if err != nil {
		return nil, err
	}

	token := base64.URLEncoding.EncodeToString(make([]byte, 32))

	verification := email.EmailVerification{
		Id:        uuid.New(),
		UserId:    usr.Id(),
		Token:     token,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}

	if err = h.emailRepo.Save(ctx, &verification); err != nil {
		return nil, err
	}

	go func() {
		err = h.emailService.SendVerificationEmail(usr.Email().String(), verification.Token)
		if err != nil {
			log.Printf("failed to send verification email to %s: %v", usr.Email().String(), err)
		}
	}()

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
