package handlers

import (
	"context"
	"errors"
	"github.com/carlosclavijo/Pinterest-User/internal/application/commands"
	"github.com/carlosclavijo/Pinterest-User/internal/application/dto"
	"github.com/carlosclavijo/Pinterest-User/internal/domain/abstractions"
	"github.com/carlosclavijo/Pinterest-User/internal/domain/shared"
	users "github.com/carlosclavijo/Pinterest-User/internal/domain/user"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
	"time"
)

func TestUserHandler_HandleUpdate(t *testing.T) {
	ctx := context.Background()

	mockRepository := new(MockRepository)
	mockFactory := new(MockFactory)

	handler := NewUserHandler(mockRepository, mockFactory)

	cmd := validUpdateUserCommand()
	username, email, password, gender, birth, country, language, phone := valueObjects(*cmd.UserName, *cmd.Email, *cmd.Password, *cmd.Gender, *cmd.Birth, *cmd.Country, *cmd.Language, cmd.Phone, t)

	usr := users.NewUser(*cmd.FirstName, *cmd.LastName, username, email, password, gender, birth, country, language, phone)
	usr.AggregateRoot = abstractions.NewAggregateRoot(cmd.Id)

	time.Sleep(10 * time.Millisecond)

	mockRepository.On("ExistsById", ctx, cmd.Id).Return(true, nil)
	mockRepository.On("GetById", ctx, cmd.Id).Return(usr, nil)
	mockRepository.On("Update", ctx, usr).Return(usr, nil)

	resp, err := handler.HandleUpdate(ctx, cmd)

	require.NotNil(t, resp)
	require.IsType(t, &dto.UserResponse{}, resp)
	require.NoError(t, err)

	assert.NotNil(t, cmd.Id, resp.Id)
	assert.Equal(t, *cmd.FirstName, resp.FirstName)
	assert.Equal(t, *cmd.LastName, resp.LastName)
	assert.Equal(t, *cmd.UserName, resp.Username)
	assert.Equal(t, *cmd.Email, resp.Email)
	assert.Equal(t, *cmd.Gender, resp.Gender)
	assert.Equal(t, *cmd.Birth, resp.Birth)
	assert.Equal(t, *cmd.Country, resp.Country)
	assert.Equal(t, *cmd.Language, resp.Language)
	assert.Equal(t, cmd.Phone, resp.Phone)
	assert.Equal(t, cmd.Information, resp.Information)
	assert.Equal(t, cmd.ProfilePic, resp.ProfilePic)
	assert.Equal(t, cmd.Website, resp.Website)
	assert.True(t, resp.Visibility)
	assert.NotNil(t, resp.LastLoginAt)
	assert.WithinDuration(t, time.Now(), resp.LastLoginAt, time.Second)
	assert.NotNil(t, resp.CreatedAt)
	assert.WithinDuration(t, time.Now(), resp.CreatedAt, time.Second)
	assert.NotNil(t, resp.UpdatedAt)
	assert.WithinDuration(t, time.Now(), resp.UpdatedAt, time.Second)
	assert.True(t, resp.UpdatedAt.After(resp.CreatedAt))
	assert.Nil(t, resp.DeletedAt)

	mockRepository.AssertExpectations(t)
}

func TestUserHandler_HandleUpdate_IdError(t *testing.T) {
	ctx := context.Background()

	mockFactory := new(MockFactory)
	mockRepository := new(MockRepository)

	handler := NewUserHandler(mockRepository, mockFactory)
	cmd := validUpdateUserCommand()

	cmd.Id = uuid.Nil
	resp, err := handler.HandleUpdate(ctx, cmd)

	require.Nil(t, resp)
	require.ErrorIs(t, err, users.ErrIdNilUser)
}

func TestUserHandler_HandleUpdate_ExistError(t *testing.T) {
	ctx := context.Background()

	mockFactory := new(MockFactory)
	mockRepository := new(MockRepository)

	handler := NewUserHandler(mockRepository, mockFactory)
	cmd := validUpdateUserCommand()

	mockRepository.On("ExistsById", ctx, cmd.Id).Return(false, errors.New("new error"))

	resp, err := handler.HandleUpdate(ctx, cmd)

	require.Nil(t, resp)
	require.Error(t, err)

	mockRepository.AssertExpectations(t)
}

func TestUserHandler_HandleUpdate_NotFoundError(t *testing.T) {
	ctx := context.Background()

	mockFactory := new(MockFactory)
	mockRepository := new(MockRepository)

	handler := NewUserHandler(mockRepository, mockFactory)
	cmd := validUpdateUserCommand()

	mockRepository.On("ExistsById", ctx, cmd.Id).Return(false, nil)

	resp, err := handler.HandleUpdate(ctx, cmd)

	require.Nil(t, resp)
	require.ErrorIs(t, err, users.ErrNotFoundUser)

	mockRepository.AssertExpectations(t)
}

func TestUserHandler_HandleUpdate_GetByIdError(t *testing.T) {
	ctx := context.Background()

	mockFactory := new(MockFactory)
	mockRepository := new(MockRepository)

	handler := NewUserHandler(mockRepository, mockFactory)
	cmd := validUpdateUserCommand()

	mockRepository.On("ExistsById", ctx, cmd.Id).Return(true, nil)
	mockRepository.On("GetById", ctx, cmd.Id).Return(nil, users.ErrNotFoundUser)

	resp, err := handler.HandleUpdate(ctx, cmd)

	require.Nil(t, resp)
	require.ErrorIs(t, err, users.ErrNotFoundUser)

	mockRepository.AssertExpectations(t)
}

func TestUserHandler_HandleUpdate_FirstNameError(t *testing.T) {
	ctx := context.Background()

	mockFactory := new(MockFactory)
	mockRepository := new(MockRepository)

	handler := NewUserHandler(mockRepository, mockFactory)

	cases := []struct {
		name        string
		firstName   string
		expectedErr error
	}{
		{"empty first name", "", users.ErrEmptyFirstNameUser},
		{"too long first name", strings.Repeat("a", 101), users.ErrLongFirstNameUser},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			cmd := validUpdateUserCommand()
			username, email, password, gender, birth, country, language, phone := valueObjects(*cmd.UserName, *cmd.Email, *cmd.Password, *cmd.Gender, *cmd.Birth, *cmd.Country, *cmd.Language, cmd.Phone, t)

			cmd.FirstName = &tc.firstName

			usr := users.NewUser("John", *cmd.LastName, username, email, password, gender, birth, country, language, phone)
			usr.AggregateRoot = abstractions.NewAggregateRoot(cmd.Id)

			mockRepository.On("ExistsById", ctx, cmd.Id).Return(true, nil)
			mockRepository.On("GetById", ctx, cmd.Id).Return(usr, nil)

			resp, err := handler.HandleUpdate(ctx, cmd)

			require.Nil(t, resp)
			require.ErrorIs(t, err, tc.expectedErr)

			mockRepository.AssertExpectations(t)
		})
	}
}

func TestUserHandler_HandleUpdate_LastNameError(t *testing.T) {
	ctx := context.Background()

	mockFactory := new(MockFactory)
	mockRepository := new(MockRepository)

	handler := NewUserHandler(mockRepository, mockFactory)

	cases := []struct {
		name        string
		lastName    string
		expectedErr error
	}{
		{"empty last name", "", users.ErrEmptyLastNameUser},
		{"too long last name", strings.Repeat("a", 101), users.ErrLongLastNameUser},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			cmd := validUpdateUserCommand()
			username, email, password, gender, birth, country, language, phone := valueObjects(*cmd.UserName, *cmd.Email, *cmd.Password, *cmd.Gender, *cmd.Birth, *cmd.Country, *cmd.Language, cmd.Phone, t)

			cmd.LastName = &tc.lastName

			usr := users.NewUser(*cmd.FirstName, "Doe", username, email, password, gender, birth, country, language, phone)
			usr.AggregateRoot = abstractions.NewAggregateRoot(cmd.Id)

			mockRepository.On("ExistsById", ctx, cmd.Id).Return(true, nil)
			mockRepository.On("GetById", ctx, cmd.Id).Return(usr, nil)

			resp, err := handler.HandleUpdate(ctx, cmd)

			require.Nil(t, resp)
			require.ErrorIs(t, err, tc.expectedErr)

			mockRepository.AssertExpectations(t)
		})
	}
}

func TestUserHandler_HandleUpdate_UserNameError(t *testing.T) {
	ctx := context.Background()

	mockFactory := new(MockFactory)
	mockRepository := new(MockRepository)

	handler := NewUserHandler(mockRepository, mockFactory)

	cases := []struct {
		name        string
		username    string
		expectedErr error
	}{
		{"empty username", "", shared.ErrEmptyUsername},
		{"too long username", strings.Repeat("a", 31), shared.ErrLongUsername},
		{"too short username", "a", shared.ErrShortUsername},
		{"invalid username", "!3%3", shared.ErrInvalidUsername},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			cmd := validUpdateUserCommand()
			username, email, password, gender, birth, country, language, phone := valueObjects(*cmd.UserName, *cmd.Email, *cmd.Password, *cmd.Gender, *cmd.Birth, *cmd.Country, *cmd.Language, cmd.Phone, t)

			cmd.UserName = &tc.username

			usr := users.NewUser(*cmd.FirstName, *cmd.LastName, username, email, password, gender, birth, country, language, phone)
			usr.AggregateRoot = abstractions.NewAggregateRoot(cmd.Id)

			mockRepository.On("ExistsById", ctx, cmd.Id).Return(true, nil)
			mockRepository.On("GetById", ctx, cmd.Id).Return(usr, nil)

			resp, err := handler.HandleUpdate(ctx, cmd)

			require.Nil(t, resp)
			require.ErrorIs(t, err, tc.expectedErr)

			mockRepository.AssertExpectations(t)
		})
	}
}

func TestUserHandler_HandleUpdate_EmailError(t *testing.T) {
	ctx := context.Background()

	mockFactory := new(MockFactory)
	mockRepository := new(MockRepository)

	handler := NewUserHandler(mockRepository, mockFactory)

	cases := []struct {
		name        string
		email       string
		expectedErr error
	}{
		{"empty email", "", shared.ErrEmptyEmail},
		{"invalid email", "invalid.email@.com", shared.ErrInvalidEmail},
		{"too long local part", strings.Repeat("a", 65) + "@mail.com", shared.ErrLongLocalEmail},
		{"too long domain part", "a@" + strings.Repeat("a", 255) + ".com", shared.ErrLongDomainEmail},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			cmd := validUpdateUserCommand()
			username, email, password, gender, birth, country, language, phone := valueObjects(*cmd.UserName, *cmd.Email, *cmd.Password, *cmd.Gender, *cmd.Birth, *cmd.Country, *cmd.Language, cmd.Phone, t)

			cmd.Email = &tc.email

			usr := users.NewUser(*cmd.FirstName, *cmd.LastName, username, email, password, gender, birth, country, language, phone)
			usr.AggregateRoot = abstractions.NewAggregateRoot(cmd.Id)

			mockRepository.On("ExistsById", ctx, cmd.Id).Return(true, nil)
			mockRepository.On("GetById", ctx, cmd.Id).Return(usr, nil)

			resp, err := handler.HandleUpdate(ctx, cmd)

			require.Nil(t, resp)
			require.ErrorIs(t, err, tc.expectedErr)

			mockRepository.AssertExpectations(t)
		})
	}
}

func TestUserHandler_HandleUpdate_PasswordError(t *testing.T) {
	ctx := context.Background()

	mockFactory := new(MockFactory)
	mockRepository := new(MockRepository)

	handler := NewUserHandler(mockRepository, mockFactory)

	cases := []struct {
		name        string
		password    string
		expectedErr error
	}{
		{"empty password", "", shared.ErrEmptyPassword},
		{"too long password", strings.Repeat("a", 65), shared.ErrLongPassword},
		{"too short password", strings.Repeat("a", 6), shared.ErrShortPassword},
		{"invalid email", "softpassword", shared.ErrSoftPassword},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			cmd := validUpdateUserCommand()
			username, email, password, gender, birth, country, language, phone := valueObjects(*cmd.UserName, *cmd.Email, *cmd.Password, *cmd.Gender, *cmd.Birth, *cmd.Country, *cmd.Language, cmd.Phone, t)

			cmd.Password = &tc.password

			usr := users.NewUser(*cmd.FirstName, *cmd.LastName, username, email, password, gender, birth, country, language, phone)
			usr.AggregateRoot = abstractions.NewAggregateRoot(cmd.Id)

			mockRepository.On("ExistsById", ctx, cmd.Id).Return(true, nil)
			mockRepository.On("GetById", ctx, cmd.Id).Return(usr, nil)

			resp, err := handler.HandleUpdate(ctx, cmd)

			require.Nil(t, resp)
			require.ErrorIs(t, err, tc.expectedErr)

			mockRepository.AssertExpectations(t)
		})
	}
}

func TestUserHandler_HandleUpdate_GenderError(t *testing.T) {
	ctx := context.Background()

	mockFactory := new(MockFactory)
	mockRepository := new(MockRepository)

	handler := NewUserHandler(mockRepository, mockFactory)

	cmd := validUpdateUserCommand()
	username, email, password, gender, birth, country, language, phone := valueObjects(*cmd.UserName, *cmd.Email, *cmd.Password, *cmd.Gender, *cmd.Birth, *cmd.Country, *cmd.Language, cmd.Phone, t)

	genderStr := "X"
	cmd.Gender = &genderStr

	usr := users.NewUser(*cmd.FirstName, *cmd.LastName, username, email, password, gender, birth, country, language, phone)
	usr.AggregateRoot = abstractions.NewAggregateRoot(cmd.Id)

	mockRepository.On("ExistsById", ctx, cmd.Id).Return(true, nil)
	mockRepository.On("GetById", ctx, cmd.Id).Return(usr, nil)

	resp, err := handler.HandleUpdate(ctx, cmd)

	require.Nil(t, resp)
	require.ErrorIs(t, err, shared.ErrNotAGender)

	mockRepository.AssertExpectations(t)
}

func TestUserHandler_HandleUpdate_BirthError(t *testing.T) {
	ctx := context.Background()

	mockFactory := new(MockFactory)
	mockRepository := new(MockRepository)

	handler := NewUserHandler(mockRepository, mockFactory)

	cases := []struct {
		name        string
		birth       time.Time
		expectedErr error
	}{
		{"future date", time.Now().AddDate(10, 0, 0), shared.ErrFutureDate},
		{"less thatn 12 years", time.Now().AddDate(-10, 0, 0), shared.ErrUnderTwelve},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			cmd := validUpdateUserCommand()
			username, email, password, gender, birth, country, language, phone := valueObjects(*cmd.UserName, *cmd.Email, *cmd.Password, *cmd.Gender, *cmd.Birth, *cmd.Country, *cmd.Language, cmd.Phone, t)

			cmd.Birth = &tc.birth

			usr := users.NewUser(*cmd.FirstName, *cmd.LastName, username, email, password, gender, birth, country, language, phone)
			usr.AggregateRoot = abstractions.NewAggregateRoot(cmd.Id)

			mockRepository.On("ExistsById", ctx, cmd.Id).Return(true, nil)
			mockRepository.On("GetById", ctx, cmd.Id).Return(usr, nil)

			resp, err := handler.HandleUpdate(ctx, cmd)

			require.Nil(t, resp)
			require.ErrorIs(t, err, tc.expectedErr)

			mockRepository.AssertExpectations(t)
		})
	}
}

func TestUserHandler_HandleUpdate_CountryError(t *testing.T) {
	ctx := context.Background()

	mockFactory := new(MockFactory)
	mockRepository := new(MockRepository)

	handler := NewUserHandler(mockRepository, mockFactory)

	cmd := validUpdateUserCommand()
	username, email, password, gender, birth, country, language, phone := valueObjects(*cmd.UserName, *cmd.Email, *cmd.Password, *cmd.Gender, *cmd.Birth, *cmd.Country, *cmd.Language, cmd.Phone, t)

	countryStr := "X"
	cmd.Country = &countryStr

	usr := users.NewUser(*cmd.FirstName, *cmd.LastName, username, email, password, gender, birth, country, language, phone)
	usr.AggregateRoot = abstractions.NewAggregateRoot(cmd.Id)

	mockRepository.On("ExistsById", ctx, cmd.Id).Return(true, nil)
	mockRepository.On("GetById", ctx, cmd.Id).Return(usr, nil)

	resp, err := handler.HandleUpdate(ctx, cmd)

	require.Nil(t, resp)
	require.ErrorIs(t, err, shared.ErrNotACountry)

	mockRepository.AssertExpectations(t)
}

func TestUserHandler_HandleUpdate_LanguageError(t *testing.T) {
	ctx := context.Background()

	mockFactory := new(MockFactory)
	mockRepository := new(MockRepository)

	handler := NewUserHandler(mockRepository, mockFactory)

	cmd := validUpdateUserCommand()
	username, email, password, gender, birth, country, language, phone := valueObjects(*cmd.UserName, *cmd.Email, *cmd.Password, *cmd.Gender, *cmd.Birth, *cmd.Country, *cmd.Language, cmd.Phone, t)

	languageStr := "X"
	cmd.Language = &languageStr

	usr := users.NewUser(*cmd.FirstName, *cmd.LastName, username, email, password, gender, birth, country, language, phone)
	usr.AggregateRoot = abstractions.NewAggregateRoot(cmd.Id)

	mockRepository.On("ExistsById", ctx, cmd.Id).Return(true, nil)
	mockRepository.On("GetById", ctx, cmd.Id).Return(usr, nil)

	resp, err := handler.HandleUpdate(ctx, cmd)

	require.Nil(t, resp)
	require.ErrorIs(t, err, shared.ErrNotALanguage)

	mockRepository.AssertExpectations(t)
}

func TestUserHandler_HandleUpdate_PhoneError(t *testing.T) {
	ctx := context.Background()

	mockFactory := new(MockFactory)
	mockRepository := new(MockRepository)

	handler := NewUserHandler(mockRepository, mockFactory)

	cases := []struct {
		name        string
		phone       string
		expectedErr error
	}{
		{"not numeric phone", "+591-AA89879", shared.ErrNotNumericPhoneNumber},
		{"too short phone", "+591-771415", shared.ErrShortPhoneNumber},
		{"too long phone", "+591-7712131415161718", shared.ErrLongPhoneNumber},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			cmd := validUpdateUserCommand()
			username, email, password, gender, birth, country, language, phone := valueObjects(*cmd.UserName, *cmd.Email, *cmd.Password, *cmd.Gender, *cmd.Birth, *cmd.Country, *cmd.Language, cmd.Phone, t)

			cmd.Phone = &tc.phone

			usr := users.NewUser(*cmd.FirstName, *cmd.LastName, username, email, password, gender, birth, country, language, phone)
			usr.AggregateRoot = abstractions.NewAggregateRoot(cmd.Id)

			mockRepository.On("ExistsById", ctx, cmd.Id).Return(true, nil)
			mockRepository.On("GetById", ctx, cmd.Id).Return(usr, nil)

			resp, err := handler.HandleUpdate(ctx, cmd)

			require.Nil(t, resp)
			require.ErrorIs(t, err, tc.expectedErr)

			mockRepository.AssertExpectations(t)
		})
	}
}

func TestUserHandler_HandleUpdate_EmptyPhone(t *testing.T) {
	ctx := context.Background()

	mockFactory := new(MockFactory)
	mockRepository := new(MockRepository)

	handler := NewUserHandler(mockRepository, mockFactory)
	cmd := validUpdateUserCommand()
	username, email, password, gender, birth, country, language, _ := valueObjects(*cmd.UserName, *cmd.Email, *cmd.Password, *cmd.Gender, *cmd.Birth, *cmd.Country, *cmd.Language, nil, t)

	phoneStr := ""
	cmd.Phone = &phoneStr

	usr := users.NewUser(*cmd.FirstName, *cmd.LastName, username, email, password, gender, birth, country, language, nil)
	usr.AggregateRoot = abstractions.NewAggregateRoot(cmd.Id)

	mockRepository.On("ExistsById", ctx, cmd.Id).Return(true, nil)
	mockRepository.On("GetById", ctx, cmd.Id).Return(usr, nil)
	mockRepository.On("Update", ctx, mock.AnythingOfType("*users.User")).Return(usr, nil)

	resp, err := handler.HandleUpdate(ctx, cmd)

	require.NotNil(t, resp)
	require.NoError(t, err)
	assert.Nil(t, resp.Phone)

	mockRepository.AssertExpectations(t)
}

func TestUserHandler_HandleUpdate_InformationError(t *testing.T) {
	ctx := context.Background()

	mockFactory := new(MockFactory)
	mockRepository := new(MockRepository)

	handler := NewUserHandler(mockRepository, mockFactory)

	cmd := validUpdateUserCommand()
	username, email, password, gender, birth, country, language, phone := valueObjects(*cmd.UserName, *cmd.Email, *cmd.Password, *cmd.Gender, *cmd.Birth, *cmd.Country, *cmd.Language, cmd.Phone, t)

	information := strings.Repeat("a", 501)
	cmd.Information = &information

	usr := users.NewUser(*cmd.FirstName, *cmd.LastName, username, email, password, gender, birth, country, language, phone)
	usr.AggregateRoot = abstractions.NewAggregateRoot(cmd.Id)

	mockRepository.On("ExistsById", ctx, cmd.Id).Return(true, nil)
	mockRepository.On("GetById", ctx, cmd.Id).Return(usr, nil)

	resp, err := handler.HandleUpdate(ctx, cmd)

	require.Nil(t, resp)
	require.ErrorIs(t, err, users.ErrLongInformationUser)

	mockRepository.AssertExpectations(t)
}

func TestUserHandler_HandleUpdate_EmptyInformation(t *testing.T) {
	ctx := context.Background()

	mockFactory := new(MockFactory)
	mockRepository := new(MockRepository)

	handler := NewUserHandler(mockRepository, mockFactory)
	cmd := validUpdateUserCommand()
	username, email, password, gender, birth, country, language, phone := valueObjects(*cmd.UserName, *cmd.Email, *cmd.Password, *cmd.Gender, *cmd.Birth, *cmd.Country, *cmd.Language, nil, t)

	information := ""
	cmd.Information = &information

	usr := users.NewUser(*cmd.FirstName, *cmd.LastName, username, email, password, gender, birth, country, language, phone)
	usr.AggregateRoot = abstractions.NewAggregateRoot(cmd.Id)

	mockRepository.On("ExistsById", ctx, cmd.Id).Return(true, nil)
	mockRepository.On("GetById", ctx, cmd.Id).Return(usr, nil)
	mockRepository.On("Update", ctx, mock.AnythingOfType("*users.User")).Return(usr, nil)

	resp, err := handler.HandleUpdate(ctx, cmd)

	require.NotNil(t, resp)
	require.NoError(t, err)
	assert.Nil(t, resp.Information)

	mockRepository.AssertExpectations(t)
}

func TestUserHandler_HandleUpdate_ProfilePicError(t *testing.T) {
	ctx := context.Background()

	mockFactory := new(MockFactory)
	mockRepository := new(MockRepository)

	handler := NewUserHandler(mockRepository, mockFactory)

	cmd := validUpdateUserCommand()
	username, email, password, gender, birth, country, language, phone := valueObjects(*cmd.UserName, *cmd.Email, *cmd.Password, *cmd.Gender, *cmd.Birth, *cmd.Country, *cmd.Language, cmd.Phone, t)

	information := strings.Repeat("a", 501)
	cmd.Information = &information

	usr := users.NewUser(*cmd.FirstName, *cmd.LastName, username, email, password, gender, birth, country, language, phone)
	usr.AggregateRoot = abstractions.NewAggregateRoot(cmd.Id)

	mockRepository.On("ExistsById", ctx, cmd.Id).Return(true, nil)
	mockRepository.On("GetById", ctx, cmd.Id).Return(usr, nil)

	resp, err := handler.HandleUpdate(ctx, cmd)

	require.Nil(t, resp)
	require.ErrorIs(t, err, users.ErrLongInformationUser)

	mockRepository.AssertExpectations(t)
}

func TestUserHandler_HandleUpdate_EmptyProfilePic(t *testing.T) {
	ctx := context.Background()

	mockFactory := new(MockFactory)
	mockRepository := new(MockRepository)

	handler := NewUserHandler(mockRepository, mockFactory)
	cmd := validUpdateUserCommand()
	username, email, password, gender, birth, country, language, phone := valueObjects(*cmd.UserName, *cmd.Email, *cmd.Password, *cmd.Gender, *cmd.Birth, *cmd.Country, *cmd.Language, nil, t)

	profilePic := ""
	cmd.ProfilePic = &profilePic

	usr := users.NewUser(*cmd.FirstName, *cmd.LastName, username, email, password, gender, birth, country, language, phone)
	usr.AggregateRoot = abstractions.NewAggregateRoot(cmd.Id)

	mockRepository.On("ExistsById", ctx, cmd.Id).Return(true, nil)
	mockRepository.On("GetById", ctx, cmd.Id).Return(usr, nil)
	mockRepository.On("Update", ctx, mock.AnythingOfType("*users.User")).Return(usr, nil)

	resp, err := handler.HandleUpdate(ctx, cmd)

	require.NotNil(t, resp)
	require.NoError(t, err)
	assert.Nil(t, resp.ProfilePic)

	mockRepository.AssertExpectations(t)
}

func TestUserHandler_HandleUpdate_WebSiteError(t *testing.T) {
	ctx := context.Background()

	mockFactory := new(MockFactory)
	mockRepository := new(MockRepository)

	handler := NewUserHandler(mockRepository, mockFactory)

	cmd := validUpdateUserCommand()
	username, email, password, gender, birth, country, language, phone := valueObjects(*cmd.UserName, *cmd.Email, *cmd.Password, *cmd.Gender, *cmd.Birth, *cmd.Country, *cmd.Language, cmd.Phone, t)

	information := strings.Repeat("a", 501)
	cmd.Website = &information

	usr := users.NewUser(*cmd.FirstName, *cmd.LastName, username, email, password, gender, birth, country, language, phone)
	usr.AggregateRoot = abstractions.NewAggregateRoot(cmd.Id)

	mockRepository.On("ExistsById", ctx, cmd.Id).Return(true, nil)
	mockRepository.On("GetById", ctx, cmd.Id).Return(usr, nil)

	resp, err := handler.HandleUpdate(ctx, cmd)

	require.Nil(t, resp)
	require.ErrorIs(t, err, shared.ErrInvalidWebsite)

	mockRepository.AssertExpectations(t)
}

func TestUserHandler_HandleUpdate_EmptyWebSite(t *testing.T) {
	ctx := context.Background()

	mockFactory := new(MockFactory)
	mockRepository := new(MockRepository)

	handler := NewUserHandler(mockRepository, mockFactory)
	cmd := validUpdateUserCommand()
	username, email, password, gender, birth, country, language, phone := valueObjects(*cmd.UserName, *cmd.Email, *cmd.Password, *cmd.Gender, *cmd.Birth, *cmd.Country, *cmd.Language, nil, t)

	webSite := ""
	cmd.Website = &webSite

	usr := users.NewUser(*cmd.FirstName, *cmd.LastName, username, email, password, gender, birth, country, language, phone)
	usr.AggregateRoot = abstractions.NewAggregateRoot(cmd.Id)

	mockRepository.On("ExistsById", ctx, cmd.Id).Return(true, nil)
	mockRepository.On("GetById", ctx, cmd.Id).Return(usr, nil)
	mockRepository.On("Update", ctx, mock.AnythingOfType("*users.User")).Return(usr, nil)

	resp, err := handler.HandleUpdate(ctx, cmd)

	require.NotNil(t, resp)
	require.NoError(t, err)
	assert.Nil(t, resp.Website)

	mockRepository.AssertExpectations(t)
}

func TestUserHandler_HandleUpdate_UpdateError(t *testing.T) {
	ctx := context.Background()

	mockFactory := new(MockFactory)
	mockRepository := new(MockRepository)

	handler := NewUserHandler(mockRepository, mockFactory)

	cmd := validUpdateUserCommand()
	username, email, password, gender, birth, country, language, phone := valueObjects(*cmd.UserName, *cmd.Email, *cmd.Password, *cmd.Gender, *cmd.Birth, *cmd.Country, *cmd.Language, cmd.Phone, t)

	usr := users.NewUser(*cmd.FirstName, *cmd.LastName, username, email, password, gender, birth, country, language, phone)
	usr.AggregateRoot = abstractions.NewAggregateRoot(cmd.Id)

	mockRepository.On("ExistsById", ctx, cmd.Id).Return(true, nil)
	mockRepository.On("GetById", ctx, cmd.Id).Return(usr, nil)
	mockRepository.On("Update", ctx, usr).Return(nil, errors.New("mock repository error"))

	resp, err := handler.HandleUpdate(ctx, cmd)

	require.Nil(t, resp)
	require.Error(t, err)

	mockRepository.AssertExpectations(t)
}

func validUpdateUserCommand() commands.UpdateUserCommand {
	firstName := "John"
	lastName := "Doe"
	username := "johndoes"
	email := "john@doe.com"
	password := "5tr0nG.!"
	gender := "Male"
	birth := time.Now().AddDate(-20, -5, -2)
	country := "Bolivia"
	language := "Spanish"
	phone := "+591-77141516"
	information := "a lot of information"
	profilePic := "./images/profilpics/id/id.jpg"
	webSite := "https://www.github.com/carlosclavijo/"
	visibility := true
	cmd := commands.UpdateUserCommand{
		Id:          uuid.New(),
		FirstName:   &firstName,
		LastName:    &lastName,
		UserName:    &username,
		Email:       &email,
		Password:    &password,
		Gender:      &gender,
		Birth:       &birth,
		Country:     &country,
		Language:    &language,
		Phone:       &phone,
		Information: &information,
		ProfilePic:  &profilePic,
		Website:     &webSite,
		Visibility:  &visibility,
	}
	return cmd
}
