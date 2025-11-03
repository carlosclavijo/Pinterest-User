package handlers

import (
	"context"
	"errors"
	"github.com/carlosclavijo/Pinterest-User/internal/application/user/dto"
	"github.com/carlosclavijo/Pinterest-User/internal/application/user/queries"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUserHandler_HandleGetAll(t *testing.T) {
	ctx := context.Background()

	mockRepository := new(MockRepository)
	mockFactory := new(MockFactory)

	handler := NewUserHandler(mockRepository, mockFactory)
	qry := queries.GetAllUsers{}

	usersList := listUsers()

	mockRepository.On("GetAll", ctx).Return(usersList, nil)

	resp, err := handler.HandleGetAll(ctx, qry)

	require.NotNil(t, resp)
	require.IsType(t, []*dto.UserDTO{}, resp)
	require.NoError(t, err)
	require.Len(t, resp, 10)

	mockRepository.AssertExpectations(t)
}

func TestUserHandler_HandleGetAll_Error(t *testing.T) {
	ctx := context.Background()

	mockRepository := new(MockRepository)
	mockFactory := new(MockFactory)

	handler := NewUserHandler(mockRepository, mockFactory)
	qry := queries.GetAllUsers{}

	mockRepository.On("GetAll", ctx).Return(nil, errors.New("new error"))

	resp, err := handler.HandleGetAll(ctx, qry)

	require.Nil(t, resp)
	require.Error(t, err)
}
