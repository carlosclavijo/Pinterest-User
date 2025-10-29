package handlers

import (
	"context"
	"errors"
	"github.com/carlosclavijo/Pinterest-User/internal/domain/shared"
	"github.com/carlosclavijo/Pinterest-User/internal/domain/user"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

type MockRepository struct {
	mock.Mock
}

type MockFactory struct {
	mock.Mock
}

func TestNewUserHandler(t *testing.T) {
	r := new(MockRepository)
	f := new(MockFactory)
	h := NewUserHandler(r, f)

	require.NotEmpty(t, h)
	require.Exactly(t, r, h.repository)
	require.Exactly(t, f, h.factory)
}

var errDbConnectionUser error = errors.New("db connection failed")

func (m *MockFactory) Create(firstName, lastName string, username shared.Username, email shared.Email, password shared.Password, gender shared.Gender, birth shared.BirthDate, country shared.Country, language shared.Language, phone *shared.Phone) (*users.User, error) {
	return nil, nil
}

func (m *MockRepository) GetAll(ctx context.Context) ([]*users.User, error) {
	args := m.Called(ctx)

	var usersList []*users.User
	if args.Get(0) != nil {
		usersList = args.Get(0).([]*users.User)
	}

	return usersList, args.Error(1)
}

func (m *MockRepository) GetList(ctx context.Context) ([]*users.User, error) {
	args := m.Called(ctx)

	var usersList []*users.User
	if args.Get(0) != nil {
		usersList = args.Get(0).([]*users.User)
	}

	return usersList, args.Error(1)
}

func (m *MockRepository) GetById(ctx context.Context, id uuid.UUID) (*users.User, error) {
	args := m.Called(ctx, id)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*users.User), args.Error(1)
}

func (m *MockRepository) GetByUsername(ctx context.Context, username string) (*users.User, error) {
	args := m.Called(ctx, username)

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
	args := m.Called(ctx, country)

	var usersList []*users.User
	if args.Get(0) != nil {
		usersList = args.Get(0).([]*users.User)
	}

	return usersList, args.Error(1)
}

func (m *MockRepository) GetListByLanguage(ctx context.Context, language string) ([]*users.User, error) {
	args := m.Called(ctx, language)

	var usersList []*users.User
	if args.Get(0) != nil {
		usersList = args.Get(0).([]*users.User)
	}

	return usersList, args.Error(1)
}

func (m *MockRepository) ExistsById(ctx context.Context, id uuid.UUID) (bool, error) {
	args := m.Called(ctx, id)

	var exists bool
	if v := args.Get(0); v != nil {
		exists, _ = v.(bool)
	}

	return exists, args.Error(1)
}

func (m *MockRepository) ExistsByUserName(ctx context.Context, username string) (bool, error) {
	args := m.Called(ctx, username)

	var exists bool
	if v := args.Get(0); v != nil {
		exists, _ = v.(bool)
	}

	return exists, args.Error(1)
}

func (m *MockRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	args := m.Called(ctx, email)

	var exists bool
	if v := args.Get(0); v != nil {
		exists, _ = v.(bool)
	}

	return exists, args.Error(1)
}

func (m *MockRepository) Create(ctx context.Context, u *users.User) (*users.User, error) {
	return nil, nil
}

func (m *MockRepository) Update(ctx context.Context, u *users.User) (*users.User, error) {
	return nil, nil
}

func (m *MockRepository) Delete(ctx context.Context, u *users.User) (*users.User, error) {
	return nil, nil
}

func listUsers() []*users.User {
	now := time.Now()
	phones := []string{"+591-7714151617", "+1-2025550143", "+49-3012345678", "+81-9012345678", "+61-412345678"}
	information := []string{"a lot of information", "lorem ipsum sit dolor amet", "dummy info", "subscribe to Carlos Clavijo Dev for more videos", "Give me a job, please"}
	profilePic := []string{"./image/user/id.jpg", "./image/user/id2.jpg", "./image/user/id3.jpg"}
	websites := []string{"https://www.github.com/carlosclavijo/", "https://www.x.com/CClavijoDev/", "https://www.linkedin.com/in/carlos-clavijo-b084a01b8", "https://leetcode.com/u/CarlosClavijo/", "https://youtube.com/@CarlosClavijoDev"}

	cases := []struct {
		id                                                     uuid.UUID
		firstName, lastName, username, email, password, gender string
		birth                                                  time.Time
		country, language                                      string
		phone, information, profilePic, webSite                *string
		visibility                                             bool
		lastLoginAt, createdAt, updatedAt                      time.Time
		deletedAt                                              *time.Time
	}{
		{uuid.New(), "John", "Doe", "johndoe", "john@doe.com", "5Tr0nG1.!", "Male", now.AddDate(-20, -10, -5), "Bolivia", "Spanish", &phones[0], &information[0], &profilePic[0], &websites[0], false, now, now, now, &now},
		{uuid.New(), "Jane", "Smith", "janesmith", "jane@smith.com", "S3cur3P@ss", "Female", now.AddDate(-25, 0, 0), "United States", "English", &phones[1], &information[1], &profilePic[1], &websites[1], true, now, now, now, nil},
		{uuid.New(), "Hans", "Müller", "hansmuller", "hans@muller.de", "Germ@ny123", "Male", now.AddDate(-30, 2, 15), "Germany", "German", &phones[2], &information[2], &profilePic[2], &websites[2], false, now, now, now, nil},
		{uuid.New(), "Akira", "Tanaka", "akira", "akira@tanaka.jp", "P@ssJPN99", "Male", now.AddDate(-28, -5, 12), "Japan", "Japanese", &phones[3], &information[3], nil, &websites[3], true, now, now, now, nil},
		{uuid.New(), "Olivia", "Brown", "oliviab", "olivia@brown.au", "Koal@123", "Female", now.AddDate(-22, 4, 3), "Australia", "English", &phones[4], &information[4], nil, &websites[4], false, now, now, now, nil},
		{uuid.New(), "Pierre", "Dupont", "pierred", "pierre@dupont.fr", "Bagu3tt3!", "Male", now.AddDate(-27, 6, 20), "France", "French", nil, nil, nil, nil, true, now, now, now, nil},
		{uuid.New(), "Giulia", "Rossi", "giuliar", "giulia@rossi.it", "T1ramisu!", "Female", now.AddDate(-24, -3, 8), "Italy", "Italian", nil, nil, nil, nil, true, now, now, now, nil},
		{uuid.New(), "Carlos", "Martínez", "cmartinez", "carlos@martinez.es", "Esp@na456", "Male", now.AddDate(-26, 9, 1), "South Korea", "Korean", nil, nil, nil, nil, false, now, now, now, nil},
		{uuid.New(), "Ana", "Silva", "anasilva", "ana@silva.br", "Braz1l!*", "Female", now.AddDate(-29, 1, 25), "Brazil", "Portuguese", nil, nil, nil, nil, true, now, now, now, nil},
		{uuid.New(), "Miguel", "Lopez", "miguell", "miguel@lopez.mx", "T@c0Power", "Male", now.AddDate(-23, 11, 9), "China", "Chinese", nil, nil, nil, nil, false, now, now, now, &now},
	}

	var usersList []*users.User
	for _, c := range cases {
		username, _ := shared.NewUsername(c.username)
		email, _ := shared.NewEmail(c.email)
		password, _ := shared.NewPassword(c.password)
		gender, _ := shared.ParseGender(c.gender)
		birth, _ := shared.NewBirthDate(c.birth)
		country, _ := shared.ParseCountry(c.country)
		language, _ := shared.ParseLanguage(c.language)
		phone, _ := shared.NewPhone(c.phone)
		website, _ := shared.NewWebSite(c.webSite)

		usr := users.NewUser(c.firstName, c.lastName, username, email, password, gender, birth, country, language, phone)

		_ = usr.ChangeInformation(c.information)
		usr.ChangeProfilePic(c.profilePic)
		usr.ChangeWebSite(website)
		usr.ChangeVisibility(c.visibility)

		time.Sleep(5 * time.Millisecond)

		usr.ChangeLastLoginAt()
		usr.ChangeUpdatedAt()

		if c.deletedAt != nil {
			time.Sleep(5 * time.Millisecond)
			_ = usr.Delete()
		}

		usersList = append(usersList, usr)
	}

	return usersList
}
