package shared

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type Email struct {
	local  string
	domain string
	total  string
}

const emailRegex = `^[a-zA-Z0-9]([a-zA-Z0-9._%+\-]*[a-zA-Z0-9])?@([a-zA-Z0-9]+(-[a-zA-Z0-9]+)*\.)+[a-zA-Z]{2,}$`

var (
	ErrEmptyEmail      = errors.New("email cannot be empty")
	ErrInvalidEmail    = errors.New("invalid email format")
	ErrLongLocalEmail  = errors.New("local part is too long, maximum size is 64")
	ErrLongDomainEmail = errors.New("domain part of email is too long, maximum size is 255")
)

func NewEmail(email string) (Email, error) {
	if email == "" {
		return Email{}, ErrEmptyEmail
	}
	if !regexp.MustCompile(emailRegex).MatchString(email) {
		return Email{}, fmt.Errorf("%w: got %s", ErrInvalidEmail, email)
	}
	parts := strings.Split(email, "@")
	localPart, domainPart := parts[0], parts[1]
	if len(localPart) > 64 {
		return Email{}, fmt.Errorf("%w: got %s", ErrLongLocalEmail, email)
	}
	if len(domainPart) > 255 {
		return Email{}, fmt.Errorf("%w: got %s", ErrLongDomainEmail, email)
	}
	return Email{local: localPart, domain: domainPart, total: email}, nil
}

func (email Email) String() string {
	return email.total
}

func (email Email) Local() string {
	return email.local
}

func (email Email) Domain() string {
	return email.domain
}
