package handlers

import (
	"context"
	"errors"
	"github.com/carlosclavijo/Pinterest-User/internal/application/user/dto"
	"github.com/carlosclavijo/Pinterest-User/internal/domain/abstractions"
	"github.com/carlosclavijo/Pinterest-User/internal/domain/user"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestUserHandler_HandleRestore(t *testing.T) {
	ctx := context.Background()

	mockRepository := new(MockRepository)
	mockFactory := new(MockFactory)

	handler := NewUserHandler(mockRepository, mockFactory)
	id := uuid.New()

	phoneStr := "+591-70926048"
	username, email, password, gender, birth, country, language, phone := valueObjects("john", "john@doe.com", "5tr0nG!.", "Male", time.Now().AddDate(-20, 0, 0), "Bolivia", "Spanish", &phoneStr, t)

	usr := users.NewUser("john", "doe", username, email, password, gender, birth, country, language, phone)
	usr.AggregateRoot = abstractions.NewAggregateRoot(id)

	_ = usr.Delete()
	require.NotNil(t, usr.DeletedAt())

	mockRepository.On("GetById", ctx, id).Return(usr, nil)
	mockRepository.On("Delete", ctx, usr).Return(usr, nil)

	resp, err := handler.HandleRestore(ctx, id)

	require.NotNil(t, resp)
	require.IsType(t, &dto.UserResponse{}, resp)
	require.NoError(t, err)
	assert.Nil(t, resp.DeletedAt)

	mockRepository.AssertExpectations(t)
}

func TestUserHandler_HandleRestore_IdError(t *testing.T) {
	ctx := context.Background()

	mockRepository := new(MockRepository)
	mockFactory := new(MockFactory)

	handler := NewUserHandler(mockRepository, mockFactory)
	id := uuid.Nil

	resp, err := handler.HandleRestore(ctx, id)

	require.Nil(t, resp)
	require.ErrorIs(t, err, users.ErrIdNilUser)
}

func TestUserHandler_HandleRestore_GetByIdError(t *testing.T) {
	ctx := context.Background()

	mockRepository := new(MockRepository)
	mockFactory := new(MockFactory)

	handler := NewUserHandler(mockRepository, mockFactory)
	id := uuid.New()

	mockRepository.On("GetById", ctx, id).Return(nil, users.ErrNotFoundUser)

	resp, err := handler.HandleRestore(ctx, id)

	require.Nil(t, resp)
	require.ErrorIs(t, err, users.ErrNotFoundUser)

	mockRepository.AssertExpectations(t)
}

func TestUserHandler_HandleRestore_ChangeDeleteError(t *testing.T) {
	ctx := context.Background()

	mockRepository := new(MockRepository)
	mockFactory := new(MockFactory)

	handler := NewUserHandler(mockRepository, mockFactory)
	id := uuid.New()

	phoneStr := "+591-70926048"
	username, email, password, gender, birth, country, language, phone := valueObjects("john", "john@doe.com", "5tr0nG!.", "Male", time.Now().AddDate(-20, 0, 0), "Bolivia", "Spanish", &phoneStr, t)

	usr := users.NewUser("john", "doe", username, email, password, gender, birth, country, language, phone)
	usr.AggregateRoot = abstractions.NewAggregateRoot(id)

	mockRepository.On("GetById", ctx, id).Return(usr, nil)

	resp, err := handler.HandleRestore(ctx, id)

	require.Nil(t, resp)
	require.ErrorIs(t, err, users.ErrAlreadyRestoredUser)

	mockRepository.AssertExpectations(t)
}

func TestUserHandler_HandleRestore_RestoreError(t *testing.T) {
	ctx := context.Background()

	mockRepository := new(MockRepository)
	mockFactory := new(MockFactory)

	handler := NewUserHandler(mockRepository, mockFactory)
	id := uuid.New()

	phoneStr := "+591-70926048"
	username, email, password, gender, birth, country, language, phone := valueObjects("john", "john@doe.com", "5tr0nG!.", "Male", time.Now().AddDate(-20, 0, 0), "Bolivia", "Spanish", &phoneStr, t)

	usr := users.NewUser("john", "doe", username, email, password, gender, birth, country, language, phone)
	usr.AggregateRoot = abstractions.NewAggregateRoot(id)

	_ = usr.Delete()
	require.NotNil(t, usr.DeletedAt())

	mockRepository.On("GetById", ctx, id).Return(usr, nil)
	mockRepository.On("Delete", ctx, usr).Return(nil, errors.New("mock repository error"))

	resp, err := handler.HandleRestore(ctx, id)

	require.Nil(t, resp)
	require.Error(t, err)

	mockRepository.AssertExpectations(t)
}
