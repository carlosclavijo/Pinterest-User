package controllers

import (
	"database/sql"
	"encoding/json"
	"github.com/carlosclavijo/Pinterest-User/internal/application/user/commands"
	command "github.com/carlosclavijo/Pinterest-User/internal/application/user/handlers"
	"github.com/carlosclavijo/Pinterest-User/internal/domain/shared"
	"github.com/carlosclavijo/Pinterest-User/internal/domain/user"
	query "github.com/carlosclavijo/Pinterest-User/internal/infrastructure/handlers"
	"github.com/carlosclavijo/Pinterest-User/internal/infrastructure/persistence/repositories"
	"github.com/carlosclavijo/Pinterest-User/internal/infrastructure/services"
	"github.com/carlosclavijo/Pinterest-User/internal/web/helpers"
	"github.com/carlosclavijo/Pinterest-User/internal/web/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"net/http"
)

type UserController struct {
	commandHandler command.UserHandler
	queryHandler   query.UserHandler
	jwtService     services.JWTService
}

func NewUserController(db *sql.DB, jwt services.JWTService) *UserController {
	repository := repositories.NewUserRepository(db)
	factory := users.NewUserFactory()
	commandHandler := command.NewUserHandler(repository, factory)
	queryHandler := query.NewUserHandler(repository, factory)
	return &UserController{
		commandHandler: *commandHandler,
		queryHandler:   *queryHandler,
		jwtService:     jwt,
	}
}

const (
	ErrFetchUsers = "Could not fetch users"
	ErrParseId    = "Invalid user ID format"
	ErrJSONFormat = "Invalid JSON format or fields"
)

func (c *UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	qry := queries.GetAllUsers{}
	usersList, err := c.queryHandler.HandleGetAll(r.Context(), qry)
	if err != nil {
		errStr := err.Error()
		helpers.WriteJSON(w, http.StatusInternalServerError, helpers.Response[any]{
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
	helpers.WriteJSON(w, http.StatusOK, helpers.Response[[]*dto2.UserDTO]{
		Success: true,
		Data:    usersList,
		Length:  &length,
	})
}

func (c *UserController) GetListUsers(w http.ResponseWriter, r *http.Request) {
	qry := queries2.GetListUsers{}
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
	helpers.WriteJSON(w, http.StatusOK, helpers.Response[[]*dto2.UserDTO]{
		Success: true,
		Data:    usersList,
		Length:  &length,
	})
}

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

	qry := queries2.GetUserById{
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

	helpers.WriteJSON(w, http.StatusOK, helpers.Response[*dto2.UserDTO]{
		Success: true,
		Data:    usr,
	})
}

func (c *UserController) GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")

	qry := queries2.GetUserByEmail{
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

	helpers.WriteJSON(w, http.StatusOK, helpers.Response[*dto2.UserDTO]{
		Success: true,
		Data:    usr,
	})
}

func (c *UserController) GetUserByUsername(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")

	qry := queries2.GetUserByUsername{
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

	helpers.WriteJSON(w, http.StatusOK, helpers.Response[*dto2.UserDTO]{
		Success: true,
		Data:    usr,
	})
}

func (c *UserController) GetUsersByCountry(w http.ResponseWriter, r *http.Request) {
	country := chi.URLParam(r, "country")

	qry := queries2.GetUsersByCountry{
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
	helpers.WriteJSON(w, http.StatusOK, helpers.Response[[]*dto2.UserDTO]{
		Success: true,
		Data:    usersList,
		Length:  &length,
	})
}

func (c *UserController) GetUsersByLanguage(w http.ResponseWriter, r *http.Request) {
	language := chi.URLParam(r, "language")

	qry := queries2.GetUsersByLanguage{
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
	helpers.WriteJSON(w, http.StatusOK, helpers.Response[[]*dto2.UserDTO]{
		Success: true,
		Data:    usersList,
		Length:  &length,
	})
}

func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var cmd users2.CreateUserCommand
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

	helpers.WriteJSON(w, http.StatusCreated, helpers.Response[*dto2.UserResponse]{
		Success: true,
		Data:    usr,
	})
}

func (c *UserController) Login(w http.ResponseWriter, r *http.Request) {
	var cmd users2.LoginUserCommand
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

func (c *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var cmd users2.UpdateUserCommand
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

	helpers.WriteJSON(w, http.StatusOK, helpers.Response[*dto2.UserResponse]{
		Success: true,
		Data:    usr,
	})
}

func (c *UserController) UploadProfilePic(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "id")
	if userID == "" {
		helpers.WriteJSON(w, http.StatusBadRequest, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "MISSING_USER_ID",
				Message: "User ID is required in the URL",
			},
		})
		return
	}

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "INVALID_MULTIPART",
				Message: "Invalid form data",
			},
		})
		return
	}

	file, handler, err := r.FormFile("profile_pic")
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "NO_FILE",
				Message: "Profile picture file is required",
			},
		})
		return
	}
	defer file.Close()

	fileService := services.NewFileService("../../")
	path, err := fileService.SaveProfilePic(file, handler.Filename)
	if err != nil {
		helpers.WriteJSON(w, http.StatusInternalServerError, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "UPLOAD_FAILED",
				Message: err.Error(),
			},
		})
		return
	}

	cmd := commands.UpdateProfilePicCommand{
		UserID:     userID,
		ProfilePic: path,
	}

	err = c.commandHandler.HandleUpdateProfilePic(r.Context(), cmd)
	if err != nil {
		helpers.WriteJSON(w, http.StatusInternalServerError, helpers.Response[any]{
			Success: false,
			Error: &helpers.Error{
				Code:    "UPDATE_FAILED",
				Message: err.Error(),
			},
		})
		return
	}

	helpers.WriteJSON(w, http.StatusOK, helpers.Response[any]{
		Success: true,
		Data: map[string]string{
			"profile_pic": path,
		},
	})
}

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

	helpers.WriteJSON(w, http.StatusOK, helpers.Response[*dto2.UserResponse]{
		Success: true,
		Data:    usr,
	})
}

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

	helpers.WriteJSON(w, http.StatusOK, helpers.Response[*dto2.UserResponse]{
		Success: true,
		Data:    usr,
	})
}

func (c *UserController) ListCountries(w http.ResponseWriter, r *http.Request) {
	helpers.WriteJSON(w, http.StatusOK, helpers.Response[[]shared.Country]{
		Success: true,
		Data:    shared.ListCountries(),
	})
}

func (c *UserController) ListLanguages(w http.ResponseWriter, r *http.Request) {
	helpers.WriteJSON(w, http.StatusOK, helpers.Response[[]shared.Language]{
		Success: true,
		Data:    shared.ListLanguages(),
	})
}

func (c *UserController) RegisterRoutes(r chi.Router) {
	r.Post("/create", c.CreateUser)
	r.Post("/login", c.Login)

	r.Group(func(r chi.Router) {
		r.Use(middleware.JWTMiddleware(&c.jwtService))

		r.Get("/all", c.GetAllUsers)
		r.Get("/list", c.GetListUsers)
		r.Get("/countries", c.ListCountries)
		r.Get("/languages", c.ListLanguages)
		r.Get("/id/{id}", c.GetUserById)
		r.Get("/email/{email}", c.GetUserByEmail)
		r.Get("/username/{username}", c.GetUserByUsername)
		r.Get("/country/{country}", c.GetUsersByCountry)
		r.Get("/language/{language}", c.GetUsersByLanguage)
		r.Put("/", c.UpdateUser)
		r.Delete("/{id}", c.DeleteUser)
		r.Patch("/profilepic/{id}", c.UploadProfilePic)
		r.Patch("/restore/{id}", c.RestoreUser)
	})
}
