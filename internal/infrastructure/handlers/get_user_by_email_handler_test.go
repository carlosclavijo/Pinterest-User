package handlers

import (
	"context"
	"errors"
	"github.com/carlosclavijo/Pinterest-User/internal/application/dto"
	"github.com/carlosclavijo/Pinterest-User/internal/application/queries"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUserHandler_HandleGetByEmail(t *testing.T) {
	ctx := context.Background()

	mockRepository := new(MockRepository)
	mockFactory := new(MockFactory)

	handler := NewUserHandler(mockRepository, mockFactory)
	usr := listUsers()[0]

	qry := queries.GetUserByEmail{
		Email: "valid@email.com",
	}

	mockRepository.On("GetByEmail", ctx, qry.Email).Return(usr, nil)

	resp, err := handler.HandleGetByEmail(ctx, qry)

	require.NotNil(t, resp)
	require.IsType(t, &dto.UserDTO{}, resp)
	require.NoError(t, err)

	mockRepository.AssertExpectations(t)
}

func TestUserHandler_HandleGetByEmail_Error(t *testing.T) {
	ctx := context.Background()

	mockRepository := new(MockRepository)
	mockFactory := new(MockFactory)

	handler := NewUserHandler(mockRepository, mockFactory)

	qry := queries.GetUserByEmail{
		Email: "valid@email.com",
	}

	mockRepository.On("GetByEmail", ctx, qry.Email).Return(nil, errors.New("new error"))

	resp, err := handler.HandleGetByEmail(ctx, qry)

	require.Nil(t, resp)
	require.Error(t, err)
}
