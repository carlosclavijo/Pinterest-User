package handlers

import (
	"context"
	"errors"
	users2 "github.com/carlosclavijo/Pinterest-User/internal/application/commands/users"
	"github.com/carlosclavijo/Pinterest-User/internal/application/user/dto"
	"github.com/carlosclavijo/Pinterest-User/internal/domain/shared"
	"github.com/carlosclavijo/Pinterest-User/internal/domain/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"testing"
	"time"
)

func TestUserHandler_HandleLogin(t *testing.T) {
	ctx := context.Background()

	mockRepository := new(MockRepository)
	mockFactory := new(MockFactory)

	handler := NewUserHandler(mockRepository, mockFactory)
	cmd := validLoginCommand()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(cmd.Password), bcrypt.DefaultCost)
	require.NoError(t, err)

	password, err := shared.NewHashedPassword(string(hashedPassword))
	require.NoError(t, err)

	username, email, _, gender, birth, country, language, phone := valueObjects("username", cmd.Email, string(hashedPassword), "Male", time.Now().AddDate(-20, -10, -5), "Bolivia", "Spanish", nil, t)

	usr := users.NewUser("john", "doe", username, email, password, gender, birth, country, language, phone)

	mockRepository.On("ExistsByEmail", ctx, cmd.Email).Return(true, nil)
	mockRepository.On("GetByEmail", ctx, cmd.Email).Return(usr, nil)
	mockRepository.On("Update", ctx, usr).Return(usr, nil)

	time.Sleep(5 * time.Millisecond)

	resp, err := handler.HandleLogin(ctx, cmd)

	require.NotNil(t, resp)
	require.IsType(t, &dto.UserResponse{}, resp)
	require.NoError(t, err)

	assert.True(t, resp.LastLoginAt.After(resp.CreatedAt))

	mockRepository.AssertExpectations(t)
}

func TestUserHandler_HandleLogin_EmailError(t *testing.T) {
	ctx := context.Background()

	mockRepository := new(MockRepository)
	mockFactory := new(MockFactory)

	handler := NewUserHandler(mockRepository, mockFactory)
	cmd := validLoginCommand()

	cmd.Email = "invalid"

	resp, err := handler.HandleLogin(ctx, cmd)

	require.Nil(t, resp)
	require.Error(t, err)

	mockRepository.AssertExpectations(t)
}

func TestUserHandler_HandleLogin_ExistsError(t *testing.T) {
	ctx := context.Background()

	mockRepository := new(MockRepository)
	mockFactory := new(MockFactory)

	handler := NewUserHandler(mockRepository, mockFactory)
	cmd := validLoginCommand()

	mockRepository.On("ExistsByEmail", ctx, cmd.Email).Return(false, errors.New("new error"))

	resp, err := handler.HandleLogin(ctx, cmd)

	require.Nil(t, resp)
	require.Error(t, err)

	mockRepository.AssertExpectations(t)
}

func TestUserHandler_HandleLogin_NotFoundError(t *testing.T) {
	ctx := context.Background()

	mockRepository := new(MockRepository)
	mockFactory := new(MockFactory)

	handler := NewUserHandler(mockRepository, mockFactory)
	cmd := validLoginCommand()

	mockRepository.On("ExistsByEmail", ctx, cmd.Email).Return(false, nil)

	resp, err := handler.HandleLogin(ctx, cmd)

	require.Nil(t, resp)
	require.ErrorIs(t, err, users.ErrNotFoundUser)

	mockRepository.AssertExpectations(t)
}

func TestUserHandler_HandleLogin_GetByEmailError(t *testing.T) {
	ctx := context.Background()

	mockRepository := new(MockRepository)
	mockFactory := new(MockFactory)

	handler := NewUserHandler(mockRepository, mockFactory)
	cmd := validLoginCommand()

	mockRepository.On("ExistsByEmail", ctx, cmd.Email).Return(true, nil)
	mockRepository.On("GetByEmail", ctx, cmd.Email).Return(nil, users.ErrNotFoundUser)

	resp, err := handler.HandleLogin(ctx, cmd)

	require.Nil(t, resp)
	require.ErrorIs(t, err, users.ErrNotFoundUser)

	mockRepository.AssertExpectations(t)
}

func TestUserHandler_HandleLogin_PasswordError(t *testing.T) {
	ctx := context.Background()

	mockRepository := new(MockRepository)
	mockFactory := new(MockFactory)

	handler := NewUserHandler(mockRepository, mockFactory)
	cmd := validLoginCommand()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(cmd.Password), bcrypt.DefaultCost)
	require.NoError(t, err)

	password, err := shared.NewHashedPassword(string(hashedPassword))
	require.NoError(t, err)

	username, email, _, gender, birth, country, language, phone := valueObjects("username", cmd.Email, string(hashedPassword), "Male", time.Now().AddDate(-20, -10, -5), "Bolivia", "Spanish", nil, t)
	usr := users.NewUser("john", "doe", username, email, password, gender, birth, country, language, phone)

	mockRepository.On("ExistsByEmail", ctx, cmd.Email).Return(true, nil)
	mockRepository.On("GetByEmail", ctx, cmd.Email).Return(usr, nil)

	cmd.Password = "softpassword"

	resp, err := handler.HandleLogin(ctx, cmd)

	require.Nil(t, resp)
	require.Error(t, err)

	mockRepository.AssertExpectations(t)
}

func TestUserHandler_HandleLogin_UpdateError(t *testing.T) {
	ctx := context.Background()

	mockRepository := new(MockRepository)
	mockFactory := new(MockFactory)

	handler := NewUserHandler(mockRepository, mockFactory)
	cmd := validLoginCommand()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(cmd.Password), bcrypt.DefaultCost)
	require.NoError(t, err)

	password, err := shared.NewHashedPassword(string(hashedPassword))
	require.NoError(t, err)

	username, email, _, gender, birth, country, language, phone := valueObjects("username", cmd.Email, string(hashedPassword), "Male", time.Now().AddDate(-20, -10, -5), "Bolivia", "Spanish", nil, t)
	usr := users.NewUser("john", "doe", username, email, password, gender, birth, country, language, phone)

	mockRepository.On("ExistsByEmail", ctx, cmd.Email).Return(true, nil)
	mockRepository.On("GetByEmail", ctx, cmd.Email).Return(usr, nil)
	mockRepository.On("Update", ctx, usr).Return(nil, errors.New("new error"))

	resp, err := handler.HandleLogin(ctx, cmd)

	require.Nil(t, resp)
	require.Error(t, err)

	mockRepository.AssertExpectations(t)
}

func validLoginCommand() users2.LoginUserCommand {
	email := "valid@email.com"
	password := "5tr0ngP4ssworD!"

	cmd := users2.LoginUserCommand{
		Email:    email,
		Password: password,
	}

	return cmd
}
