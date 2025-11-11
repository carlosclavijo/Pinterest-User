package handlers

import (
	"context"
	"errors"
	"github.com/carlosclavijo/Pinterest-Services/internal/domain/shared"
	"github.com/carlosclavijo/Pinterest-Services/internal/domain/user"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

type MockFactory struct {
	mock.Mock
}

type MockRepository struct {
	mock.Mock
}

var ErrDbFailureUser error = errors.New("db failure")

func TestNewUserHandler(t *testing.T) {
	factory := new(MockFactory)
	repository := new(MockRepository)
	handler := NewUserHandler(repository, factory)

	require.NotEmpty(t, handler)
	require.Exactly(t, factory, handler.factory)
	require.Exactly(t, repository, handler.repository)
}

func (m *MockFactory) Create(firstName, lastName string, usersName shared.Username, email shared.Email, password shared.Password, gender shared.Gender, birth shared.BirthDate, country shared.Country, language shared.Language, phone *shared.Phone) (*users.User, error) {
	args := m.Called(firstName, lastName, usersName, email, password, gender, birth, country, language, &phone)

	var result *users.User
	if v := args.Get(0); v != nil {
		result = v.(*users.User)
	}

	return result, args.Error(1)
}
func (m *MockRepository) GetAll(ctx context.Context) ([]*users.User, error) {
	return nil, nil
}

func (m *MockRepository) GetList(ctx context.Context) ([]*users.User, error) {
	return nil, nil
}

func (m *MockRepository) GetById(ctx context.Context, id uuid.UUID) (*users.User, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*users.User), args.Error(1)
}

func (m *MockRepository) GetByUsername(ctx context.Context, userName string) (*users.User, error) {
	args := m.Called(ctx, userName)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*users.User), args.Error(1)
}

func (m *MockRepository) GetByEmail(ctx context.Context, email string) (*users.User, error) {
	args := m.Called(ctx, email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*users.User), args.Error(1)
}

func (m *MockRepository) GetListByCountry(ctx context.Context, country string) ([]*users.User, error) {
	return nil, nil
}

func (m *MockRepository) GetListByLanguage(ctx context.Context, language string) ([]*users.User, error) {
	return nil, nil
}

func (m *MockRepository) ExistsById(ctx context.Context, id uuid.UUID) (bool, error) {
	args := m.Called(ctx, id)
	return args.Bool(0), args.Error(1)
}

func (m *MockRepository) ExistsByUserName(ctx context.Context, username string) (bool, error) {
	args := m.Called(ctx, username)
	return args.Bool(0), args.Error(1)
}

func (m *MockRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	args := m.Called(ctx, email)
	return args.Bool(0), args.Error(1)
}

func (m *MockRepository) Create(ctx context.Context, u *users.User) (*users.User, error) {
	args := m.Called(ctx, u)

	var result *users.User
	if v := args.Get(0); v != nil {
		result = v.(*users.User)
	}

	return result, args.Error(1)
}

func (m *MockRepository) Update(ctx context.Context, u *users.User) (*users.User, error) {
	args := m.Called(ctx, u)

	var result *users.User
	if v := args.Get(0); v != nil {
		result = v.(*users.User)
	}

	return result, args.Error(1)
}

func (m *MockRepository) Delete(ctx context.Context, u *users.User) (*users.User, error) {
	args := m.Called(ctx, u)

	var usr *users.User
	if v := args.Get(0); v != nil {
		usr, _ = v.(*users.User)
	}

	return usr, args.Error(1)
}

func valueObjects(username, email, password, gender string, birth time.Time, country, language string, phone *string, t *testing.T) (shared.Username, shared.Email, shared.Password, shared.Gender, shared.BirthDate, shared.Country, shared.Language, *shared.Phone) {
	usernameVo, err := shared.NewUsername(username)
	assert.NotEmpty(t, username)
	assert.NoError(t, err)

	emailVo, err := shared.NewEmail(email)
	assert.NotEmpty(t, email)
	assert.NoError(t, err)

	passwordVo, err := shared.NewPassword(password)
	assert.NotEmpty(t, password)
	assert.NoError(t, err)

	genderVo, err := shared.ParseGender(gender)
	assert.NotEmpty(t, gender)
	assert.NoError(t, err)

	birthVo, err := shared.NewBirthDate(birth)
	assert.NotEmpty(t, birth)
	assert.NoError(t, err)

	countryVo, err := shared.ParseCountry(country)
	assert.NotEmpty(t, country)
	assert.NoError(t, err)

	languageVo, err := shared.ParseLanguage(language)
	assert.NotEmpty(t, language)
	assert.NoError(t, err)

	phoneVo, err := shared.NewPhone(phone)
	if phone != nil {
		assert.NotEmpty(t, phone)
		assert.NoError(t, err)
	} else {
		assert.Nil(t, phoneVo)
		assert.NoError(t, err)
	}

	return usernameVo, emailVo, passwordVo, genderVo, birthVo, countryVo, languageVo, phoneVo
}
