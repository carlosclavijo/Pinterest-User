package user

import (
	"errors"
	"github.com/carlosclavijo/Pinterest-User/internal/domain/abstractions"
	"github.com/carlosclavijo/Pinterest-User/internal/domain/shared"
	"github.com/google/uuid"
	"time"
)

type User struct {
	*abstractions.AggregateRoot
	firstName   string
	lastName    string
	username    shared.Username
	email       shared.Email
	password    shared.Password
	gender      shared.Gender
	birth       shared.BirthDate
	country     shared.Country
	language    shared.Language
	phone       *shared.Phone
	information *string
	profilePic  *string
	webSite     *string
	visibility  bool
	createdAt   time.Time
	updatedAt   time.Time
	deletedAt   *time.Time
}

var (
	ErrIdNilUser             = errors.New("id cannot be nil")
	ErrEmptyFirstNameUser    = errors.New("first name cannot be empty")
	ErrEmptyLastNameUser     = errors.New("last name cannot be empty")
	ErrLongFirstNameUser     = errors.New("first name cannot be longer than 100 characters")
	ErrLongLastNameUser      = errors.New("last name cannot be longer than 100 characters")
	ErrNonAlphaFirstNameUser = errors.New("first name has non alphabetical characters")
	ErrNonAlphaLastNameUser  = errors.New("last name has non alphabetical characters")
	ErrNotFoundUser          = errors.New("user not found")
)

func NewUser(firstName, lastName string, username shared.Username, email shared.Email, password shared.Password, gender shared.Gender, birth shared.BirthDate, country shared.Country, language shared.Language, phone *shared.Phone) *User {
	return &User{
		AggregateRoot: abstractions.NewAggregateRoot(uuid.New()),
		firstName:     firstName,
		lastName:      lastName,
		username:      username,
		email:         email,
		password:      password,
		gender:        gender,
		birth:         birth,
		country:       country,
		language:      language,
		phone:         phone,
	}
}

func (u *User) Id() uuid.UUID {
	return u.Entity.Id
}

func (u *User) FirstName() string {
	return u.firstName
}

func (u *User) LastName() string {
	return u.LastName()
}

func (u *User) Username() shared.Username {
	return u.username
}

func (u *User) Email() shared.Email {
	return u.email
}

func (u *User) Gender() shared.Gender {
	return u.gender
}

func (u *User) Birth() shared.BirthDate {
	return u.birth
}

func (u *User) Phone() *shared.Phone {
	return u.phone
}

func (u *User) Country() shared.Country {
	return u.country
}

func (u *User) Language() shared.Language {
	return u.language
}

func (u *User) Information() *string {
	return u.information
}

func (u *User) ProfilePic() *string {
	return u.profilePic
}

func (u *User) WebSite() *string {
	return u.webSite
}

func (u *User) Visibility() bool {
	return u.visibility
}

func (u *User) CreatedAt() time.Time {
	return u.createdAt
}

func (u *User) UpdatedAt() time.Time {
	return u.updatedAt
}

func (u *User) DeletedAt() *time.Time {
	return u.deletedAt
}

func NewUserFromDB(id uuid.UUID, firstName, lastName string, username shared.Username, email shared.Email, password shared.Password, gender shared.Gender, birth shared.BirthDate, country shared.Country, language shared.Language, phone *shared.Phone, information, profilePic, webSite *string, visibility bool, createdAt, updatedAt time.Time, deletedAt *time.Time) *User {
	return &User{
		AggregateRoot: abstractions.NewAggregateRoot(id),
		firstName:     firstName,
		lastName:      lastName,
		username:      username,
		email:         email,
		password:      password,
		gender:        gender,
		birth:         birth,
		country:       country,
		language:      language,
		phone:         phone,
		information:   information,
		profilePic:    profilePic,
		webSite:       webSite,
		visibility:    visibility,
		createdAt:     createdAt,
		updatedAt:     updatedAt,
		deletedAt:     deletedAt,
	}
}
