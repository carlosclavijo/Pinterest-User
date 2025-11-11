package handlers

import (
	"context"
	"errors"
	"github.com/carlosclavijo/Pinterest-Services/internal/application/user/dto"
	"github.com/carlosclavijo/Pinterest-Services/internal/domain/abstractions"
	"github.com/carlosclavijo/Pinterest-Services/internal/domain/user"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestUserHandler_HandleDelete(t *testing.T) {
	ctx := context.Background()

	mockRepository := new(MockRepository)
	mockFactory := new(MockFactory)

	handler := NewUserHandler(mockRepository, mockFactory)
	id := uuid.New()

	phoneStr := "+591-70926048"
	username, email, password, gender, birth, country, language, phone := valueObjects("john", "john@doe.com", "5tr0nG!.", "Male", time.Now().AddDate(-20, 0, 0), "Bolivia", "Spanish", &phoneStr, t)

	usr := users.NewUser("john", "doe", username, email, password, gender, birth, country, language, phone)
	usr.AggregateRoot = abstractions.NewAggregateRoot(id)

	mockRepository.On("ExistsById", ctx, id).Return(true, nil)
	mockRepository.On("GetById", ctx, id).Return(usr, nil)
	mockRepository.On("Delete", ctx, usr).Return(usr, nil)

	resp, err := handler.HandleDelete(ctx, id)

	require.NotNil(t, resp)
	require.IsType(t, &dto.UserResponse{}, resp)
	require.NoError(t, err)

	assert.NotNil(t, resp.DeletedAt)

	mockRepository.AssertExpectations(t)
}

func TestUserHandler_HandleDelete_IdError(t *testing.T) {
	ctx := context.Background()

	mockRepository := new(MockRepository)
	mockFactory := new(MockFactory)

	handler := NewUserHandler(mockRepository, mockFactory)
	id := uuid.Nil

	resp, err := handler.HandleDelete(ctx, id)

	require.Nil(t, resp)
	require.ErrorIs(t, err, users.ErrIdNilUser)
}

func TestUserHandler_HandleDelete_ExistError(t *testing.T) {
	ctx := context.Background()

	mockFactory := new(MockFactory)
	mockRepository := new(MockRepository)

	handler := NewUserHandler(mockRepository, mockFactory)
	id := uuid.New()

	mockRepository.On("ExistsById", ctx, id).Return(false, errors.New("new error"))

	resp, err := handler.HandleDelete(ctx, id)

	require.Nil(t, resp)
	require.Error(t, err)

	mockRepository.AssertExpectations(t)
}

func TestUserHandler_HandleDelete_NotFoundError(t *testing.T) {
	ctx := context.Background()

	mockRepository := new(MockRepository)
	mockFactory := new(MockFactory)

	handler := NewUserHandler(mockRepository, mockFactory)
	id := uuid.New()

	mockRepository.On("ExistsById", ctx, id).Return(false, nil)

	resp, err := handler.HandleDelete(ctx, id)

	require.Nil(t, resp)
	require.ErrorIs(t, err, users.ErrNotFoundUser)

	mockRepository.AssertExpectations(t)
}

func TestUserHandler_HandleDelete_GetByIdError(t *testing.T) {
	ctx := context.Background()

	mockRepository := new(MockRepository)
	mockFactory := new(MockFactory)

	handler := NewUserHandler(mockRepository, mockFactory)
	id := uuid.New()

	phoneStr := "+591-70926048"
	username, email, password, gender, birth, country, language, phone := valueObjects("john", "john@doe.com", "5tr0nG!.", "Male", time.Now().AddDate(-20, 0, 0), "Bolivia", "Spanish", &phoneStr, t)

	usr := users.NewUser("john", "doe", username, email, password, gender, birth, country, language, phone)
	usr.AggregateRoot = abstractions.NewAggregateRoot(id)

	mockRepository.On("ExistsById", ctx, id).Return(true, nil)
	mockRepository.On("GetById", ctx, id).Return(nil, users.ErrNotFoundUser)

	resp, err := handler.HandleDelete(ctx, id)

	require.Nil(t, resp)
	require.ErrorIs(t, err, users.ErrNotFoundUser)

	mockRepository.AssertExpectations(t)
}

func TestUserHandler_HandleDelete_ChangeDeleteError(t *testing.T) {
	ctx := context.Background()

	mockRepository := new(MockRepository)
	mockFactory := new(MockFactory)

	handler := NewUserHandler(mockRepository, mockFactory)
	id := uuid.New()

	phoneStr := "+591-70926048"
	username, email, password, gender, birth, country, language, phone := valueObjects("john", "john@doe.com", "5tr0nG!.", "Male", time.Now().AddDate(-20, 0, 0), "Bolivia", "Spanish", &phoneStr, t)

	usr := users.NewUser("john", "doe", username, email, password, gender, birth, country, language, phone)
	usr.AggregateRoot = abstractions.NewAggregateRoot(id)

	mockRepository.On("ExistsById", ctx, id).Return(true, nil)
	mockRepository.On("GetById", ctx, id).Return(usr, nil)

	err := usr.Delete()

	resp, err := handler.HandleDelete(ctx, id)

	require.Nil(t, resp)
	require.ErrorIs(t, err, users.ErrAlreadyDeletedUser)

	mockRepository.AssertExpectations(t)
}

func TestUserHandler_HandleDelete_DeleteError(t *testing.T) {
	ctx := context.Background()

	mockRepository := new(MockRepository)
	mockFactory := new(MockFactory)

	handler := NewUserHandler(mockRepository, mockFactory)
	id := uuid.New()

	phoneStr := "+591-70926048"
	username, email, password, gender, birth, country, language, phone := valueObjects("john", "john@doe.com", "5tr0nG!.", "Male", time.Now().AddDate(-20, 0, 0), "Bolivia", "Spanish", &phoneStr, t)

	usr := users.NewUser("john", "doe", username, email, password, gender, birth, country, language, phone)
	usr.AggregateRoot = abstractions.NewAggregateRoot(id)

	mockRepository.On("ExistsById", ctx, id).Return(true, nil)
	mockRepository.On("GetById", ctx, id).Return(usr, nil)
	mockRepository.On("Delete", ctx, usr).Return(nil, errors.New("mock repository error"))

	resp, err := handler.HandleDelete(ctx, id)

	require.Nil(t, resp)
	require.Error(t, err)

	mockRepository.AssertExpectations(t)
}
