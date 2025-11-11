package users

import (
	"context"
	"errors"
	"github.com/carlosclavijo/Pinterest-Services/internal/application/user/dto"
	"github.com/carlosclavijo/Pinterest-Services/internal/application/user/queries"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUserHandler_HandleGetById(t *testing.T) {
	ctx := context.Background()

	mockRepository := new(MockRepository)
	mockFactory := new(MockFactory)

	handler := NewUserHandler(mockRepository, mockFactory)
	usr := listUsers()[0]

	qry := queries.GetUserByIdQuery{
		Id: uuid.New(),
	}

	mockRepository.On("GetById", ctx, qry.Id).Return(usr, nil)

	resp, err := handler.HandleGetById(ctx, qry)

	require.NotNil(t, resp)
	require.IsType(t, &dto.UserDTO{}, resp)
	require.NoError(t, err)

	mockRepository.AssertExpectations(t)
}

func TestUserHandler_HandleGetById_Error(t *testing.T) {
	ctx := context.Background()

	mockRepository := new(MockRepository)
	mockFactory := new(MockFactory)

	handler := NewUserHandler(mockRepository, mockFactory)

	qry := queries.GetUserByIdQuery{
		Id: uuid.New(),
	}

	mockRepository.On("GetById", ctx, qry.Id).Return(nil, errors.New("new error"))

	resp, err := handler.HandleGetById(ctx, qry)

	require.Nil(t, resp)
	require.Error(t, err)
}
