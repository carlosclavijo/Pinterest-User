package user

import (
	"fmt"
	"github.com/carlosclavijo/Pinterest-User/internal/domain/shared"
	"log"
	"strings"
	"unicode"
)

type UserFactory interface {
	Create(firstName, lastName string, username shared.Username, email shared.Email, password shared.Password, gender shared.Gender, birth shared.BirthDate, country shared.Country, language shared.Language, phone *shared.Phone) (*User, error)
}

type userFactory struct{}

func (u userFactory) Create(firstName, lastName string, username shared.Username, email shared.Email, password shared.Password, gender shared.Gender, birth shared.BirthDate, country shared.Country, language shared.Language, phone *shared.Phone) (*User, error) {
	if firstName == "" {
		log.Printf("[factory:administrator] firstName '%s' is empty", firstName)
		return nil, ErrEmptyFirstNameUser
	}

	if lastName == "" {
		log.Printf("[factory:administrator] lastName '%s' is empty", lastName)
		return nil, ErrEmptyLastNameUser
	}

	if len(firstName) > 100 {
		log.Printf("[factory:administrator] firstName '%s' is too long, length %d, maximum is 100)", firstName, len(firstName))
		return nil, fmt.Errorf("%w: got %s, size %d", ErrLongFirstNameUser, firstName, len(firstName))
	}

	if len(lastName) > 100 {
		log.Printf("[factory:administrator] lastName '%s' is too long (length %d, maximum is 100)", lastName, len(lastName))
		return nil, fmt.Errorf("%w: got %s, size %d", ErrLongLastNameUser, lastName, len(lastName))
	}

	if !isAlpha(firstName) {
		log.Printf("[factory:administrator] firstName '%s' contains non-alphabetic characters", firstName)
		return nil, fmt.Errorf("%w: got %s", ErrNonAlphaFirstNameUser, firstName)
	}

	if !isAlpha(lastName) {
		log.Printf("[factory:administrator] lastName '%s' contains non-alphabetic characters", lastName)
		return nil, fmt.Errorf("%w: got %s", ErrNonAlphaLastNameUser, lastName)
	}

	log.Printf("[factory:administrator][SUCCESS] administrator created")
	return NewUser(firstName, lastName, username, email, password, gender, birth, country, language, phone), nil
}

func NewUserFactory() UserFactory {
	return &userFactory{}
}

func isAlpha(s string) bool {
	s = strings.TrimSpace(s)
	if s == "" {
		return false
	}

	previousWasSpace := false
	for _, r := range s {
		if unicode.IsLetter(r) {
			previousWasSpace = false
			continue
		} else if r == ' ' {
			if previousWasSpace {
				return false
			}
			previousWasSpace = true
		} else {
			return false
		}
	}

	return true
}
