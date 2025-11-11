package shared

import (
	"errors"
	"fmt"
	"time"
)

type BirthDate struct {
	birth time.Time
}

var (
	ErrEmptyBirth  = errors.New("birthdate cannot be nil")
	ErrFutureDate  = errors.New("birthdate cannot be in the future")
	ErrUnderTwelve = errors.New("user must be at least 12 years old")
)

func NewBirthDate(birth time.Time) (BirthDate, error) {
	if birth.After(time.Now()) {
		return BirthDate{}, fmt.Errorf("%w: got %v", ErrFutureDate, birth)
	}
	if !isAnAdult(birth) {
		return BirthDate{}, fmt.Errorf("%w: got %v", ErrUnderTwelve, birth)
	}
	return BirthDate{birth: birth}, nil
}

func isAnAdult(v time.Time) bool {
	yearsAgo := time.Now().AddDate(-18, 0, 0)

	return !v.After(yearsAgo)
}

func (birth BirthDate) Time() time.Time {
	return birth.birth
}
