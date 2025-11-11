package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/carlosclavijo/Pinterest-Services/internal/domain/user"
	"github.com/google/uuid"
	"time"
)

const (
	got              = "%w: got %w"
	QueryGetAllUsers = `SELECT id, first_name, last_name, user_name, email, password, gender, birth_date, country, language, phone, information, profile_pic, web_site, visibility, last_login_at, created_at, updated_at, deleted_at
						FROM users`
	QueryGetListUsers = `SELECT id, first_name, last_name, user_name, email, password, gender, birth_date, country, language, phone, information, profile_pic, web_site, visibility, last_login_at, created_at, updated_at, deleted_at
						FROM users
						WHERE deleted_at IS NULL`
	QueryGetUserById = `SELECT first_name, last_name, user_name, email, password, gender, birth_date, country, language, phone, information, profile_pic, web_site, visibility, last_login_at, created_at, updated_at, deleted_at
						FROM users
						WHERE id = $1`
	QueryGetUserByUsername = `SELECT id, first_name, last_name, email, password, gender, birth_date, country, language, phone, information, profile_pic, web_site, visibility, last_login_at, created_at, updated_at, deleted_at
						  	  FROM users 
						  	  WHERE user_name = $1`
	QueryGetUserByEmail = `SELECT id, first_name, last_name, user_name, password, gender, birth_date, country, language, phone, information, profile_pic, web_site, visibility, last_login_at, created_at, updated_at, deleted_at
						   FROM users 
						   WHERE email = $1`
	QueryGetUsersByCountry = `SELECT id, first_name, last_name, user_name, email, password, gender, birth_date, language, phone, information, profile_pic, web_site, visibility, last_login_at, created_at, updated_at, deleted_at
							 	  FROM users
							 	  WHERE country = $1`
	QueryGetUsersByLanguage = `SELECT id, first_name, last_name, user_name, email, password, gender, birth_date, country, phone, information, profile_pic, web_site, visibility, last_login_at, created_at, updated_at, deleted_at
							  	   FROM users
							  	   WHERE language = $1`
	QueryGetUsersLikeUsername = `SELECT id, first_name, last_name, user_name, email, password, gender, birth_date, country, phone, information, profile_pic, web_site, visibility, last_login_at, created_at, updated_at, deleted_at
							  	   FROM users
							  	   WHERE language ILIKE '%' || $1 || '%' AND deleted_at IS NULL`
	QueryExistUserById = `SELECT EXISTS(
							SELECT 1 
							FROM users
				  			WHERE id = $1 
				  			AND deleted_at IS NULL
				  		  )`
	QueryExistUserByUsername = `SELECT EXISTS(
								SELECT 1 
								FROM users
				  				WHERE user_name = $1 
				  				AND deleted_at IS NULL
				  		  	)`
	QueryExistUserByEmail = `SELECT EXISTS(
								SELECT 1 
								FROM users
				  				WHERE email = $1 
				  				AND deleted_at IS NULL
				  		  	)`
	QueryCreateUser = `INSERT INTO users(id, first_name, last_name, user_name, email, password, gender, birth_date, country, language, phone, visibility, last_login_at, created_at, updated_at)
					   VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
					   RETURNING id, first_name, last_name, user_name, email, password, gender, birth_date, country, language, phone, information, profile_pic, web_site, visibility, last_login_at, created_at, updated_at, deleted_at`
	QueryUpdateUser = `UPDATE users
					   SET first_name = $2, last_name = $3, user_name = $4, email = $5, password = $6, gender = $7, birth_date = $8, country = $9, language = $10, phone = $11, information = $12, profile_pic = $13, web_site = $14, visibility = $15, last_login_at = $16, updated_at = $17
					   WHERE id = $1`
	QueryDeleteUser = `UPDATE users
					   SET deleted_at = $2
					   WHERE id = $1`
)

var (
	ErrQuery         = errors.New("query failed")
	ErrScan          = errors.New("scan failed")
	ErrConcatenating = errors.New("error concatenating values from DB")
	ErrIterationRows = errors.New("rows iteration error")
)

type userRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) users.UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (r *userRepository) GetAll(ctx context.Context) ([]*users.User, error) {
	var (
		usersList                                                                 []*users.User
		id                                                                        uuid.UUID
		firstName, lastName, username, email, password, gender, country, language string
		birth, lastLoginAt, createdAt, updatedAt                                  time.Time
		phone, information, profilePic, webSite                                   *string
		visibility                                                                bool
		deletedAt                                                                 *time.Time
	)

	rows, err := r.DB.QueryContext(ctx, QueryGetAllUsers)
	if err != nil {
		return nil, fmt.Errorf(got, ErrQuery, err)
	}

	defer func(rows *sql.Rows) {
		if err = rows.Close(); err != nil {
			return
		}
	}(rows)
	for rows.Next() {
		err = rows.Scan(&id, &firstName, &lastName, &username, &email, &password, &gender, &birth, &country, &language, &phone, &information, &profilePic, &webSite, &visibility, &lastLoginAt, &createdAt, &updatedAt, &deletedAt)
		if err != nil {
			return nil, fmt.Errorf(got, ErrScan, err)
		}

		usr, err := users.NewUserFromDB(id, firstName, lastName, username, email, password, gender, birth, country, language, phone, information, profilePic, webSite, visibility, lastLoginAt, createdAt, updatedAt, deletedAt)
		if err != nil {
			return nil, fmt.Errorf(got, ErrConcatenating, err)
		}

		usersList = append(usersList, usr)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf(got, ErrIterationRows, err)
	}

	return usersList, nil
}

func (r *userRepository) GetList(ctx context.Context) ([]*users.User, error) {
	var (
		usersList                                                                 []*users.User
		id                                                                        uuid.UUID
		firstName, lastName, username, email, password, gender, country, language string
		birth, lastLoginAt, createdAt, updatedAt                                  time.Time
		phone, information, profilePic, webSite                                   *string
		visibility                                                                bool
		deletedAt                                                                 *time.Time
	)

	rows, err := r.DB.QueryContext(ctx, QueryGetListUsers)
	if err != nil {
		return nil, fmt.Errorf(got, ErrQuery, err)
	}

	defer func(rows *sql.Rows) {
		if err = rows.Close(); err != nil {
			return
		}
	}(rows)
	for rows.Next() {
		err = rows.Scan(
			&id, &firstName, &lastName, &username, &email, &password, &gender, &birth, &country, &language, &phone, &information, &profilePic, &webSite, &visibility, &lastLoginAt, &createdAt, &updatedAt, &deletedAt,
		)
		if err != nil {
			return nil, fmt.Errorf(got, ErrScan, err)
		}

		usr, err := users.NewUserFromDB(id, firstName, lastName, username, email, password, gender, birth, country, language, phone, information, profilePic, webSite, visibility, lastLoginAt, createdAt, updatedAt, deletedAt)
		if err != nil {
			return nil, fmt.Errorf(got, ErrConcatenating, err)
		}

		usersList = append(usersList, usr)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf(got, ErrIterationRows, err)
	}

	return usersList, nil
}

func (r *userRepository) GetById(ctx context.Context, id uuid.UUID) (*users.User, error) {
	var (
		firstName, lastName, username, email, password, gender, country, language string
		birth, lastLoginAt, createdAt, updatedAt                                  time.Time
		phone, information, profilePic, webSite                                   *string
		visibility                                                                bool
		deletedAt                                                                 *time.Time
	)

	err := r.DB.QueryRowContext(ctx, QueryGetUserById, id).Scan(
		&firstName, &lastName, &username, &email, &password, &gender, &birth, &country, &language, &phone, &information, &profilePic, &webSite, &visibility, &lastLoginAt, &createdAt, &updatedAt, &deletedAt,
	)
	if err != nil {
		return nil, fmt.Errorf(got, ErrQuery, err)
	}

	usr, err := users.NewUserFromDB(id, firstName, lastName, username, email, password, gender, birth, country, language, phone, information, profilePic, webSite, visibility, lastLoginAt, createdAt, updatedAt, deletedAt)
	if err != nil {
		return nil, fmt.Errorf(got, ErrConcatenating, err)
	}

	return usr, nil
}

func (r *userRepository) GetByUsername(ctx context.Context, username string) (*users.User, error) {
	var (
		id                                                              uuid.UUID
		firstName, lastName, email, password, gender, country, language string
		birth, lastLoginAt, createdAt, updatedAt                        time.Time
		phone, information, profilePic, webSite                         *string
		visibility                                                      bool
		deletedAt                                                       *time.Time
	)

	err := r.DB.QueryRowContext(ctx, QueryGetUserByUsername, username).Scan(
		&id, &firstName, &lastName, &email, &password, &gender, &birth, &country, &language, &phone, &information, &profilePic, &webSite, &visibility, &lastLoginAt, &createdAt, &updatedAt, &deletedAt,
	)
	if err != nil {
		return nil, fmt.Errorf(got, ErrQuery, err)
	}

	usr, err := users.NewUserFromDB(id, firstName, lastName, username, email, password, gender, birth, country, language, phone, information, profilePic, webSite, visibility, lastLoginAt, createdAt, updatedAt, deletedAt)
	if err != nil {
		return nil, fmt.Errorf(got, ErrConcatenating, err)
	}

	return usr, nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*users.User, error) {
	var (
		id                                                                 uuid.UUID
		firstName, lastName, username, password, gender, country, language string
		birth, lastLoginAt, createdAt, updatedAt                           time.Time
		phone, information, profilePic, webSite                            *string
		visibility                                                         bool
		deletedAt                                                          *time.Time
	)

	err := r.DB.QueryRowContext(ctx, QueryGetUserByEmail, email).Scan(
		&id, &firstName, &lastName, &username, &password, &gender, &birth, &country, &language, &phone, &information, &profilePic, &webSite, &visibility, &lastLoginAt, &createdAt, &updatedAt, &deletedAt,
	)
	if err != nil {
		return nil, fmt.Errorf(got, ErrQuery, err)
	}

	usr, err := users.NewUserFromDB(id, firstName, lastName, username, email, password, gender, birth, country, language, phone, information, profilePic, webSite, visibility, lastLoginAt, createdAt, updatedAt, deletedAt)
	if err != nil {
		return nil, fmt.Errorf(got, ErrConcatenating, err)
	}

	return usr, nil
}

func (r *userRepository) GetListByCountry(ctx context.Context, country string) ([]*users.User, error) {
	var (
		usersList                                                        []*users.User
		id                                                               uuid.UUID
		firstName, lastName, username, email, password, gender, language string
		birth, lastLoginAt, createdAt, updatedAt                         time.Time
		phone, information, profilePic, webSite                          *string
		visibility                                                       bool
		deletedAt                                                        *time.Time
	)

	rows, err := r.DB.QueryContext(ctx, QueryGetUsersByCountry, country)
	if err != nil {
		return nil, fmt.Errorf(got, ErrQuery, err)
	}

	defer func(rows *sql.Rows) {
		if err = rows.Close(); err != nil {
			return
		}
	}(rows)
	for rows.Next() {
		err = rows.Scan(&id, &firstName, &lastName, &username, &email, &password, &gender, &birth, &language, &phone, &information, &profilePic, &webSite, &visibility, &lastLoginAt, &createdAt, &updatedAt, &deletedAt)
		if err != nil {
			return nil, fmt.Errorf(got, ErrScan, err)
		}

		usr, err := users.NewUserFromDB(id, firstName, lastName, username, email, password, gender, birth, country, language, phone, information, profilePic, webSite, visibility, lastLoginAt, createdAt, updatedAt, deletedAt)
		if err != nil {
			return nil, fmt.Errorf(got, ErrConcatenating, err)
		}

		usersList = append(usersList, usr)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf(got, ErrIterationRows, err)
	}

	return usersList, nil
}

func (r *userRepository) GetListByLanguage(ctx context.Context, language string) ([]*users.User, error) {
	var (
		usersList                                                       []*users.User
		id                                                              uuid.UUID
		firstName, lastName, username, email, password, gender, country string
		birth, lastLoginAt, createdAt, updatedAt                        time.Time
		phone, information, profilePic, webSite                         *string
		visibility                                                      bool
		deletedAt                                                       *time.Time
	)

	rows, err := r.DB.QueryContext(ctx, QueryGetUsersByLanguage, language)
	if err != nil {
		return nil, fmt.Errorf(got, ErrQuery, err)
	}

	defer func(rows *sql.Rows) {
		if err = rows.Close(); err != nil {
			return
		}
	}(rows)
	for rows.Next() {
		err = rows.Scan(&id, &firstName, &lastName, &username, &email, &password, &gender, &birth, &country, &phone, &information, &profilePic, &webSite, &visibility, &lastLoginAt, &createdAt, &updatedAt, &deletedAt)
		if err != nil {
			return nil, fmt.Errorf(got, ErrScan, err)
		}

		usr, err := users.NewUserFromDB(id, firstName, lastName, username, email, password, gender, birth, country, language, phone, information, profilePic, webSite, visibility, lastLoginAt, createdAt, updatedAt, deletedAt)
		if err != nil {
			return nil, fmt.Errorf(got, ErrConcatenating, err)
		}

		usersList = append(usersList, usr)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf(got, ErrIterationRows, err)
	}

	return usersList, nil
}

func (r *userRepository) GetListLikeUsername(ctx context.Context, name string) ([]*users.User, error) {
	var (
		usersList                                                                 []*users.User
		id                                                                        uuid.UUID
		firstName, lastName, username, email, password, gender, country, language string
		birth, lastLoginAt, createdAt, updatedAt                                  time.Time
		phone, information, profilePic, webSite                                   *string
		visibility                                                                bool
		deletedAt                                                                 *time.Time
	)

	rows, err := r.DB.QueryContext(ctx, QueryGetUsersLikeUsername, name)
	if err != nil {
		return nil, fmt.Errorf(got, ErrQuery, err)
	}

	defer func(rows *sql.Rows) {
		if err = rows.Close(); err != nil {
			return
		}
	}(rows)
	for rows.Next() {
		err = rows.Scan(&id, &firstName, &lastName, &username, &email, &password, &gender, &birth, &country, &language, &phone, &information, &profilePic, &webSite, &visibility, &lastLoginAt, &createdAt, &updatedAt, &deletedAt)
		if err != nil {
			return nil, fmt.Errorf(got, ErrScan, err)
		}

		usr, err := users.NewUserFromDB(id, firstName, lastName, username, email, password, gender, birth, country, language, phone, information, profilePic, webSite, visibility, lastLoginAt, createdAt, updatedAt, deletedAt)
		if err != nil {
			return nil, fmt.Errorf(got, ErrConcatenating, err)
		}

		usersList = append(usersList, usr)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf(got, ErrIterationRows, err)
	}

	return usersList, nil
}

func (r *userRepository) ExistsById(ctx context.Context, id uuid.UUID) (bool, error) {
	var exist bool

	err := r.DB.QueryRowContext(ctx, QueryExistUserById, id).Scan(&exist)
	if err != nil {
		return false, fmt.Errorf(got, ErrQuery, err)
	}

	return exist, nil
}

func (r *userRepository) ExistsByUserName(ctx context.Context, userName string) (bool, error) {
	var exist bool

	err := r.DB.QueryRowContext(ctx, QueryExistUserByUsername, userName).Scan(&exist)
	if err != nil {
		return false, fmt.Errorf(got, ErrQuery, err)
	}

	return exist, nil
}

func (r *userRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	var exist bool

	err := r.DB.QueryRowContext(ctx, QueryExistUserByEmail, email).Scan(&exist)
	if err != nil {
		return false, fmt.Errorf(got, ErrQuery, err)
	}

	return exist, nil
}

func (r *userRepository) Create(ctx context.Context, u *users.User) (*users.User, error) {
	var (
		id                                                                        uuid.UUID
		firstName, lastName, username, email, password, gender, country, language string
		birth, lastLoginAt, createdAt, updatedAt                                  time.Time
		phone, information, profilePic, webSite                                   *string
		visibility                                                                bool
		deletedAt                                                                 *time.Time
	)

	if u.Phone() != nil {
		p := u.Phone().String()
		phone = &p
	}

	err := r.DB.QueryRowContext(ctx, QueryCreateUser,
		u.Id(), u.FirstName(), u.LastName(), u.Username().String(), u.Email().String(), u.Password().String(), u.Gender(), u.Birth().Time(), u.Country(), u.Language(), phone, u.Visibility(), u.LastLoginAt(), u.CreatedAt(), u.UpdatedAt(),
	).Scan(
		&id, &firstName, &lastName, &username, &email, &password, &gender, &birth, &country, &language, &phone, &information, &profilePic, &webSite, &visibility, &lastLoginAt, &createdAt, &updatedAt, &deletedAt,
	)

	if err != nil {
		return nil, fmt.Errorf(got, ErrQuery, err)
	}

	usr, err := users.NewUserFromDB(id, firstName, lastName, username, email, password, gender, birth, country, language, phone, information, profilePic, webSite, visibility, lastLoginAt, createdAt, updatedAt, deletedAt)
	if err != nil {
		return nil, fmt.Errorf(got, ErrConcatenating, err)
	}

	return usr, nil
}

func (r *userRepository) Update(ctx context.Context, u *users.User) error {
	var phone, information, profilePic, webSite *string

	if u.Phone() != nil {
		p := u.Phone().String()
		phone = &p
	}
	if u.Information() != nil {
		i := *u.Information()
		information = &i
	}
	if u.ProfilePic() != nil {
		p := *u.ProfilePic()
		profilePic = &p
	}
	if u.WebSite() != nil {
		w := u.WebSite().String()
		webSite = &w
	}

	_, err := r.DB.ExecContext(ctx, QueryUpdateUser,
		u.Id(), u.FirstName(), u.LastName(), u.Username().String(), u.Email().String(), u.Password().String(), u.Gender(), u.Birth().Time(), u.Country(), u.Language(), phone, information, profilePic, webSite, u.Visibility(), u.LastLoginAt(), u.UpdatedAt(),
	)

	if err != nil {
		return fmt.Errorf(got, ErrQuery, err)
	}

	return nil
}

func (r *userRepository) Delete(ctx context.Context, u *users.User) error {
	_, err := r.DB.ExecContext(ctx, QueryDeleteUser, u.Id(), u.DeletedAt())
	if err != nil {
		return fmt.Errorf(got, ErrQuery, err)
	}

	return nil
}
