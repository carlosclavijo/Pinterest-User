package controllers

import (
	"database/sql"
	"encoding/json"
	"github.com/carlosclavijo/Pinterest-Services/internal/application/user/commands"
	"github.com/carlosclavijo/Pinterest-Services/internal/application/user/dto"
	command "github.com/carlosclavijo/Pinterest-Services/internal/application/user/handlers"
	"github.com/carlosclavijo/Pinterest-Services/internal/application/user/queries"
	"github.com/carlosclavijo/Pinterest-Services/internal/domain/shared"
	"github.com/carlosclavijo/Pinterest-Services/internal/domain/user"
	query "github.com/carlosclavijo/Pinterest-Services/internal/infrastructure/handlers/users"
	"github.com/carlosclavijo/Pinterest-Services/internal/infrastructure/persistence/repositories"
	"github.com/carlosclavijo/Pinterest-Services/internal/infrastructure/services"
	"github.com/carlosclavijo/Pinterest-Services/internal/web/helpers"
	"github.com/carlosclavijo/Pinterest-Services/internal/web/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"go/types"
	"net/http"
	"strings"
	"time"
)

type UserController struct {
	commandHandler *command.UserHandler
	queryHandler   *query.UserHandler
	jwtService     *services.JWTService
	blacklistRepo  *services.TokenBlacklist
}

func NewUserController(db *sql.DB, jwt *services.JWTService, blacklistRepo *services.TokenBlacklist, emService *services.EmailService) *UserController {
	repository := repositories.NewUserRepository(db)
	factory := users.NewUserFactory()
	emailRepo := repositories.NewEmailVerificationRepo(db)
	commandHandler := command.NewUserHandler(repository, emailRepo, emService, factory)
	queryHandler := query.NewUserHandler(repository, factory)
	return &UserController{
		commandHandler: commandHandler,
		queryHandler:   queryHandler,
		jwtService:     jwt,
		blacklistRepo:  blacklistRepo,
	}
}

const (
	ErrFetchUsers = "Could not fetch users"
	ErrParseId    = "Invalid ID format"
	ErrJSONFormat = "Invalid JSON format or fields"
)

// GetAllUsers godoc
// @Summary      Get all users
// @Description  Returns a list of all registered users
// @Tags         users
// @Produce      json
// @Success      200  {object}  helpers.GetListUsersDTO
// @Failure      500  {object}  helpers.GetListUsersDTO "Server error"
// @Router       /users/all [get]
func (c *UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	qry := queries.GetAllUsersQuery{}
	usersList, err := c.queryHandler.HandleGetAll(r.Context(), qry)
	if err != nil {
		errStr := err.Error()
		helpers.WriteJSON(w, http.StatusInternalServerError, helpers.Response[types.Nil]{
			Success: false,
			Error: &helpers.Error{
				Code:    "GET_ALL_FAILED",
				Message: ErrFetchUsers,
				Err:     &errStr,
			},
		})
		return
	}

	length := len(usersList)
	helpers.WriteJSON(w, http.StatusOK, helpers.Response[[]*dto.UserDTO]{
		Success: true,
		Data:    usersList,
		Length:  &length,
	})
}

// GetListUsers godoc
// @Summary Get list of active users
// @Description Returns all users where deleted_at IS NULL
// @Tags Users
// @Produce json
// @Success 200 {object} helpers.GetListUsersDTO
// @Failure 500 {object} helpers.GetListUsersDTO "Server error"
// @Router /users/list [get]
func (c *UserController) GetListUsers(w http.ResponseWriter, r *http.Request) {
	qry := queries.GetListUsersQuery{}
	usersList, err := c.queryHandler.HandleGetList(r.Context(), qry)
	if err != nil {
		errStr := err.Error()
		helpers.WriteJSON(w, http.StatusInternalServerError, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "GET_LIST_FAILED",
				Message: ErrFetchUsers,
				Err:     &errStr,
			},
		})
		return
	}

	length := len(usersList)
	helpers.WriteJSON(w, http.StatusOK, helpers.Response[[]*dto.UserDTO]{
		Success: true,
		Data:    usersList,
		Length:  &length,
	})
}

// GetUserById godoc
// @Summary      Get user by ID
// @Description  Returns a single user by UUID
// @Tags         users
// @Produce      json
// @Param        id   path      string  true  "User ID (UUID)"
// @Success      200  {object}  helpers.GetUserDTO
// @Failure      400  {object}  helpers.GetUserDTO  "Invalid id"
// @Failure      500  {object}  helpers.GetUserDTO  "Server error"
// @Router       /users/id/{id} [get]
func (c *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		errStr := err.Error()
		helpers.WriteJSON(w, http.StatusBadRequest, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "PARSING_UUID_FAILED",
				Message: ErrParseId,
				Err:     &errStr,
			},
		})
		return
	}

	qry := queries.GetUserByIdQuery{
		Id: id,
	}

	usr, err := c.queryHandler.HandleGetById(r.Context(), qry)
	if err != nil {
		errStr := err.Error()
		helpers.WriteJSON(w, http.StatusInternalServerError, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "GET_BY_ID_FAILED",
				Message: ErrFetchUsers,
				Err:     &errStr,
			},
		})
		return
	}

	helpers.WriteJSON(w, http.StatusOK, helpers.Response[*dto.UserDTO]{
		Success: true,
		Data:    usr,
	})
}

// GetUserByEmail godoc
// @Summary      Get user by email
// @Description  Returns a single user by email address
// @Tags         users
// @Produce      json
// @Param        email  path      string  true  "User email"
// @Success      200    {object}  helpers.GetUserDTO
// @Failure      400    {object}  helpers.GetUserDTO  "Invalid email"
// @Failure      500    {object}  helpers.GetUserDTO  "Server error"
// @Router       /users/email/{email} [get]
func (c *UserController) GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")

	qry := queries.GetUserByEmailQuery{
		Email: email,
	}

	usr, err := c.queryHandler.HandleGetByEmail(r.Context(), qry)
	if err != nil {
		errStr := err.Error()
		helpers.WriteJSON(w, http.StatusInternalServerError, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "GET_BY_EMAIL_FAILED",
				Message: ErrFetchUsers,
				Err:     &errStr,
			},
		})
		return
	}

	helpers.WriteJSON(w, http.StatusOK, helpers.Response[*dto.UserDTO]{
		Success: true,
		Data:    usr,
	})
}

// GetUserByUsername godoc
// @Summary      Get user by username
// @Description  Returns a single user by username
// @Tags         users
// @Produce      json
// @Param        username  path      string  true  "Username"
// @Success      200       {object}  helpers.GetUserDTO
// @Failure      400       {object}  helpers.GetUserDTO  "Invalid username"
// @Failure      500       {object}  helpers.GetUserDTO  "Server error"
// @Router       /users/username/{username} [get]
func (c *UserController) GetUserByUsername(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")

	qry := queries.GetUserByUsernameQuery{
		Username: username,
	}

	usr, err := c.queryHandler.HandleGetByUsername(r.Context(), qry)
	if err != nil {
		errStr := err.Error()
		helpers.WriteJSON(w, http.StatusInternalServerError, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "GET_BY_USERNAME_FAILED",
				Message: ErrFetchUsers,
				Err:     &errStr,
			},
		})
		return
	}

	helpers.WriteJSON(w, http.StatusOK, helpers.Response[*dto.UserDTO]{
		Success: true,
		Data:    usr,
	})
}

// GetUsersByCountry godoc
// @Summary      Get users by country
// @Description  Returns a list of users filtered by country
// @Tags         users
// @Produce      json
// @Param        country  path      string  true  "Country name"
// @Success      200      {object}  helpers.GetListUsersDTO
// @Failure      400      {object}  helpers.GetListUsersDTO  "Invalid country"
// @Failure      500      {object}  helpers.GetListUsersDTO  "Server error"
// @Router       /users/country/{country} [get]
func (c *UserController) GetUsersByCountry(w http.ResponseWriter, r *http.Request) {
	country := chi.URLParam(r, "country")

	qry := queries.GetUsersByCountryQuery{
		Country: country,
	}

	usersList, err := c.queryHandler.HandleGetListByCountry(r.Context(), qry)
	if err != nil {
		errStr := err.Error()
		helpers.WriteJSON(w, http.StatusInternalServerError, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "GET_BY_COUNTRY_FAILED",
				Message: ErrFetchUsers,
				Err:     &errStr,
			},
		})
		return
	}

	length := len(usersList)
	helpers.WriteJSON(w, http.StatusOK, helpers.Response[[]*dto.UserDTO]{
		Success: true,
		Data:    usersList,
		Length:  &length,
	})
}

// GetUsersByLanguage godoc
// @Summary      Get users by language
// @Description  Returns a list of users filtered by language
// @Tags         users
// @Produce      json
// @Param        country  path      string  true  "Country name"
// @Success      200      {object}  helpers.GetListUsersDTO
// @Failure      400      {object}  helpers.GetListUsersDTO  "Invalid country"
// @Failure      500      {object}  helpers.GetListUsersDTO  "Server error"
// @Router       /users/language/{language} [get]
func (c *UserController) GetUsersByLanguage(w http.ResponseWriter, r *http.Request) {
	language := chi.URLParam(r, "language")

	qry := queries.GetUsersByLanguageQuery{
		Language: language,
	}

	usersList, err := c.queryHandler.HandleGetListByLanguage(r.Context(), qry)
	if err != nil {
		errStr := err.Error()
		helpers.WriteJSON(w, http.StatusInternalServerError, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "GET_BY_LANGUAGE_FAILED",
				Message: ErrFetchUsers,
				Err:     &errStr,
			},
		})
		return
	}

	length := len(usersList)
	helpers.WriteJSON(w, http.StatusOK, helpers.Response[[]*dto.UserDTO]{
		Success: true,
		Data:    usersList,
		Length:  &length,
	})
}

// GetUsersLikeUsername godoc
// @Summary      Get users by partial username
// @Description  Returns a list of users whose usernames match the provided pattern
// @Tags         users
// @Produce      json
// @Param        username  path      string  true  "Username pattern"
// @Success      200       {object}  helpers.GetListUsersDTO
// @Failure      400       {object}  helpers.GetListUsersDTO  "Invalid username"
// @Failure      500       {object}  helpers.GetListUsersDTO  "Server error"
// @Router       /users/like-username/{username} [get]
func (c *UserController) GetUsersLikeUsername(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")

	qry := queries.GetUsersLikeUsernameQuery{
		Username: username,
	}

	usersList, err := c.queryHandler.HandleGetListLikeUsername(r.Context(), qry)
	if err != nil {
		errStr := err.Error()
		helpers.WriteJSON(w, http.StatusInternalServerError, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "GET_LIKE_USERNAME_FAILED",
				Message: ErrFetchUsers,
				Err:     &errStr,
			},
		})
		return
	}

	length := len(usersList)
	helpers.WriteJSON(w, http.StatusOK, helpers.Response[[]*dto.UserDTO]{
		Success: true,
		Data:    usersList,
		Length:  &length,
	})
}

// CreateUser godoc
// @Summary      Create a new user
// @Description  Creates a new user in the system
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      commands.CreateUserCommand  true  "User creation payload"
// @Success      201   {object}  helpers.GetUserResponse
// @Failure      400   {object}  helpers.GetUserResponse  "Invalid request body"
// @Failure      500   {object}  helpers.GetUserResponse  "Server error"
// @Router       /users/create [post]
func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var cmd commands.CreateUserCommand
	if err := json.NewDecoder(r.Body).Decode(&cmd); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "INVALID_REQUEST_BODY",
				Message: ErrJSONFormat,
			},
		})
		return
	}

	usr, err := c.commandHandler.HandleCreate(r.Context(), cmd)
	if err != nil {
		errStr := err.Error()
		helpers.WriteJSON(w, http.StatusInternalServerError, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "CREATE_FAILED",
				Message: "Could not create user",
				Err:     &errStr,
			},
		})
		return
	}

	helpers.WriteJSON(w, http.StatusCreated, helpers.Response[*dto.UserResponse]{
		Success: true,
		Data:    usr,
	})
}

// LoginUser godoc
// @Summary      Login a user
// @Description  Authenticates a user and returns a JWT token
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        credentials  body      commands.LoginUserCommand  true  "User login payload"
// @Success      200          {object}  helpers.LoginSuccessResponse  "User data and token"
// @Failure      400          {object}  helpers.GetUserResponse  "Invalid request body"
// @Failure      401          {object}  helpers.GetUserResponse  "Authentication failed"
// @Failure      500          {object}  helpers.GetUserResponse  "Server error"
// @Router       /users/login [post]
func (c *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {
	var cmd commands.LoginUserCommand
	if err := json.NewDecoder(r.Body).Decode(&cmd); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "INVALID_REQUEST_BODY",
				Message: "Invalid JSON",
			},
		})
		return
	}

	usr, err := c.commandHandler.HandleLogin(r.Context(), cmd)
	if err != nil {
		errStr := err.Error()
		helpers.WriteJSON(w, http.StatusUnauthorized, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "LOGIN_FAILED",
				Message: "Could not login user",
				Err:     &errStr,
			},
		})
		return
	}

	token, err := c.jwtService.Generate(usr.Id.String())
	if err != nil {
		errStr := err.Error()
		helpers.WriteJSON(w, http.StatusInternalServerError, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "TOKEN_ERROR",
				Message: "Could not generate token",
				Err:     &errStr,
			},
		})
		return
	}

	helpers.WriteJSON(w, http.StatusOK, helpers.Response[any]{
		Success: true,
		Data: map[string]any{
			"user":  usr,
			"token": token,
		},
	})
}

// Logout godoc
// @Summary      Logout user
// @Description  Revokes the current JWT token
// @Tags         users
// @Produce      json
// @Param        Authorization  header  string  true  "Bearer token"
// @Success      200  {object}  helpers.LogoutSuccessResponse  "Logout successful"
// @Failure      401  {object}  helpers.GetUserResponse  "Missing or invalid token"
// @Failure      500  {object}  helpers.GetUserResponse  "Logout failed"
// @Router       /users/logout [post]
func (c *UserController) Logout(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		helpers.WriteJSON(w, http.StatusUnauthorized, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "MISS_AUTH_HEADER",
				Message: "missing authorization header",
			},
		})
		return
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		helpers.WriteJSON(w, http.StatusUnauthorized, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "INVALID_AUTH_HEADER",
				Message: "invalid authorization header",
			},
		})
		return
	}
	tokenStr := parts[1]

	claims, err := c.jwtService.ParseToken(tokenStr)
	if err != nil {
		helpers.WriteJSON(w, http.StatusUnauthorized, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "INVALID_TOKEN",
				Message: "invalid token",
			},
		})
		return
	}

	exp := time.Unix(int64(claims["exp"].(float64)), 0)

	if err = c.blacklistRepo.Add(tokenStr, exp); err != nil {
		helpers.WriteJSON(w, http.StatusInternalServerError, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "LOG_FAILED",
				Message: "logout failed",
			},
		})
		return
	}

	helpers.WriteJSON(w, http.StatusOK, helpers.Response[string]{
		Success: true,
		Data:    "Successfully log out",
	})
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Logged out successfully"))
}

// UpdateUser godoc
// @Summary      Update a user
// @Description  Updates the authenticated user's information
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      commands.UpdateUserCommand  true  "User update payload"
// @Success      200   {object}  helpers.GetUserResponse
// @Failure      400   {object}  helpers.GetUserResponse  "Invalid request body"
// @Failure      403   {object}  helpers.GetUserResponse  "Forbidden: cannot update another user"
// @Failure      500   {object}  helpers.GetUserResponse  "Server error"
// @Router       /users/ [put]
func (c *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var cmd commands.UpdateUserCommand
	if err := json.NewDecoder(r.Body).Decode(&cmd); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "INVALID_REQUEST_BODY",
				Message: ErrJSONFormat,
			},
		})
		return
	}

	authUserID := r.Context().Value("user_id").(string)

	if cmd.Id.String() != authUserID {
		helpers.WriteJSON(w, http.StatusForbidden, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "FORBIDDEN",
				Message: "Cannot update another user",
			},
		})
		return
	}

	usr, err := c.commandHandler.HandleUpdate(r.Context(), cmd)
	if err != nil {
		errStr := err.Error()
		helpers.WriteJSON(w, http.StatusInternalServerError, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "UPDATE_FAILED",
				Message: "Could not update user",
				Err:     &errStr,
			},
		})
		return
	}

	helpers.WriteJSON(w, http.StatusOK, helpers.Response[*dto.UserResponse]{
		Success: true,
		Data:    usr,
	})
}

// UploadProfilePic godoc
// @Summary      Upload a user's profile picture
// @Description  Uploads a profile picture for the given user
// @Tags         users
// @Accept       multipart/form-data
// @Produce      json
// @Param        id           path      string  true  "User ID"
// @Param        profile_pic  formData  file    true  "Profile picture file"
// @Success      200          {object}  helpers.LogoutSuccessResponse  "Uploaded file info"
// @Failure      400          {object}  helpers.GetUserDTO          "Bad request / missing file"
// @Failure      500          {object}  helpers.GetUserDTO          "Server error"
// @Router       /users/profilepic/{id} [patch]
func (c *UserController) UploadProfilePic(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "id")
	if userID == "" {
		helpers.WriteJSON(w, http.StatusBadRequest, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "INVALID_ID",
				Message: "User ID is required in the URL",
			},
		})
		return
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "INVALID_MULTIPART",
				Message: "Invalid multipart form data",
			},
		})
		return
	}

	file, handler, err := r.FormFile("profile_pic")
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "NO_PROFILE_PIC",
				Message: "Profile picture file is required",
			},
		})
		return
	}
	defer file.Close()

	fileService := services.NewFileService("../../")
	fileName, err := fileService.SaveProfilePic(file, handler.Filename)
	if err != nil {
		newErr := err.Error()
		helpers.WriteJSON(w, http.StatusInternalServerError, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "SAVE_FAILED",
				Message: "Failed to save file",
				Err:     &newErr,
			},
		})
		return
	}

	cmd := commands.UpdateProfilePicCommand{
		UserID:     userID,
		ProfilePic: fileName,
	}

	if err = c.commandHandler.HandleUpdateProfilePic(r.Context(), cmd); err != nil {
		helpers.WriteJSON(w, http.StatusInternalServerError, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "SAVE_FAILED",
				Message: "Failed to update user profile picture",
			},
		})
		return
	}

	helpers.WriteJSON(w, http.StatusOK, helpers.Response[string]{
		Success: true,
		Data:    "Successfully update profile pic",
	})
}

// DeleteUser godoc
// @Summary      Delete a user
// @Description  Deletes a user by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "User ID"
// @Success      200  {object}  helpers.GetUserResponse  "Deleted user"
// @Failure      400  {object}  helpers.GetUserResponse  "Invalid UUID"
// @Failure      500  {object}  helpers.GetUserResponse  "Server error"
// @Router       /users/{id} [delete]
func (c *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		errStr := err.Error()
		helpers.WriteJSON(w, http.StatusBadRequest, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "PARSING_UUID_FAILED",
				Message: ErrParseId,
				Err:     &errStr,
			},
		})
		return
	}

	usr, err := c.commandHandler.HandleDelete(r.Context(), id)
	if err != nil {
		errStr := err.Error()
		helpers.WriteJSON(w, http.StatusInternalServerError, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "DELETE_FAILED",
				Message: ErrFetchUsers,
				Err:     &errStr,
			},
		})
		return
	}

	helpers.WriteJSON(w, http.StatusOK, helpers.Response[*dto.UserResponse]{
		Success: true,
		Data:    usr,
	})
}

// RestoreUser godoc
// @Summary      Restore a deleted user
// @Description  Restores a user by ID that was previously soft-deleted
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "User ID"
// @Success      200  {object}  helpers.GetUserResponse  "Restored user"
// @Failure      400  {object}  helpers.GetUserResponse  "Invalid UUID"
// @Failure      500  {object}  helpers.GetUserResponse  "Server error"
// @Router       /users/restore/{id} [patch]
func (c *UserController) RestoreUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		errStr := err.Error()
		helpers.WriteJSON(w, http.StatusBadRequest, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "PARSING_UUID_FAILED",
				Message: ErrParseId,
				Err:     &errStr,
			},
		})
		return
	}

	usr, err := c.commandHandler.HandleRestore(r.Context(), id)
	if err != nil {
		errStr := err.Error()
		helpers.WriteJSON(w, http.StatusInternalServerError, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "RESTORE_FAILED",
				Message: ErrFetchUsers,
				Err:     &errStr,
			},
		})
		return
	}

	helpers.WriteJSON(w, http.StatusOK, helpers.Response[*dto.UserResponse]{
		Success: true,
		Data:    usr,
	})
}

// ListCountries godoc
// @Summary      List all countries
// @Description  Returns a list of available countries
// @Tags         users
// @Produce      json
// @Success      200  {object}  helpers.GetCountriesList  "List of countries"
// @Router       /users/countries [get]
func (c *UserController) ListCountries(w http.ResponseWriter, r *http.Request) {
	helpers.WriteJSON(w, http.StatusOK, helpers.Response[[]shared.Country]{
		Success: true,
		Data:    shared.ListCountries(),
	})
}

// ListLanguages godoc
// @Summary      List all languages
// @Description  Returns a list of available languages
// @Tags         users
// @Produce      json
// @Success      200  {object}  helpers.GetLanguagesList  "List of languages"
// @Router       /users/languages [get]
func (c *UserController) ListLanguages(w http.ResponseWriter, r *http.Request) {
	helpers.WriteJSON(w, http.StatusOK, helpers.Response[[]shared.Language]{
		Success: true,
		Data:    shared.ListLanguages(),
	})
}

// VerifyEmail godoc
// @Summary      Verify user's email
// @Description  Verifies the email of a user using a token
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        token  query  string  true  "Verification token"
// @Success      303  "Redirects to login page with verification success"
// @Failure      400  {string}  string  "Missing or invalid token"
// @Router       /verify-email [get]
func (c *UserController) VerifyEmail(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	if token == "" {
		http.Error(w, "missing token", http.StatusBadRequest)
		return
	}

	err := c.commandHandler.VerifyEmailToken(r.Context(), token)
	if err != nil {
		http.Error(w, "invalid or expired token", http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "http://localhost:3000/login?verified=true", http.StatusSeeOther)
}

func (c *UserController) RegisterRoutes(r chi.Router) {
	r.Post("/create", c.CreateUser)
	r.Post("/login", c.LoginUser)

	r.Group(func(r chi.Router) {
		r.Use(middleware.JWTMiddleware(c.jwtService, c.blacklistRepo))

		r.Get("/all", c.GetAllUsers)
		r.Get("/list", c.GetListUsers)
		r.Get("/countries", c.ListCountries)
		r.Get("/languages", c.ListLanguages)
		r.Get("/id/{id}", c.GetUserById)
		r.Get("/email/{email}", c.GetUserByEmail)
		r.Get("/username/{username}", c.GetUserByUsername)
		r.Get("/like-username/{username}", c.GetUsersLikeUsername)
		r.Get("/country/{country}", c.GetUsersByCountry)
		r.Get("/language/{language}", c.GetUsersByLanguage)
		r.Put("/", c.UpdateUser)
		r.Delete("/{id}", c.DeleteUser)
		r.Patch("/profilepic/{id}", c.UploadProfilePic)
		r.Patch("/restore/{id}", c.RestoreUser)
		r.Post("/out", c.Logout)
	})
}
