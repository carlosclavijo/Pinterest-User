package users

import (
	"context"
	"errors"
	"github.com/carlosclavijo/Pinterest-Services/internal/application/user/dto"
	"github.com/carlosclavijo/Pinterest-Services/internal/application/user/queries"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUserHandler_HandleGetListByLanguage(t *testing.T) {
	ctx := context.Background()

	mockRepository := new(MockRepository)
	mockFactory := new(MockFactory)

	handler := NewUserHandler(mockRepository, mockFactory)
	qry := queries.GetUsersByLanguageQuery{
		Language: "Spanish",
	}

	usersList := listUsers()

	mockRepository.On("GetListByLanguage", ctx, "Spanish").Return(usersList, nil)

	resp, err := handler.HandleGetListByLanguage(ctx, qry)

	require.NotNil(t, resp)
	require.IsType(t, []*dto.UserDTO{}, resp)
	require.NoError(t, err)

	mockRepository.AssertExpectations(t)
}

func TestUserHandler_HandleGetListByLanguage_Error(t *testing.T) {
	ctx := context.Background()

	mockRepository := new(MockRepository)
	mockFactory := new(MockFactory)

	handler := NewUserHandler(mockRepository, mockFactory)
	qry := queries.GetUsersByLanguageQuery{
		Language: "Spanish",
	}

	mockRepository.On("GetListByLanguage", ctx, qry.Language).Return(nil, errors.New("new error"))

	resp, err := handler.HandleGetListByLanguage(ctx, qry)

	require.Nil(t, resp)
	require.Error(t, err)
}
