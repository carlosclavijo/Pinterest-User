package handlers

import (
	"context"
	"errors"
	"github.com/carlosclavijo/Pinterest-Services/internal/application/user/commands"
	"github.com/carlosclavijo/Pinterest-Services/internal/application/user/dto"
	"github.com/carlosclavijo/Pinterest-Services/internal/domain/shared"
	"github.com/carlosclavijo/Pinterest-Services/internal/domain/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"testing"
	"time"
)

func TestUserHandler_HandleCreate(t *testing.T) {
	ctx := context.Background()

	mockRepository := new(MockRepository)
	mockFactory := new(MockFactory)

	handler := NewUserHandler(mockRepository, mockFactory)
	cmd := validCreateUserCommand()

	username, email, password, gender, birth, country, language, phone := valueObjects(cmd.Username, cmd.Email, cmd.Password, cmd.Gender, cmd.Birth, cmd.Country, cmd.Language, cmd.Phone, t)
	usr := users.NewUser(cmd.FirstName, cmd.LastName, username, email, password, gender, birth, country, language, phone)

	mockRepository.On("ExistsByUserName", ctx, cmd.Username).Return(false, nil)
	mockRepository.On("ExistsByEmail", ctx, cmd.Email).Return(false, nil)
	mockFactory.On("Create", cmd.FirstName, cmd.LastName, username, email,
		mock.MatchedBy(func(p shared.Password) bool {

			return bcrypt.CompareHashAndPassword([]byte(p.String()), []byte(cmd.Password)) == nil
		}), gender, birth, country, language, &phone).Return(usr, nil)

	mockRepository.On("Create", ctx, usr).Return(usr, nil)

	resp, err := handler.HandleCreate(ctx, cmd)

	require.NotNil(t, resp)
	require.IsType(t, &dto.UserResponse{}, resp)
	require.NoError(t, err)

	assert.NotNil(t, resp.Id)
	assert.Equal(t, cmd.FirstName, resp.FirstName)
	assert.Equal(t, cmd.LastName, resp.LastName)
	assert.Equal(t, cmd.Username, resp.Username)
	assert.Equal(t, cmd.Email, resp.Email)
	assert.Equal(t, cmd.Gender, resp.Gender)
	assert.Equal(t, cmd.Birth, resp.Birth)
	assert.Equal(t, cmd.Country, resp.Country)
	assert.Equal(t, cmd.Language, resp.Language)
	assert.Equal(t, cmd.Phone, resp.Phone)
	assert.Nil(t, resp.Information)
	assert.Nil(t, resp.ProfilePic)
	assert.Nil(t, resp.Website)
	assert.False(t, resp.Visibility)
	assert.NotNil(t, resp.LastLoginAt)
	assert.WithinDuration(t, time.Now(), resp.LastLoginAt, time.Second)
	assert.NotNil(t, resp.CreatedAt)
	assert.WithinDuration(t, time.Now(), resp.CreatedAt, time.Second)
	assert.NotNil(t, resp.UpdatedAt)
	assert.WithinDuration(t, time.Now(), resp.UpdatedAt, time.Second)
	assert.Nil(t, resp.DeletedAt)

	mockRepository.AssertExpectations(t)
	mockFactory.AssertExpectations(t)
}

func TestUserHandler_HandleCreate_UsernameError(t *testing.T) {
	ctx := context.Background()

	mockFactory := new(MockFactory)
	mockRepository := new(MockRepository)

	handler := NewUserHandler(mockRepository, mockFactory)
	cmd := validCreateUserCommand()

	cmd.Username = ""
	resp, err := handler.HandleCreate(ctx, cmd)

	assert.Nil(t, resp)
	assert.ErrorIs(t, err, shared.ErrEmptyUsername)

	cmd.Username = "lognusernamewithmorethanthirtycharacters"
	resp, err = handler.HandleCreate(ctx, cmd)

	assert.Nil(t, resp)
	assert.ErrorIs(t, err, shared.ErrLongUsername)

	cmd.Username = "no"
	resp, err = handler.HandleCreate(ctx, cmd)

	assert.Nil(t, resp)
	assert.ErrorIs(t, err, shared.ErrShortUsername)

	cmd.Username = "InvalidUsername!12.&"
	resp, err = handler.HandleCreate(ctx, cmd)

	assert.Nil(t, resp)
	assert.ErrorIs(t, err, shared.ErrInvalidUsername)
}

func TestUserHandler_HandleCreate_EmailError(t *testing.T) {
	ctx := context.Background()

	mockFactory := new(MockFactory)
	mockRepository := new(MockRepository)

	handler := NewUserHandler(mockRepository, mockFactory)
	cmd := validCreateUserCommand()

	cmd.Email = ""
	resp, err := handler.HandleCreate(ctx, cmd)

	assert.Nil(t, resp)
	assert.ErrorIs(t, err, shared.ErrEmptyEmail)

	cmd.Email = "invalidemail.com"
	resp, err = handler.HandleCreate(ctx, cmd)

	assert.Nil(t, resp)
	assert.ErrorIs(t, err, shared.ErrInvalidEmail)

	cmd.Email = "longlocalnametoolongthatcannotbealocalnamesowouldfailedeveryunittest@test.com"
	resp, err = handler.HandleCreate(ctx, cmd)

	assert.Nil(t, resp)
	assert.ErrorIs(t, err, shared.ErrLongLocalEmail)

	cmd.Email = "test@toomuchlongdomainnamethatwouldfailedeveryoneofthetesteventhehardestoneandidontknowwhatelsetoputheretoreachthe255oflongandevenmoretoreachthefailmaybeishouldtrytousechatgptforthisonebutiamtiredofthattoolnotbecauseisnnthelpfulbutbecauseaiistakingmyjobandinneedtofindajobtolivethatswhyiamdoingtheselivestreamspleaseifyouarereadingthisgivemeajobineedthemoneyipromiseiwillpushmylimitstowritegoodandqualityjobs.com"
	resp, err = handler.HandleCreate(ctx, cmd)

	assert.Nil(t, resp)
	assert.ErrorIs(t, err, shared.ErrLongDomainEmail)
}

func TestUserHandler_HandleCreate_PasswordError(t *testing.T) {
	ctx := context.Background()

	mockFactory := new(MockFactory)
	mockRepository := new(MockRepository)

	handler := NewUserHandler(mockRepository, mockFactory)
	cmd := validCreateUserCommand()

	cmd.Password = ""
	resp, err := handler.HandleCreate(ctx, cmd)

	assert.Nil(t, resp)
	assert.ErrorIs(t, err, shared.ErrEmptyPassword)

	cmd.Username = "lognusernamewithmorethanthirtycharacters"
	resp, err = handler.HandleCreate(ctx, cmd)

	assert.Nil(t, resp)
	assert.ErrorIs(t, err, shared.ErrLongUsername)

	cmd.Username = "no"
	resp, err = handler.HandleCreate(ctx, cmd)

	assert.Nil(t, resp)
	assert.ErrorIs(t, err, shared.ErrShortUsername)

	cmd.Username = "InvalidUsername!12.&"
	resp, err = handler.HandleCreate(ctx, cmd)

	assert.Nil(t, resp)
	assert.ErrorIs(t, err, shared.ErrInvalidUsername)
}

func TestUserHandler_HandleCreate_GenderError(t *testing.T) {
	ctx := context.Background()

	mockFactory := new(MockFactory)
	mockRepository := new(MockRepository)

	handler := NewUserHandler(mockRepository, mockFactory)
	cmd := validCreateUserCommand()

	cmd.Gender = "X"
	resp, err := handler.HandleCreate(ctx, cmd)

	assert.Nil(t, resp)
	assert.ErrorIs(t, err, shared.ErrNotAGender)
}

func TestUserHandler_HandleCreate_BirthError(t *testing.T) {
	ctx := context.Background()

	mockFactory := new(MockFactory)
	mockRepository := new(MockRepository)

	handler := NewUserHandler(mockRepository, mockFactory)
	cmd := validCreateUserCommand()

	cmd.Birth = time.Now().AddDate(1, 0, 0)
	resp, err := handler.HandleCreate(ctx, cmd)

	assert.Nil(t, resp)
	assert.ErrorIs(t, err, shared.ErrFutureDate)

	cmd.Birth = time.Now().AddDate(-10, 0, 0)
	resp, err = handler.HandleCreate(ctx, cmd)

	assert.Nil(t, resp)
	assert.ErrorIs(t, err, shared.ErrUnderTwelve)
}

func TestUserHandler_HandleCreate_CountryError(t *testing.T) {
	ctx := context.Background()

	mockFactory := new(MockFactory)
	mockRepository := new(MockRepository)

	handler := NewUserHandler(mockRepository, mockFactory)
	cmd := validCreateUserCommand()

	cmd.Country = "X"
	resp, err := handler.HandleCreate(ctx, cmd)

	assert.Nil(t, resp)
	assert.ErrorIs(t, err, shared.ErrNotACountry)
}

func TestUserHandler_HandleCreate_LanguageError(t *testing.T) {
	ctx := context.Background()

	mockFactory := new(MockFactory)
	mockRepository := new(MockRepository)

	handler := NewUserHandler(mockRepository, mockFactory)
	cmd := validCreateUserCommand()

	cmd.Language = "X"
	resp, err := handler.HandleCreate(ctx, cmd)

	assert.Nil(t, resp)
	assert.ErrorIs(t, err, shared.ErrNotALanguage)
}

func TestUserHandler_HandleCreate_PhoneError(t *testing.T) {
	ctx := context.Background()

	mockFactory := new(MockFactory)
	mockRepository := new(MockRepository)

	handler := NewUserHandler(mockRepository, mockFactory)
	cmd := validCreateUserCommand()

	newPhone := "a"
	cmd.Phone = &newPhone
	resp, err := handler.HandleCreate(ctx, cmd)

	assert.Nil(t, resp)
	assert.ErrorIs(t, err, shared.ErrNotNumericPhoneNumber)

	newPhone = "+591-77"
	cmd.Phone = &newPhone
	resp, err = handler.HandleCreate(ctx, cmd)

	assert.Nil(t, resp)
	assert.ErrorIs(t, err, shared.ErrShortPhoneNumber)

	newPhone = "+591-778878778787878787878"
	cmd.Phone = &newPhone
	resp, err = handler.HandleCreate(ctx, cmd)

	assert.Nil(t, resp)
	assert.ErrorIs(t, err, shared.ErrLongPhoneNumber)
}

func TestUserHandler_HandleCreate_ExistenceByUserNameCheck(t *testing.T) {
	ctx := context.Background()
	cases := []struct {
		name        string
		exist       bool
		repoError   error
		expectedErr error
	}{
		{"fail when username already exists", true, nil, users.ErrExistsUser},
		{"fail when repository returns error", false, ErrDbFailureUser, ErrDbFailureUser},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepository := new(MockRepository)
			mockFactory := new(MockFactory)
			handler := NewUserHandler(mockRepository, mockFactory)
			cmd := validCreateUserCommand()

			mockRepository.On("ExistsByUserName", ctx, cmd.Username).Return(tc.exist, tc.repoError)

			resp, err := handler.HandleCreate(ctx, cmd)

			assert.Nil(t, resp)
			assert.ErrorIs(t, err, tc.expectedErr)

			mockRepository.AssertExpectations(t)
			mockFactory.AssertExpectations(t)
		})
	}
}

func TestUserHandler_HandleCreate_ExistenceByEmailCheck(t *testing.T) {
	ctx := context.Background()
	cases := []struct {
		name        string
		exist       bool
		repoError   error
		expectedErr error
	}{
		{"fail when email already exists", true, nil, users.ErrExistsUser},
		{"fail when repository returns error", false, ErrDbFailureUser, ErrDbFailureUser},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepository := new(MockRepository)
			mockFactory := new(MockFactory)
			handler := NewUserHandler(mockRepository, mockFactory)
			cmd := validCreateUserCommand()

			mockRepository.On("ExistsByUserName", ctx, cmd.Username).Return(false, nil)
			mockRepository.On("ExistsByEmail", ctx, cmd.Email).Return(tc.exist, tc.repoError)

			resp, err := handler.HandleCreate(ctx, cmd)

			assert.Nil(t, resp)
			assert.ErrorIs(t, err, tc.expectedErr)

			mockRepository.AssertExpectations(t)
			mockFactory.AssertExpectations(t)
		})
	}
}

func TestUserHandler_HandleCreate_FactoryError(t *testing.T) {
	ctx := context.Background()

	mockFactory := new(MockFactory)
	mockRepository := new(MockRepository)

	handler := NewUserHandler(mockRepository, mockFactory)
	cmd := validCreateUserCommand()

	mockRepository.On("ExistsByUserName", ctx, cmd.Username).Return(false, nil)
	mockRepository.On("ExistsByEmail", ctx, cmd.Email).Return(false, nil)
	mockFactory.On("Create", cmd.FirstName, cmd.LastName, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("mock error"))

	resp, err := handler.HandleCreate(ctx, cmd)

	require.Nil(t, resp)
	require.Error(t, err)

	mockFactory.AssertExpectations(t)
	mockRepository.AssertExpectations(t)
}

func TestUserHandler_HandleCreate_CreateError(t *testing.T) {
	ctx := context.Background()

	mockFactory := new(MockFactory)
	mockRepository := new(MockRepository)

	handler := NewUserHandler(mockRepository, mockFactory)
	cmd := validCreateUserCommand()

	mockRepository.On("ExistsByUserName", ctx, cmd.Username).Return(false, nil)
	mockRepository.On("ExistsByEmail", ctx, cmd.Email).Return(false, nil)
	mockFactory.On("Create", cmd.FirstName, cmd.LastName, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&users.User{}, nil)
	mockRepository.On("Create", ctx, mock.AnythingOfType("*users.User")).Return(nil, errors.New("mock repository error"))

	resp, err := handler.HandleCreate(ctx, cmd)

	require.Nil(t, resp)
	require.Error(t, err)

	mockFactory.AssertExpectations(t)
	mockRepository.AssertExpectations(t)
}

func validCreateUserCommand() commands.CreateUserCommand {
	phoneStr := "+591-77141516"
	cmd := commands.CreateUserCommand{
		FirstName: "John",
		LastName:  "Doe",
		Username:  "johndoe",
		Email:     "john@doe.com",
		Password:  "5tr0nG!.",
		Gender:    "Male",
		Birth:     time.Now().AddDate(-20, -1, -2),
		Country:   "Bolivia",
		Language:  "Spanish",
		Phone:     &phoneStr,
	}
	return cmd
}
