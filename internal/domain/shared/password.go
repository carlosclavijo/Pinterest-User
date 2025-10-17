package shared

import (
	"errors"
	"fmt"
	"regexp"
)

type Password struct {
	value string
}

var (
	ErrEmptyPassword = errors.New("password cannot be empty")
	ErrLongPassword  = errors.New("password is too long, maximum size is 64 characters")
	ErrShortPassword = errors.New("password is too short, minimum size is 8 characters")
	ErrSoftPassword  = errors.New("password isn't strong enough")

	ErrEmptyHashedPassword   = errors.New("hashed password cannot be empty")
	ErrInvalidHashedPassword = errors.New("invalid hashed password format")
	ErrLengthHashedPassword  = errors.New("unexpected hashed password length")
)

func NewPassword(password string) (Password, error) {
	if password == "" {
		return Password{}, ErrEmptyPassword
	}
	if len(password) > 64 {
		return Password{}, fmt.Errorf("%w: got %s", ErrLongPassword, password)
	}
	if len(password) < 8 {
		return Password{}, fmt.Errorf("%w: got %s", ErrShortPassword, password)
	}
	if !isStrongPassword(password) {
		return Password{}, fmt.Errorf("%w: got %s", ErrSoftPassword, password)
	}
	return Password{password}, nil
}

func NewHashedPassword(password string) (Password, error) {
	if password == "" {
		return Password{}, ErrEmptyHashedPassword
	}

	if !regexp.MustCompile(`^\$2[ayb]\$[0-9]{2}\$[./A-Za-z0-9]{53}$`).MatchString(password) {
		return Password{}, ErrInvalidHashedPassword
	}

	if len(password) != 60 {
		return Password{}, ErrLengthHashedPassword
	}

	return Password{password}, nil
}

func (p Password) String() string {
	return p.value
}

func isStrongPassword(v string) bool {
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(v)
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(v)
	hasDigit := regexp.MustCompile(`[0-9]`).MatchString(v)
	hasSpecial := regexp.MustCompile(`[!@#$%^&*()_\-+=<>?]`).MatchString(v)

	return hasLower && hasUpper && hasDigit && hasSpecial
}
