package controllers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/carlosclavijo/Pinterest-Services/internal/application/user/commands"
	"github.com/carlosclavijo/Pinterest-Services/internal/application/user/dto"
	"github.com/carlosclavijo/Pinterest-Services/internal/infrastructure/persistence/repositories"
	"github.com/carlosclavijo/Pinterest-Services/internal/infrastructure/services"
	"github.com/carlosclavijo/Pinterest-Services/internal/web/helpers"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"io"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"
	"time"
)

var (
	password = "sTr0nG!!"
	columns  = []string{
		"id", "first_name", "last_name", "user_name", "email", "password", "gender", "birth_date", "country", "language", "phone", "information", "profile_pic", "web_site",
		"visibility", "last_login_at", "created_at", "updated_at", "deleted_at",
	}
)

func TestUserController_GetAllUsers(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	ctrl := NewUserController(db, services.JWTService{})

	userDto := mockUserDto()
	rows := sqlmock.NewRows(columns).AddRow(
		userDto.Id, userDto.FirstName, userDto.LastName, userDto.Username, userDto.Email, password, userDto.Gender, userDto.Birth, userDto.Country, userDto.Language,
		*userDto.Phone, *userDto.Information, *userDto.ProfilePic, *userDto.Website, userDto.Visibility, time.Now(), time.Now(), time.Now(), nil,
	)
	mock.ExpectQuery(regexp.QuoteMeta(repositories.QueryGetAllUsers)).WillReturnRows(rows)

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rr := httptest.NewRecorder()

	ctrl.GetAllUsers(rr, req)

	resp := rr.Result()
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	require.Equal(t, http.StatusOK, resp.StatusCode)
	require.Contains(t, string(body), `"length":1`)

	var usr helpers.Response[[]*dto.UserDTO]

	require.NoError(t, json.Unmarshal(body, &usr))
	require.Len(t, usr.Data, 1)

	userResponse := usr.Data[0]

	listUserCases(t, body, userDto, userResponse)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserController_GetAllUsers_Error(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	ctrl := NewUserController(db, services.JWTService{})

	mock.ExpectQuery(regexp.QuoteMeta(repositories.QueryGetAllUsers)).WillReturnError(errors.New("DB connection failed"))

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rr := httptest.NewRecorder()

	ctrl.GetAllUsers(rr, req)

	resp := rr.Result()
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	assert.Contains(t, string(body), `"success":false`)
	assert.Contains(t, string(body), `"code":"GET_ALL_FAILED"`)
	assert.Contains(t, string(body), `"message":"Could not fetch users"`)
	assert.Contains(t, string(body), "DB connection failed")

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserController_GetListUsers(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	ctrl := NewUserController(db, services.JWTService{})

	userDto := mockUserDto()
	rows := sqlmock.NewRows(columns).AddRow(
		userDto.Id, userDto.FirstName, userDto.LastName, userDto.Username, userDto.Email, password, userDto.Gender, userDto.Birth, userDto.Country, userDto.Language,
		*userDto.Phone, *userDto.Information, *userDto.ProfilePic, *userDto.Website, userDto.Visibility, time.Now(), time.Now(), time.Now(), nil,
	)
	mock.ExpectQuery(regexp.QuoteMeta(repositories.QueryGetListUsers)).WillReturnRows(rows)

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rr := httptest.NewRecorder()

	ctrl.GetListUsers(rr, req)

	resp := rr.Result()
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	var usr helpers.Response[[]*dto.UserDTO]

	require.NoError(t, json.Unmarshal(body, &usr))
	require.Len(t, usr.Data, 1)

	userResponse := usr.Data[0]

	listUserCases(t, body, userDto, userResponse)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserController_GetListUsers_Error(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	ctrl := NewUserController(db, services.JWTService{})

	mock.ExpectQuery(regexp.QuoteMeta(repositories.QueryGetListUsers)).WillReturnError(errors.New("DB connection failed"))

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rr := httptest.NewRecorder()

	ctrl.GetListUsers(rr, req)

	resp := rr.Result()
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	assert.Contains(t, string(body), `"success":false`)
	assert.Contains(t, string(body), `"code":"GET_LIST_FAILED"`)
	assert.Contains(t, string(body), `"message":"Could not fetch users"`)
	assert.Contains(t, string(body), "DB connection failed")

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserController_GetUserById(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	ctrl := NewUserController(db, services.JWTService{})

	cols := append([]string(nil), columns...)
	cols = cols[1:]
	userDto := mockUserDto()
	rows := sqlmock.NewRows(cols).AddRow(
		userDto.FirstName, userDto.LastName, userDto.Username, userDto.Email, password, userDto.Gender, userDto.Birth, userDto.Country, userDto.Language,
		*userDto.Phone, *userDto.Information, *userDto.ProfilePic, *userDto.Website, userDto.Visibility, time.Now(), time.Now(), time.Now(), nil,
	)
	mock.ExpectQuery(regexp.QuoteMeta(repositories.QueryGetUserById)).WithArgs(userDto.Id).WillReturnRows(rows)

	req := httptest.NewRequest(http.MethodGet, "/users/"+userDto.Id.String(), nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", userDto.Id.String())
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	rr := httptest.NewRecorder()
	ctrl.GetUserById(rr, req)

	resp := rr.Result()
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	var usr helpers.Response[*dto.UserDTO]

	require.NoError(t, json.Unmarshal(body, &usr))

	userResponse := usr.Data

	listUserCases(t, body, userDto, userResponse)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserController_GetUserById_InvalidUUID(t *testing.T) {
	db, _, _ := sqlmock.New()
	defer db.Close()

	ctrl := NewUserController(db, services.JWTService{})

	req := httptest.NewRequest(http.MethodGet, "/users/invalid-uuid", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "invalid-uuid")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	rr := httptest.NewRecorder()
	ctrl.GetUserById(rr, req)

	resp := rr.Result()
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	assert.Contains(t, string(body), `"success":false`)
	assert.Contains(t, string(body), `"code":"PARSING_UUID_FAILED"`)
	assert.Contains(t, string(body), `"message":"Invalid user ID format"`)
}

func TestUserController_GetUserById_Error(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	ctrl := NewUserController(db, services.JWTService{})

	userDto := mockUserDto()
	mock.ExpectQuery(regexp.QuoteMeta(repositories.QueryGetUserById)).WithArgs(userDto.Id).WillReturnError(errors.New("DB connection failed"))

	req := httptest.NewRequest(http.MethodGet, "/users/"+userDto.Id.String(), nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", userDto.Id.String())
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	rr := httptest.NewRecorder()
	ctrl.GetUserById(rr, req)

	resp := rr.Result()
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	assert.Contains(t, string(body), `"success":false`)
	assert.Contains(t, string(body), `"code":"GET_BY_ID_FAILED"`)
	assert.Contains(t, string(body), `"message":"Could not fetch users"`)
	assert.Contains(t, string(body), "DB connection failed")

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserController_GetUserByUsername(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	ctrl := NewUserController(db, services.JWTService{})

	cols := append([]string(nil), columns...)
	cols = append(cols[:3], cols[4:]...)
	userDto := mockUserDto()
	rows := sqlmock.NewRows(cols).AddRow(
		userDto.Id, userDto.FirstName, userDto.LastName, userDto.Email, password, userDto.Gender, userDto.Birth, userDto.Country, userDto.Language,
		*userDto.Phone, *userDto.Information, *userDto.ProfilePic, *userDto.Website, userDto.Visibility, time.Now(), time.Now(), time.Now(), nil,
	)
	mock.ExpectQuery(regexp.QuoteMeta(repositories.QueryGetUserByUsername)).WithArgs(userDto.Username).WillReturnRows(rows)

	req := httptest.NewRequest(http.MethodGet, "/users/"+userDto.Username, nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("username", userDto.Username)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	rr := httptest.NewRecorder()
	ctrl.GetUserByUsername(rr, req)

	resp := rr.Result()
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	var usr helpers.Response[*dto.UserDTO]

	require.NoError(t, json.Unmarshal(body, &usr))

	userResponse := usr.Data

	listUserCases(t, body, userDto, userResponse)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserController_GetUserByUsername_Error(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	ctrl := NewUserController(db, services.JWTService{})

	userDto := mockUserDto()
	mock.ExpectQuery(regexp.QuoteMeta(repositories.QueryGetUserByUsername)).WithArgs(userDto.Username).WillReturnError(errors.New("DB connection failed"))

	req := httptest.NewRequest(http.MethodGet, "/users/"+userDto.Username, nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("username", userDto.Username)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	rr := httptest.NewRecorder()
	ctrl.GetUserByUsername(rr, req)

	resp := rr.Result()
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	assert.Contains(t, string(body), `"success":false`)
	assert.Contains(t, string(body), `"code":"GET_BY_USERNAME_FAILED"`)
	assert.Contains(t, string(body), `"message":"Could not fetch users"`)
	assert.Contains(t, string(body), "DB connection failed")

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserController_GetUserByEmail(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	ctrl := NewUserController(db, services.JWTService{})

	cols := append([]string(nil), columns...)
	cols = append(cols[:4], cols[5:]...)
	userDto := mockUserDto()
	rows := sqlmock.NewRows(cols).AddRow(
		userDto.Id, userDto.FirstName, userDto.LastName, userDto.Username, password, userDto.Gender, userDto.Birth, userDto.Country, userDto.Language,
		*userDto.Phone, *userDto.Information, *userDto.ProfilePic, *userDto.Website, userDto.Visibility, time.Now(), time.Now(), time.Now(), nil,
	)
	mock.ExpectQuery(regexp.QuoteMeta(repositories.QueryGetUserByEmail)).WithArgs(userDto.Email).WillReturnRows(rows)

	req := httptest.NewRequest(http.MethodGet, "/users/"+userDto.Email, nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("email", userDto.Email)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	rr := httptest.NewRecorder()
	ctrl.GetUserByEmail(rr, req)

	resp := rr.Result()
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	var usr helpers.Response[*dto.UserDTO]

	require.NoError(t, json.Unmarshal(body, &usr))

	userResponse := usr.Data

	listUserCases(t, body, userDto, userResponse)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserController_GetUserByEmail_Error(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	ctrl := NewUserController(db, services.JWTService{})

	userDto := mockUserDto()
	mock.ExpectQuery(regexp.QuoteMeta(repositories.QueryGetUserByEmail)).WithArgs(userDto.Email).WillReturnError(errors.New("DB connection failed"))

	req := httptest.NewRequest(http.MethodGet, "/users/"+userDto.Email, nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("email", userDto.Email)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	rr := httptest.NewRecorder()
	ctrl.GetUserByEmail(rr, req)

	resp := rr.Result()
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	assert.Contains(t, string(body), `"success":false`)
	assert.Contains(t, string(body), `"code":"GET_BY_EMAIL_FAILED"`)
	assert.Contains(t, string(body), `"message":"Could not fetch users"`)
	assert.Contains(t, string(body), "DB connection failed")

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserController_GetUsersByCountry(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	ctrl := NewUserController(db, services.JWTService{})

	cols := append([]string(nil), columns...)
	cols = append(cols[:9], cols[10:]...)
	userDto := mockUserDto()
	rows := sqlmock.NewRows(cols).AddRow(
		userDto.Id, userDto.FirstName, userDto.LastName, userDto.Username, userDto.Email, password, userDto.Gender, userDto.Birth, userDto.Language,
		*userDto.Phone, *userDto.Information, *userDto.ProfilePic, *userDto.Website, userDto.Visibility, time.Now(), time.Now(), time.Now(), nil,
	)
	mock.ExpectQuery(regexp.QuoteMeta(repositories.QueryGetUsersByCountry)).WithArgs(userDto.Country).WillReturnRows(rows)

	req := httptest.NewRequest(http.MethodGet, "/users/"+userDto.Country, nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("country", userDto.Country)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	rr := httptest.NewRecorder()
	ctrl.GetUsersByCountry(rr, req)

	resp := rr.Result()
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	var usr helpers.Response[[]*dto.UserDTO]

	require.NoError(t, json.Unmarshal(body, &usr))

	userResponse := usr.Data[0]

	listUserCases(t, body, userDto, userResponse)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserController_GetUsersByCountry_Error(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	ctrl := NewUserController(db, services.JWTService{})

	userDto := mockUserDto()
	mock.ExpectQuery(regexp.QuoteMeta(repositories.QueryGetUsersByCountry)).WithArgs(userDto.Country).WillReturnError(errors.New("DB connection failed"))

	req := httptest.NewRequest(http.MethodGet, "/users/"+userDto.Country, nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("country", userDto.Country)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	rr := httptest.NewRecorder()
	ctrl.GetUsersByCountry(rr, req)

	resp := rr.Result()
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	assert.Contains(t, string(body), `"success":false`)
	assert.Contains(t, string(body), `"code":"GET_BY_COUNTRY_FAILED"`)
	assert.Contains(t, string(body), `"message":"Could not fetch users"`)
	assert.Contains(t, string(body), "DB connection failed")

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserController_GetUsersByLanguage(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	ctrl := NewUserController(db, services.JWTService{})

	cols := append([]string(nil), columns...)
	cols = append(cols[:10], cols[11:]...)
	userDto := mockUserDto()
	rows := sqlmock.NewRows(cols).AddRow(
		userDto.Id, userDto.FirstName, userDto.LastName, userDto.Username, userDto.Email, password, userDto.Gender, userDto.Birth, userDto.Country,
		*userDto.Phone, *userDto.Information, *userDto.ProfilePic, *userDto.Website, userDto.Visibility, time.Now(), time.Now(), time.Now(), nil,
	)
	mock.ExpectQuery(regexp.QuoteMeta(repositories.QueryGetUsersByLanguage)).WithArgs(userDto.Language).WillReturnRows(rows)

	req := httptest.NewRequest(http.MethodGet, "/users/"+userDto.Language, nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("language", userDto.Language)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	rr := httptest.NewRecorder()
	ctrl.GetUsersByLanguage(rr, req)

	resp := rr.Result()
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	var usr helpers.Response[[]*dto.UserDTO]

	require.NoError(t, json.Unmarshal(body, &usr))

	userResponse := usr.Data[0]

	listUserCases(t, body, userDto, userResponse)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserController_GetUsersByLanguage_Error(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	ctrl := NewUserController(db, services.JWTService{})

	userDto := mockUserDto()
	mock.ExpectQuery(regexp.QuoteMeta(repositories.QueryGetUsersByLanguage)).WithArgs(userDto.Language).WillReturnError(errors.New("DB connection failed"))

	req := httptest.NewRequest(http.MethodGet, "/users/"+userDto.Language, nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("language", userDto.Language)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	rr := httptest.NewRecorder()
	ctrl.GetUsersByLanguage(rr, req)

	resp := rr.Result()
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	assert.Contains(t, string(body), `"success":false`)
	assert.Contains(t, string(body), `"code":"GET_BY_LANGUAGE_FAILED"`)
	assert.Contains(t, string(body), `"message":"Could not fetch users"`)
	assert.Contains(t, string(body), "DB connection failed")

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserController_CreateUser(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	ctrl := NewUserController(db, services.JWTService{})

	userDto := mockUserDto()

	cmd := commands.CreateUserCommand{
		FirstName: userDto.FirstName,
		LastName:  userDto.LastName,
		Username:  userDto.Username,
		Email:     userDto.Email,
		Password:  password,
		Gender:    userDto.Gender,
		Birth:     userDto.Birth,
		Country:   userDto.Country,
		Language:  userDto.Language,
		Phone:     userDto.Phone,
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	mock.ExpectQuery(regexp.QuoteMeta(repositories.QueryExistUserByUsername)).WithArgs(userDto.Username).WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(false))
	mock.ExpectQuery(regexp.QuoteMeta(repositories.QueryExistUserByEmail)).WithArgs(userDto.Email).WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(false))
	mock.ExpectQuery(regexp.QuoteMeta(repositories.QueryCreateUser)).WithArgs(
		sqlmock.AnyArg(), userDto.FirstName, userDto.LastName, userDto.Username, userDto.Email, sqlmock.AnyArg(),
		userDto.Gender, userDto.Birth, "BO", "ES", userDto.Phone, false, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
	).WillReturnRows(sqlmock.NewRows(columns).AddRow(
		userDto.Id, userDto.FirstName, userDto.LastName, userDto.Username, userDto.Email, hashedPassword, userDto.Gender, userDto.Birth,
		"BO", "ES", userDto.Phone, userDto.Information, userDto.ProfilePic, userDto.Website, userDto.Visibility, time.Now(), time.Now(), time.Now(), nil,
	))

	body, err := json.Marshal(cmd)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(string(body)))
	rr := httptest.NewRecorder()

	ctrl.CreateUser(rr, req)

	resp := rr.Result()
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)

	assert.Equal(t, http.StatusCreated, resp.StatusCode)
	assert.Contains(t, string(bodyBytes), `"success":true`)

	var usr helpers.Response[*dto.UserDTO]
	require.NoError(t, json.Unmarshal(bodyBytes, &usr))

	userResponse := usr.Data

	listUserCases(t, bodyBytes, userDto, userResponse)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserController_CreateUser_InvalidBody(t *testing.T) {
	db, _, _ := sqlmock.New()
	defer db.Close()

	ctrl := NewUserController(db, services.JWTService{})
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader("{invalid_json}"))
	rr := httptest.NewRecorder()

	ctrl.CreateUser(rr, req)

	resp := rr.Result()
	defer resp.Body.Close()
	bodyBytes, _ := io.ReadAll(resp.Body)

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	assert.Contains(t, string(bodyBytes), `"success":false`)
	assert.Contains(t, string(bodyBytes), `"code":"INVALID_REQUEST_BODY"`)
}

func TestUserController_CreateUser_FailCreate(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	ctrl := NewUserController(db, services.JWTService{})

	userDto := mockUserDto()

	cmd := commands.CreateUserCommand{
		FirstName: userDto.FirstName,
		LastName:  userDto.LastName,
		Username:  userDto.Username,
		Email:     userDto.Email,
		Password:  password,
		Gender:    userDto.Gender,
		Birth:     userDto.Birth,
		Country:   userDto.Country,
		Language:  userDto.Language,
		Phone:     userDto.Phone,
	}

	mock.ExpectQuery(regexp.QuoteMeta(repositories.QueryExistUserByUsername)).WithArgs(userDto.Username).WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(false))
	mock.ExpectQuery(regexp.QuoteMeta(repositories.QueryExistUserByEmail)).WithArgs(userDto.Email).WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(false))

	mock.ExpectQuery(regexp.QuoteMeta(repositories.QueryCreateUser)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
			sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
			sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnError(fmt.Errorf("DB insert failed"))

	body, _ := json.Marshal(cmd)
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(string(body)))
	rr := httptest.NewRecorder()

	ctrl.CreateUser(rr, req)

	resp := rr.Result()
	defer resp.Body.Close()
	bodyBytes, _ := io.ReadAll(resp.Body)

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	assert.Contains(t, string(bodyBytes), `"success":false`)
	assert.Contains(t, string(bodyBytes), `"code":"CREATE_FAILED"`)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func listUserCases(t *testing.T, body []byte, userDto, got *dto.UserDTO) {
	require.Contains(t, string(body), `"success":true`)

	assert.Contains(t, string(body), userDto.Id.String())
	assert.Contains(t, string(body), userDto.FirstName)
	assert.Contains(t, string(body), userDto.LastName)
	assert.Contains(t, string(body), userDto.Username)
	assert.Contains(t, string(body), userDto.Email)
	assert.Contains(t, string(body), userDto.Birth.Format("2006-01-02"))
	assert.Contains(t, string(body), userDto.Country)
	assert.Contains(t, string(body), userDto.Language)
	assert.Contains(t, string(body), *userDto.Phone)
	assert.Contains(t, string(body), *userDto.Information)
	assert.Contains(t, string(body), *userDto.ProfilePic)
	assert.Contains(t, string(body), *userDto.Website)
	assert.Contains(t, string(body), `"visibility":true`)

	assert.Equal(t, userDto.Id, got.Id)
	assert.Equal(t, userDto.FirstName, got.FirstName)
	assert.Equal(t, userDto.LastName, got.LastName)
	assert.Equal(t, userDto.Username, got.Username)
	assert.Equal(t, userDto.Email, got.Email)
	assert.Equal(t, userDto.Birth.Format(time.RFC3339), got.Birth.Format(time.RFC3339))
	assert.Equal(t, userDto.Country, got.Country)
	assert.Equal(t, userDto.Language, got.Language)
	assert.Equal(t, *userDto.Phone, *got.Phone)
	assert.Equal(t, *userDto.Information, *got.Information)
	assert.Equal(t, *userDto.ProfilePic, *got.ProfilePic)
	assert.Equal(t, *userDto.Website, *got.Website)
	assert.Equal(t, userDto.Visibility, got.Visibility)
}

func mockUserDto() *dto.UserDTO {
	phone := "+591-77887878"
	information := "a lot of information"
	profilePic := "./public/images/user/user.jpg"
	website := "https://www.github.com/carlosclavijo/"

	userDto := dto.UserDTO{
		Id:          uuid.New(),
		FirstName:   "John",
		LastName:    "Doe",
		Username:    "johndoe",
		Email:       "john@example.com",
		Gender:      "M",
		Birth:       time.Now().AddDate(-20, -6, -19),
		Country:     "Bolivia",
		Language:    "Spanish",
		Phone:       &phone,
		Information: &information,
		ProfilePic:  &profilePic,
		Website:     &website,
		Visibility:  true,
	}

	return &userDto
}
