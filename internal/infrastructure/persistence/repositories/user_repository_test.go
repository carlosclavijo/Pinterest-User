package repositories

import (
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/carlosclavijo/Pinterest-User/internal/domain/shared"
	"github.com/carlosclavijo/Pinterest-User/internal/domain/user"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"regexp"
	"testing"
	"time"
)

var (
	ctx         = context.Background()
	columns     = []string{"id", "first_name", "last_name", "username", "email", "password", "gender", "birth", "country", "language", "phone", "information", "profile_pic", "web_site", "visibility", "last_login_at", "created_at", "updated_at", "deleted_at"}
	ErrDatabase = errors.New("database is down")
)

type Case struct {
	name string
	*users.User
}

func TestNewUserRepository(t *testing.T) {
	db, _, err := sqlmock.New()

	require.NotNil(t, db)
	require.NoError(t, err)

	defer db.Close()

	repo := NewUserRepository(db)

	require.NotNil(t, repo)
	require.NotEmpty(t, repo)

	ur, ok := repo.(*userRepository)

	require.True(t, ok)
	assert.Equal(t, db, ur.DB)
}

func TestUserRepository_GetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()

	repo := NewUserRepository(db)
	cases := userCases()
	rows := sqlmock.NewRows(columns)

	for _, tc := range cases {
		var phone, information, profilePic, website *string
		if tc.Phone() != nil {
			p := tc.Phone().String()
			phone = &p
		}
		if tc.Information() != nil {
			i := tc.Information()
			information = i
		}
		if tc.ProfilePic() != nil {
			pp := tc.ProfilePic()
			profilePic = pp
		}
		if tc.WebSite() != nil {
			ws := tc.WebSite().String()
			website = &ws
		}

		rows.AddRow(
			tc.Id(), tc.FirstName(), tc.LastName(), tc.Username().String(), tc.Email().String(), tc.Password().String(), tc.Gender().String(), tc.Birth().Time(), tc.Country().String(),
			tc.Language().String(), phone, information, profilePic, website, tc.Visibility(), tc.LastLoginAt(), tc.CreatedAt(), tc.UpdatedAt(), tc.DeletedAt(),
		)
	}

	mock.ExpectQuery(QueryGetAllUsers).WillReturnRows(rows)

	usersList, err := repo.GetAll(ctx)

	require.NotNil(t, usersList)
	require.NoError(t, err)

	assert.Len(t, usersList, len(cases))

	for i, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			testCases(t, tc, usersList[i])
		})
	}

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_GetAll_QueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()

	repo := NewUserRepository(db)

	mock.ExpectQuery(QueryGetAllUsers).WillReturnError(ErrDatabase)

	usersList, err := repo.GetAll(ctx)

	require.Nil(t, usersList)
	require.Error(t, err)

	assert.ErrorIs(t, err, ErrQueryUser)
	assert.ErrorIs(t, err, ErrDatabase)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_GetAll_ScanError(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()

	repo := NewUserRepository(db)
	rows := sqlmock.NewRows([]string{"id"}).AddRow(uuid.New())

	mock.ExpectQuery(QueryGetAllUsers).WillReturnRows(rows)

	usersList, err := repo.GetAll(ctx)

	require.Nil(t, usersList)
	require.Error(t, err)

	assert.ErrorIs(t, err, ErrScanUser)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_GetAll_NewUserError(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()

	repo := NewUserRepository(db)
	rows := sqlmock.NewRows(columns).AddRow(uuid.New(), "", "", "", "", "", "", time.Now(), "", "", nil, nil, nil, nil, false, time.Now(), time.Now(), time.Now(), nil)

	mock.ExpectQuery(QueryGetAllUsers).WillReturnRows(rows)

	usersList, err := repo.GetAll(ctx)

	require.Nil(t, usersList)
	require.Error(t, err)

	assert.ErrorIs(t, err, ErrConcatenatingUser)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_GetList(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()

	repo := NewUserRepository(db)
	cases := userCases()
	rows := sqlmock.NewRows(columns)

	for _, tc := range cases {
		var phone, information, profilePic, website *string
		if tc.Phone() != nil {
			p := tc.Phone().String()
			phone = &p
		}
		if tc.Information() != nil {
			i := tc.Information()
			information = i
		}
		if tc.ProfilePic() != nil {
			pp := tc.ProfilePic()
			profilePic = pp
		}
		if tc.WebSite() != nil {
			ws := tc.WebSite().String()
			website = &ws
		}

		rows.AddRow(
			tc.Id(), tc.FirstName(), tc.LastName(), tc.Username().String(), tc.Email().String(), tc.Password().String(), tc.Gender().String(), tc.Birth().Time(), tc.Country().String(),
			tc.Language().String(), phone, information, profilePic, website, tc.Visibility(), tc.LastLoginAt(), tc.CreatedAt(), tc.UpdatedAt(), tc.DeletedAt(),
		)
	}

	mock.ExpectQuery(QueryGetListUsers).WillReturnRows(rows)

	usersList, err := repo.GetList(ctx)

	require.NotNil(t, usersList)
	require.NoError(t, err)

	assert.Len(t, usersList, len(cases))

	for i, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			testCases(t, tc, usersList[i])
		})
	}

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_GetList_QueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()

	repo := NewUserRepository(db)

	mock.ExpectQuery(QueryGetListUsers).WillReturnError(ErrDatabase)

	usersList, err := repo.GetList(ctx)

	require.Nil(t, usersList)
	require.Error(t, err)

	assert.ErrorIs(t, err, ErrQueryUser)
	assert.ErrorIs(t, err, ErrDatabase)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_GetList_ScanError(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()

	repo := NewUserRepository(db)
	rows := sqlmock.NewRows([]string{"id"}).AddRow(uuid.New())

	mock.ExpectQuery(QueryGetListUsers).WillReturnRows(rows)

	usersList, err := repo.GetList(ctx)

	require.Nil(t, usersList)
	require.Error(t, err)
	assert.ErrorIs(t, err, ErrScanUser)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_GetList_NewUserError(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()

	repo := NewUserRepository(db)
	rows := sqlmock.NewRows(columns).AddRow(uuid.New(), "", "", "", "", "", "", time.Now(), "", "", nil, nil, nil, nil, false, time.Now(), time.Now(), time.Now(), nil)

	mock.ExpectQuery(QueryGetListUsers).WillReturnRows(rows)

	usersList, err := repo.GetList(ctx)

	require.Nil(t, usersList)
	require.Error(t, err)
	assert.ErrorIs(t, err, ErrConcatenatingUser)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_GetById(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()

	repo := NewUserRepository(db)
	tc := userCases()[0]

	cols := append([]string(nil), columns...)
	cols = cols[1:]

	var phone, information, profilePic, website *string
	if tc.Phone() != nil {
		p := tc.Phone().String()
		phone = &p
	}
	if tc.Information() != nil {
		i := tc.Information()
		information = i
	}
	if tc.ProfilePic() != nil {
		pp := tc.ProfilePic()
		profilePic = pp
	}
	if tc.WebSite() != nil {
		ws := tc.WebSite().String()
		website = &ws
	}

	rows := sqlmock.NewRows(cols).AddRow(
		tc.FirstName(), tc.LastName(), tc.Username().String(), tc.Email().String(), tc.Password().String(), tc.Gender(), tc.Birth().Time(), tc.Country(),
		tc.Language(), phone, information, profilePic, website, tc.Visibility(), tc.LastLoginAt(), tc.CreatedAt(), tc.UpdatedAt(), tc.DeletedAt(),
	)

	mock.ExpectQuery(regexp.QuoteMeta(QueryGetUserById)).WithArgs(tc.Id()).WillReturnRows(rows)

	usr, err := repo.GetById(ctx, tc.Id())

	testCases(t, tc, usr)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_GetById_QueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()

	repo := NewUserRepository(db)
	id := uuid.New()

	mock.ExpectQuery(regexp.QuoteMeta(QueryGetUserById)).WithArgs(id).WillReturnError(ErrDatabase)

	usr, err := repo.GetById(ctx, id)

	require.Nil(t, usr)
	require.Error(t, err)

	assert.ErrorIs(t, err, ErrQueryUser)
	assert.ErrorIs(t, err, ErrDatabase)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_GetById_NewUserError(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()

	repo := NewUserRepository(db)
	id := uuid.New()
	cols := append([]string(nil), columns...)
	cols = cols[1:]

	rows := sqlmock.NewRows(cols).AddRow("", "", "", "", "", "", time.Now(), "", "", nil, nil, nil, nil, false, time.Now(), time.Now(), time.Now(), nil)

	mock.ExpectQuery(regexp.QuoteMeta(QueryGetUserById)).WithArgs(id).WillReturnRows(rows)

	usr, err := repo.GetById(ctx, id)

	require.Nil(t, usr)
	require.Error(t, err)

	assert.ErrorIs(t, err, ErrConcatenatingUser)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_GetByUsername(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()

	repo := NewUserRepository(db)
	tc := userCases()[0]

	cols := append([]string(nil), columns...)
	cols = append(cols[:3], cols[4:]...)

	var phone, information, profilePic, website *string
	if tc.Phone() != nil {
		p := tc.Phone().String()
		phone = &p
	}
	if tc.Information() != nil {
		i := tc.Information()
		information = i
	}
	if tc.ProfilePic() != nil {
		pp := tc.ProfilePic()
		profilePic = pp
	}
	if tc.WebSite() != nil {
		ws := tc.WebSite().String()
		website = &ws
	}

	rows := sqlmock.NewRows(cols).AddRow(
		tc.Id(), tc.FirstName(), tc.LastName(), tc.Email().String(), tc.Password().String(), tc.Gender(), tc.Birth().Time(), tc.Country(),
		tc.Language(), phone, information, profilePic, website, tc.Visibility(), tc.LastLoginAt(), tc.CreatedAt(), tc.UpdatedAt(), tc.DeletedAt(),
	)

	mock.ExpectQuery(regexp.QuoteMeta(QueryGetUserByUsername)).WithArgs(tc.Username().String()).WillReturnRows(rows)

	usr, err := repo.GetByUsername(ctx, tc.Username().String())

	testCases(t, tc, usr)
	require.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_GetByUsername_QueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()

	repo := NewUserRepository(db)
	username := "username"

	mock.ExpectQuery(regexp.QuoteMeta(QueryGetUserByUsername)).WithArgs(username).WillReturnError(ErrDatabase)

	usr, err := repo.GetByUsername(ctx, username)

	require.Nil(t, usr)
	require.Error(t, err)

	assert.ErrorIs(t, err, ErrQueryUser)
	assert.ErrorIs(t, err, ErrDatabase)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_GetByUsername_NewUserError(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()

	repo := NewUserRepository(db)
	username := "username"
	cols := append([]string(nil), columns...)
	cols = append(cols[:3], cols[4:]...)

	rows := sqlmock.NewRows(cols).AddRow(uuid.New(), "", "", "", "", "", time.Now(), "", "", nil, nil, nil, nil, false, time.Now(), time.Now(), time.Now(), nil)

	mock.ExpectQuery(regexp.QuoteMeta(QueryGetUserByUsername)).WithArgs(username).WillReturnRows(rows)

	usr, err := repo.GetByUsername(ctx, username)

	require.Nil(t, usr)
	require.Error(t, err)

	assert.ErrorIs(t, err, ErrConcatenatingUser)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_GetByEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()

	repo := NewUserRepository(db)
	tc := userCases()[0]

	cols := append([]string(nil), columns...)
	cols = append(cols[:4], cols[5:]...)

	var phone, information, profilePic, website *string
	if tc.Phone() != nil {
		p := tc.Phone().String()
		phone = &p
	}
	if tc.Information() != nil {
		i := tc.Information()
		information = i
	}
	if tc.ProfilePic() != nil {
		pp := tc.ProfilePic()
		profilePic = pp
	}
	if tc.WebSite() != nil {
		ws := tc.WebSite().String()
		website = &ws
	}

	rows := sqlmock.NewRows(cols).AddRow(
		tc.Id(), tc.FirstName(), tc.LastName(), tc.Username().String(), tc.Password().String(), tc.Gender(), tc.Birth().Time(), tc.Country(),
		tc.Language(), phone, information, profilePic, website, tc.Visibility(), tc.LastLoginAt(), tc.CreatedAt(), tc.UpdatedAt(), tc.DeletedAt(),
	)

	mock.ExpectQuery(regexp.QuoteMeta(QueryGetUserByEmail)).WithArgs(tc.Email().String()).WillReturnRows(rows)

	usr, err := repo.GetByEmail(ctx, tc.Email().String())

	testCases(t, tc, usr)
	require.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_GetByEmail_QueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()

	repo := NewUserRepository(db)
	email := "valid@email.com"

	mock.ExpectQuery(regexp.QuoteMeta(QueryGetUserByEmail)).WithArgs(email).WillReturnError(ErrDatabase)

	usr, err := repo.GetByEmail(ctx, email)

	require.Nil(t, usr)
	require.Error(t, err)

	assert.ErrorIs(t, err, ErrQueryUser)
	assert.ErrorIs(t, err, ErrDatabase)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_GetByEmail_NewUserError(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()

	repo := NewUserRepository(db)
	email := "valid@email.com"
	cols := append([]string(nil), columns...)
	cols = append(cols[:4], cols[5:]...)

	rows := sqlmock.NewRows(cols).AddRow(uuid.New(), "", "", "", "", "", time.Now(), "", "", nil, nil, nil, nil, false, time.Now(), time.Now(), time.Now(), nil)

	mock.ExpectQuery(regexp.QuoteMeta(QueryGetUserByEmail)).WithArgs(email).WillReturnRows(rows)

	usr, err := repo.GetByEmail(ctx, email)

	require.Nil(t, usr)
	require.Error(t, err)

	assert.ErrorIs(t, err, ErrConcatenatingUser)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_GetListByCountry(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()

	repo := NewUserRepository(db)
	cases := userCases()
	country := "Bolivia"
	cols := append([]string(nil), columns...)
	cols = append(cols[:9], cols[10:]...)
	rows := sqlmock.NewRows(cols)

	for _, tc := range cases {
		if tc.Country().String() != country {
			continue
		}
		var phone, information, profilePic, website *string
		if tc.Phone() != nil {
			p := tc.Phone().String()
			phone = &p
		}
		if tc.Information() != nil {
			i := tc.Information()
			information = i
		}
		if tc.ProfilePic() != nil {
			pp := tc.ProfilePic()
			profilePic = pp
		}
		if tc.WebSite() != nil {
			ws := tc.WebSite().String()
			website = &ws
		}

		rows.AddRow(
			tc.Id(), tc.FirstName(), tc.LastName(), tc.Username().String(), tc.Email().String(), tc.Password().String(), tc.Gender().String(), tc.Birth().Time(),
			tc.Language().String(), phone, information, profilePic, website, tc.Visibility(), tc.LastLoginAt(), tc.CreatedAt(), tc.UpdatedAt(), tc.DeletedAt(),
		)
	}

	mock.ExpectQuery(regexp.QuoteMeta(QueryGetUsersByCountry)).WithArgs(country).WillReturnRows(rows)

	usersList, err := repo.GetListByCountry(ctx, country)

	require.NotNil(t, usersList)
	require.NoError(t, err)

	for i, tc := range cases {
		if tc.Country().String() != country {
			continue
		}
		t.Run(tc.name, func(t *testing.T) {
			testCases(t, tc, usersList[i])
		})
	}

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_GetListByCountry_QueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()

	repo := NewUserRepository(db)
	country := "Bolivia"

	mock.ExpectQuery(regexp.QuoteMeta(QueryGetUsersByCountry)).WithArgs(country).WillReturnError(ErrDatabase)

	usersList, err := repo.GetListByCountry(ctx, country)

	require.Nil(t, usersList)
	require.Error(t, err)

	assert.ErrorIs(t, err, ErrQueryUser)
	assert.ErrorIs(t, err, ErrDatabase)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_GetListByCountry_ScanError(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()

	repo := NewUserRepository(db)
	country := "Bolivia"

	newColumns := []string{"id"}
	rows := sqlmock.NewRows(newColumns).AddRow(uuid.New())

	mock.ExpectQuery(regexp.QuoteMeta(QueryGetUsersByCountry)).WithArgs(country).WillReturnRows(rows)

	usersList, err := repo.GetListByCountry(ctx, country)

	require.Nil(t, usersList)
	require.Error(t, err)

	assert.ErrorIs(t, err, ErrScanUser)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_GetListByCountry_NewUserError(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()

	repo := NewUserRepository(db)
	country := "Bolivia"
	cols := append([]string(nil), columns...)
	cols = append(cols[:9], cols[10:]...)
	rows := sqlmock.NewRows(cols).AddRow(uuid.New(), "", "", "", "", "", "", time.Now(), "", nil, nil, nil, nil, false, time.Now(), time.Now(), time.Now(), nil)

	mock.ExpectQuery(regexp.QuoteMeta(QueryGetUsersByCountry)).WithArgs(country).WillReturnRows(rows)

	usersList, err := repo.GetListByCountry(ctx, country)

	require.Nil(t, usersList)
	require.Error(t, err)

	assert.ErrorIs(t, err, ErrConcatenatingUser)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_GetListByLanguage(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()

	repo := NewUserRepository(db)
	cases := userCases()
	language := "Spanish"
	cols := append([]string(nil), columns...)
	cols = append(cols[:10], cols[11:]...)
	rows := sqlmock.NewRows(cols)

	for _, tc := range cases {
		if tc.Language().String() != language {
			continue
		}
		var phone, information, profilePic, website *string
		if tc.Phone() != nil {
			p := tc.Phone().String()
			phone = &p
		}
		if tc.Information() != nil {
			i := tc.Information()
			information = i
		}
		if tc.ProfilePic() != nil {
			pp := tc.ProfilePic()
			profilePic = pp
		}
		if tc.WebSite() != nil {
			ws := tc.WebSite().String()
			website = &ws
		}

		rows.AddRow(
			tc.Id(), tc.FirstName(), tc.LastName(), tc.Username().String(), tc.Email().String(), tc.Password().String(), tc.Gender().String(), tc.Birth().Time(),
			tc.Country().String(), phone, information, profilePic, website, tc.Visibility(), tc.LastLoginAt(), tc.CreatedAt(), tc.UpdatedAt(), tc.DeletedAt(),
		)
	}

	mock.ExpectQuery(regexp.QuoteMeta(QueryGetUsersByLanguage)).WithArgs(language).WillReturnRows(rows)

	usersList, err := repo.GetListByLanguage(ctx, language)

	require.NotNil(t, usersList)
	require.NoError(t, err)

	for i, tc := range cases {
		if tc.Language().String() != language {
			continue
		}
		t.Run(tc.name, func(t *testing.T) {
			testCases(t, tc, usersList[i])
		})
	}

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_GetListByLanguage_QueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()

	repo := NewUserRepository(db)
	language := "Spanish"

	mock.ExpectQuery(regexp.QuoteMeta(QueryGetUsersByLanguage)).WithArgs(language).WillReturnError(ErrDatabase)

	usersList, err := repo.GetListByLanguage(ctx, language)

	require.Nil(t, usersList)
	require.Error(t, err)

	assert.ErrorIs(t, err, ErrQueryUser)
	assert.ErrorIs(t, err, ErrDatabase)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_GetListByLanguage_ScanError(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()

	repo := NewUserRepository(db)
	language := "Spanish"

	newColumns := []string{"id"}
	rows := sqlmock.NewRows(newColumns).AddRow(uuid.New())

	mock.ExpectQuery(regexp.QuoteMeta(QueryGetUsersByLanguage)).WithArgs(language).WillReturnRows(rows)

	usersList, err := repo.GetListByLanguage(ctx, language)

	require.Nil(t, usersList)
	require.Error(t, err)

	assert.ErrorIs(t, err, ErrScanUser)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_GetListByLanguage_NewUserError(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()

	repo := NewUserRepository(db)
	language := "Spanish"
	cols := append([]string(nil), columns...)
	cols = append(cols[:10], cols[11:]...)
	rows := sqlmock.NewRows(cols).AddRow(uuid.New(), "", "", "", "", "", "", time.Now(), "", nil, nil, nil, nil, false, time.Now(), time.Now(), time.Now(), nil)

	mock.ExpectQuery(regexp.QuoteMeta(QueryGetUsersByLanguage)).WithArgs(language).WillReturnRows(rows)

	usersList, err := repo.GetListByLanguage(ctx, language)

	require.Nil(t, usersList)
	require.Error(t, err)

	assert.ErrorIs(t, err, ErrConcatenatingUser)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_ExistsById(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()

	repo := NewUserRepository(db)
	id := uuid.New()

	rows := sqlmock.NewRows([]string{"exist"}).AddRow(true)

	mock.ExpectQuery(regexp.QuoteMeta(QueryExistUserById)).WithArgs(id).WillReturnRows(rows)

	exists, err := repo.ExistsById(ctx, id)

	require.True(t, exists)
	require.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_ExistsById_QueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()

	repo := NewUserRepository(db)
	id := uuid.New()

	mock.ExpectQuery(regexp.QuoteMeta(QueryExistUserById)).WithArgs(id).WillReturnError(ErrDatabase)

	exists, err := repo.ExistsById(ctx, id)

	require.False(t, exists)
	require.Error(t, err)

	assert.ErrorIs(t, err, ErrQueryUser)
	assert.ErrorIs(t, err, ErrDatabase)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_ExistsByUserName(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()

	repo := NewUserRepository(db)
	username := "username"

	rows := sqlmock.NewRows([]string{"exist"}).AddRow(true)

	mock.ExpectQuery(regexp.QuoteMeta(QueryExistUserByUsername)).WithArgs(username).WillReturnRows(rows)

	exists, err := repo.ExistsByUserName(ctx, username)

	require.True(t, exists)
	require.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_ExistsByUserName_QueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()

	repo := NewUserRepository(db)
	username := "username"

	mock.ExpectQuery(regexp.QuoteMeta(QueryExistUserByUsername)).WithArgs(username).WillReturnError(ErrDatabase)

	exists, err := repo.ExistsByUserName(ctx, username)

	require.False(t, exists)
	require.Error(t, err)

	assert.ErrorIs(t, err, ErrQueryUser)
	assert.ErrorIs(t, err, ErrDatabase)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_ExistsByEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()

	repo := NewUserRepository(db)
	email := "valid@email.com"

	rows := sqlmock.NewRows([]string{"exist"}).AddRow(true)

	mock.ExpectQuery(regexp.QuoteMeta(QueryExistUserByEmail)).WithArgs(email).WillReturnRows(rows)

	exists, err := repo.ExistsByEmail(ctx, email)

	require.True(t, exists)
	require.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_ExistsByEmail_QueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()

	repo := NewUserRepository(db)
	email := "valid@email.com"

	mock.ExpectQuery(regexp.QuoteMeta(QueryExistUserByEmail)).WithArgs(email).WillReturnError(ErrDatabase)

	exists, err := repo.ExistsByEmail(ctx, email)

	require.False(t, exists)
	require.Error(t, err)

	assert.ErrorIs(t, err, ErrQueryUser)
	assert.ErrorIs(t, err, ErrDatabase)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()

	repo := NewUserRepository(db)
	tc := userCases()[0]

	rows := sqlmock.NewRows(columns).AddRow(
		tc.Id(), tc.FirstName(), tc.LastName(), tc.Username().String(), tc.Email().String(), tc.Password().String(), tc.Gender(), tc.Birth().Time(), tc.Country(), tc.Language(),
		tc.Phone().String(), tc.Information(), tc.ProfilePic(), tc.WebSite().String(), tc.Visibility(), tc.LastLoginAt(), tc.CreatedAt(), tc.UpdatedAt(), tc.DeletedAt(),
	)

	mock.ExpectQuery(regexp.QuoteMeta(QueryCreateUser)).WithArgs(
		tc.Id(), tc.FirstName(), tc.LastName(), tc.Username().String(), tc.Email().String(), tc.Password().String(), tc.Gender(), tc.Birth().Time(), tc.Country(),
		tc.Language(), tc.Phone().String(), tc.Visibility(), tc.LastLoginAt(), tc.CreatedAt(), tc.UpdatedAt(),
	).WillReturnRows(rows)

	usr, err := repo.Create(context.Background(), tc.User)

	require.NotNil(t, usr)
	require.NoError(t, err)
	testCases(t, tc, usr)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_Create_QueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()

	repo := NewUserRepository(db)
	tc := userCases()[0]

	mock.ExpectQuery(regexp.QuoteMeta(QueryCreateUser)).WithArgs(
		tc.Id(), tc.FirstName(), tc.LastName(), tc.Username().String(), tc.Email().String(), tc.Password().String(),
		tc.Gender(), tc.Birth().Time(), tc.Country(), tc.Language(), tc.Phone().String(), tc.Visibility(), tc.LastLoginAt(), tc.CreatedAt(), tc.UpdatedAt(),
	).WillReturnError(ErrDatabase)

	usr, err := repo.Create(context.Background(), tc.User)

	require.Nil(t, usr)
	require.Error(t, err)

	assert.ErrorIs(t, err, ErrQueryUser)
	assert.ErrorIs(t, err, ErrDatabase)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_Create_NewUserError(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()

	repo := NewUserRepository(db)
	tc := userCases()[0]

	rows := sqlmock.NewRows(columns).AddRow(uuid.Nil, "", "", "", "", "", "", time.Now(), "", "", nil, nil, nil, nil, false, time.Now(), time.Now(), time.Now(), nil)

	mock.ExpectQuery(regexp.QuoteMeta(QueryCreateUser)).WithArgs(
		tc.Id(), tc.FirstName(), tc.LastName(), tc.Username().String(), tc.Email().String(),
		tc.Password().String(), tc.Gender(), tc.Birth().Time(), tc.Country(),
		tc.Language(), tc.Phone().String(), tc.Visibility(), tc.LastLoginAt(), tc.CreatedAt(), tc.UpdatedAt(),
	).WillReturnRows(rows)

	usr, err := repo.Create(context.Background(), tc.User)

	require.Nil(t, usr)
	require.Error(t, err)

	assert.ErrorIs(t, err, ErrConcatenatingUser)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_Update(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()

	repo := NewUserRepository(db)
	tc := userCases()[0]

	rows := sqlmock.NewRows(columns).AddRow(
		tc.Id(), tc.FirstName(), tc.LastName(), tc.Username().String(), tc.Email().String(), tc.Password().String(), tc.Gender(), tc.Birth().Time(), tc.Country(), tc.Language(),
		tc.Phone().String(), tc.Information(), tc.ProfilePic(), tc.WebSite().String(), tc.Visibility(), tc.LastLoginAt(), tc.CreatedAt(), tc.UpdatedAt(), tc.DeletedAt(),
	)

	mock.ExpectQuery(regexp.QuoteMeta(QueryUpdateUser)).WithArgs(
		tc.Id(), tc.FirstName(), tc.LastName(), tc.Username().String(), tc.Email().String(), tc.Password().String(), tc.Gender(), tc.Birth().Time(), tc.Country(),
		tc.Language(), tc.Phone().String(), tc.Information(), tc.ProfilePic(), tc.WebSite().String(), tc.Visibility(), tc.LastLoginAt(), tc.UpdatedAt(),
	).WillReturnRows(rows)

	usr, err := repo.Update(context.Background(), tc.User)

	require.NotNil(t, usr)
	require.NoError(t, err)

	testCases(t, tc, usr)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_Update_QueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()

	repo := NewUserRepository(db)
	tc := userCases()[0]

	mock.ExpectQuery(regexp.QuoteMeta(QueryUpdateUser)).WithArgs(
		tc.Id(), tc.FirstName(), tc.LastName(), tc.Username().String(), tc.Email().String(), tc.Password().String(), tc.Gender(), tc.Birth().Time(), tc.Country(),
		tc.Language(), tc.Phone().String(), tc.Information(), tc.ProfilePic(), tc.WebSite().String(), tc.Visibility(), tc.LastLoginAt(), tc.UpdatedAt(),
	).WillReturnError(ErrDatabase)

	usr, err := repo.Update(context.Background(), tc.User)

	require.Nil(t, usr)
	require.Error(t, err)

	assert.ErrorIs(t, err, ErrDatabase)
	assert.ErrorIs(t, err, ErrQueryUser)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_Update_NewUserError(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()

	repo := NewUserRepository(db)
	tc := userCases()[0]

	rows := sqlmock.NewRows(columns).AddRow(uuid.Nil, "", "", "", "", "", "", time.Now(), "", "", nil, nil, nil, nil, false, time.Now(), time.Now(), time.Now(), nil)

	mock.ExpectQuery(regexp.QuoteMeta(QueryUpdateUser)).WithArgs(
		tc.Id(), tc.FirstName(), tc.LastName(), tc.Username().String(), tc.Email().String(), tc.Password().String(), tc.Gender(), tc.Birth().Time(), tc.Country(),
		tc.Language(), tc.Phone().String(), tc.Information(), tc.ProfilePic(), tc.WebSite().String(), tc.Visibility(), tc.LastLoginAt(), tc.UpdatedAt(),
	).WillReturnRows(rows)

	usr, err := repo.Update(context.Background(), tc.User)

	require.Nil(t, usr)
	require.Error(t, err)

	assert.ErrorIs(t, err, ErrConcatenatingUser)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_Delete(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()

	repo := NewUserRepository(db)
	tc := userCases()[0]

	rows := sqlmock.NewRows(columns).AddRow(
		tc.Id(), tc.FirstName(), tc.LastName(), tc.Username().String(), tc.Email().String(), tc.Password().String(), tc.Gender(), tc.Birth().Time(), tc.Country(), tc.Language(),
		tc.Phone().String(), tc.Information(), tc.ProfilePic(), tc.WebSite().String(), tc.Visibility(), tc.LastLoginAt(), tc.CreatedAt(), tc.UpdatedAt(), tc.DeletedAt(),
	)

	mock.ExpectQuery(regexp.QuoteMeta(QueryDeleteUser)).WithArgs(tc.Id(), tc.DeletedAt()).WillReturnRows(rows)

	usr, err := repo.Delete(context.Background(), tc.User)

	require.NotNil(t, usr)
	require.NoError(t, err)

	testCases(t, tc, usr)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_Delete_QueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()

	repo := NewUserRepository(db)
	tc := userCases()[0]

	mock.ExpectQuery(regexp.QuoteMeta(QueryDeleteUser)).WithArgs(tc.Id(), tc.DeletedAt()).WillReturnError(ErrDatabase)

	usr, err := repo.Delete(context.Background(), tc.User)

	require.Nil(t, usr)
	require.Error(t, err)

	assert.ErrorIs(t, err, ErrDatabase)
	assert.ErrorIs(t, err, ErrQueryUser)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_Delete_NewUserError(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()

	repo := NewUserRepository(db)
	tc := userCases()[0]

	rows := sqlmock.NewRows(columns).AddRow(uuid.Nil, "", "", "", "", "", "", time.Now(), "", "", nil, nil, nil, nil, false, time.Now(), time.Now(), time.Now(), nil)

	mock.ExpectQuery(regexp.QuoteMeta(QueryDeleteUser)).WithArgs(tc.Id(), tc.DeletedAt()).WillReturnRows(rows)

	usr, err := repo.Delete(context.Background(), tc.User)

	require.Nil(t, usr)
	require.Error(t, err)

	assert.ErrorIs(t, err, ErrConcatenatingUser)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func testCases(t *testing.T, tc Case, admin *users.User) {
	require.NotNil(t, admin)
	assert.Equal(t, tc.Id(), admin.Id())
	assert.Equal(t, tc.FirstName(), admin.FirstName())
	assert.Equal(t, tc.Username(), admin.Username())
	assert.Equal(t, tc.Email(), admin.Email())
	assert.Equal(t, tc.Password(), admin.Password())
	assert.Equal(t, tc.Gender(), admin.Gender())
	assert.Equal(t, tc.Birth(), admin.Birth())
	assert.Equal(t, tc.Country(), admin.Country())
	assert.Equal(t, tc.Language(), admin.Language())
	assert.Equal(t, tc.Phone(), admin.Phone())
	assert.Equal(t, tc.Information(), admin.Information())
	assert.Equal(t, tc.ProfilePic(), admin.ProfilePic())
	assert.Equal(t, tc.WebSite(), admin.WebSite())
	assert.Equal(t, tc.Visibility(), admin.Visibility())
	assert.Equal(t, tc.LastLoginAt(), admin.LastLoginAt())
	assert.Equal(t, tc.UpdatedAt(), admin.UpdatedAt())
	assert.Equal(t, tc.DeletedAt(), admin.DeletedAt())
}

func userCases() []Case {
	var cases []Case
	for v, c := range listUsers() {
		newCase := Case{"Case " + string(rune(v)), c}
		cases = append(cases, newCase)
	}
	return cases
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
		usr.Update()

		if c.deletedAt != nil {
			time.Sleep(5 * time.Millisecond)
			_ = usr.Delete()
		}

		usersList = append(usersList, usr)
	}

	return usersList
}
