package users

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
	webSite     *shared.Website
	visibility  bool
	lastLoginAt time.Time
	createdAt   time.Time
	updatedAt   time.Time
	deletedAt   *time.Time
}

var (
	ErrIdNilUser              = errors.New("id cannot be nil")
	ErrEmptyFirstNameUser     = errors.New("first name cannot be empty")
	ErrEmptyLastNameUser      = errors.New("last name cannot be empty")
	ErrLongFirstNameUser      = errors.New("first name cannot be longer than 100 characters")
	ErrLongLastNameUser       = errors.New("last name cannot be longer than 100 characters")
	ErrNonAlphaFirstNameUser  = errors.New("first name has non alphabetical characters")
	ErrNonAlphaLastNameUser   = errors.New("last name has non alphabetical characters")
	ErrLongInformationUser    = errors.New("information cannot be longer than 500 characters")
	ErrNotFoundUser           = errors.New("user not found")
	ErrExistsUser             = errors.New("user already exist")
	ErrAlreadyDeletedUser     = errors.New("user already deleted")
	ErrAlreadyRestoredUser    = errors.New("user already restored")
	ErrInvalidCredentialsUser = errors.New("invalid credentials")
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
		lastLoginAt:   time.Now(),
		createdAt:     time.Now(),
		updatedAt:     time.Now(),
	}
}

func (u *User) Id() uuid.UUID {
	return u.Entity.Id
}

func (u *User) FirstName() string {
	return u.firstName
}

func (u *User) LastName() string {
	return u.lastName
}

func (u *User) Username() shared.Username {
	return u.username
}

func (u *User) Email() shared.Email {
	return u.email
}

func (u *User) Password() shared.Password {
	return u.password
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

func (u *User) WebSite() *shared.Website {
	return u.webSite
}

func (u *User) Visibility() bool {
	return u.visibility
}

func (u *User) LastLoginAt() time.Time {
	return u.lastLoginAt
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

func (u *User) ChangeFirstName(name string) error {
	if name == "" {
		return ErrEmptyFirstNameUser
	} else if len(name) > 100 {
		return ErrLongFirstNameUser
	}
	u.firstName = name
	return nil
}

func (u *User) ChangeLastName(name string) error {
	if name == "" {
		return ErrEmptyLastNameUser
	} else if len(name) > 100 {
		return ErrLongLastNameUser
	}
	u.lastName = name
	return nil
}

func (u *User) ChangeUsername(username shared.Username) error {
	if username == (shared.Username{}) {
		return shared.ErrEmptyUsername
	}
	u.username = username
	return nil
}

func (u *User) ChangeEmail(email shared.Email) error {
	if email == (shared.Email{}) {
		return shared.ErrEmptyEmail
	}
	u.email = email
	return nil
}

func (u *User) ChangePassword(password shared.Password) error {
	if password == (shared.Password{}) {
		return shared.ErrEmptyPassword
	}
	u.password = password
	return nil
}

func (u *User) ChangeGender(gender shared.Gender) error {
	if gender.String() == "Unknown" {
		return shared.ErrNotAGender
	}
	u.gender = gender
	return nil
}

func (u *User) ChangeBirth(birth shared.BirthDate) error {
	if birth == (shared.BirthDate{}) {
		return shared.ErrEmptyBirth
	}
	u.birth = birth
	return nil
}

func (u *User) ChangeCountry(country shared.Country) error {
	if country.String() == "Unknown" {
		return shared.ErrNotACountry
	}
	u.country = country
	return nil
}

func (u *User) ChangeLanguage(language shared.Language) error {
	if language.String() == "Unknown" {
		return shared.ErrNotALanguage
	}
	u.language = language
	return nil
}

func (u *User) ChangePhone(phone *shared.Phone) {
	u.phone = phone
}

func (u *User) ChangeInformation(information *string) error {
	if information != nil {
		if len(*information) > 500 {
			return ErrLongInformationUser
		}
	}
	u.information = information
	return nil
}

func (u *User) ChangeProfilePic(profilePic *string) {
	u.profilePic = profilePic
}

func (u *User) ChangeWebSite(webSite *shared.Website) {
	u.webSite = webSite
}

func (u *User) ChangeVisibility(visibility bool) {
	u.visibility = visibility
}

func (u *User) ChangeLastLoginAt() {
	u.lastLoginAt = time.Now()
}

func (u *User) ChangeUpdatedAt() {
	u.updatedAt = time.Now()
}

func (u *User) Delete() error {
	if u.deletedAt != nil {
		return ErrAlreadyDeletedUser
	}

	now := time.Now()
	u.deletedAt = &now

	return nil
}

func (u *User) Restore() error {
	if u.deletedAt == nil {
		return ErrAlreadyRestoredUser
	}

	u.deletedAt = nil

	return nil
}

func NewUserFromDB(id uuid.UUID, firstName, lastName, username, email, password, gender string, birth time.Time, country, language string, phone, information, profilePic, webSite *string, visibility bool, lastLoginAt time.Time, createdAt, updatedAt time.Time, deletedAt *time.Time) (*User, error) {
	Username, err := shared.NewUsername(username)
	if err != nil {
		return nil, err
	}

	Email, err := shared.NewEmail(email)
	if err != nil {
		return nil, err
	}

	Password, err := shared.NewPassword(password)
	if err != nil {
		return nil, err
	}

	Gender, err := shared.ParseGender(gender)
	if err != nil {
		return nil, err
	}

	Birth, err := shared.NewBirthDate(birth)
	if err != nil {
		return nil, err
	}

	Country, err := shared.ParseCountry(country)
	if err != nil {
		return nil, err
	}

	Language, err := shared.ParseLanguage(language)
	if err != nil {
		return nil, err
	}

	Phone, err := shared.NewPhone(phone)
	if err != nil {
		return nil, err
	}

	WebSite, err := shared.NewWebSite(webSite)
	if err != nil {
		return nil, err
	}

	return &User{
		AggregateRoot: abstractions.NewAggregateRoot(id),
		firstName:     firstName,
		lastName:      lastName,
		username:      Username,
		email:         Email,
		password:      Password,
		gender:        Gender,
		birth:         Birth,
		country:       Country,
		language:      Language,
		phone:         Phone,
		information:   information,
		profilePic:    profilePic,
		webSite:       WebSite,
		visibility:    visibility,
		lastLoginAt:   lastLoginAt,
		createdAt:     createdAt,
		updatedAt:     updatedAt,
		deletedAt:     deletedAt,
	}, nil
}
