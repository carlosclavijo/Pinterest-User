package handlers

import (
	"context"
	"errors"
	"github.com/carlosclavijo/Pinterest-User/internal/application/user/dto"
	"github.com/carlosclavijo/Pinterest-User/internal/application/user/queries"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUserHandler_HandleGetByUsername(t *testing.T) {
	ctx := context.Background()

	mockRepository := new(MockRepository)
	mockFactory := new(MockFactory)

	handler := NewUserHandler(mockRepository, mockFactory)
	usr := listUsers()[0]

	qry := queries.GetUserByUsername{
		Username: "username",
	}

	mockRepository.On("GetByUsername", ctx, qry.Username).Return(usr, nil)

	resp, err := handler.HandleGetByUsername(ctx, qry)

	require.NotNil(t, resp)
	require.IsType(t, &dto.UserDTO{}, resp)
	require.NoError(t, err)

	mockRepository.AssertExpectations(t)
}

func TestUserHandler_HandleGetByUsername_Error(t *testing.T) {
	ctx := context.Background()

	mockRepository := new(MockRepository)
	mockFactory := new(MockFactory)

	handler := NewUserHandler(mockRepository, mockFactory)

	qry := queries.GetUserByUsername{
		Username: "username",
	}

	mockRepository.On("GetByUsername", ctx, qry.Username).Return(nil, errors.New("new error"))

	resp, err := handler.HandleGetByUsername(ctx, qry)

	require.Nil(t, resp)
	require.Error(t, err)
}
