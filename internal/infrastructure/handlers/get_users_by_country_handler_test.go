package handlers

import (
	"context"
	"errors"
	"github.com/carlosclavijo/Pinterest-User/internal/application/user/dto"
	"github.com/carlosclavijo/Pinterest-User/internal/application/user/queries"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUserHandler_HandleGetListByCountry(t *testing.T) {
	ctx := context.Background()

	mockRepository := new(MockRepository)
	mockFactory := new(MockFactory)

	handler := NewUserHandler(mockRepository, mockFactory)
	qry := queries.GetUsersByCountry{
		Country: "Bolivia",
	}

	usersList := listUsers()

	mockRepository.On("GetListByCountry", ctx, "Bolivia").Return(usersList, nil)

	resp, err := handler.HandleGetListByCountry(ctx, qry)

	require.NotNil(t, resp)
	require.IsType(t, []*dto.UserDTO{}, resp)
	require.NoError(t, err)

	mockRepository.AssertExpectations(t)
}

func TestUserHandler_HandleGetListByCountry_Error(t *testing.T) {
	ctx := context.Background()

	mockRepository := new(MockRepository)
	mockFactory := new(MockFactory)

	handler := NewUserHandler(mockRepository, mockFactory)
	qry := queries.GetUsersByCountry{
		Country: "Bolivia",
	}

	mockRepository.On("GetListByCountry", ctx, qry.Country).Return(nil, errors.New("new error"))

	resp, err := handler.HandleGetListByCountry(ctx, qry)

	require.Nil(t, resp)
	require.Error(t, err)
}
