package shared

import (
	"errors"
	"fmt"
)

type Username struct {
	name string
}

var (
	ErrEmptyUsername = errors.New("username cannot be empty")
	ErrLongUsername  = errors.New("username is too long, maximum size is 30")
	ErrShortUsername = errors.New("username is too short, minimum size is 3")
	ErrInvalidNName  = errors.New("username contains invalid characters")
)

func NewUsername(name string) (Username, error) {
	if name == "" {
		return Username{}, ErrEmptyUsername
	}
	if len(name) > 30 {
		return Username{}, fmt.Errorf("%w: got %s", ErrLongUsername, name)
	}
	if len(name) < 3 {
		return Username{}, fmt.Errorf("%w: got %s", ErrShortUsername, name)
	}
	for _, char := range name {
		if (char < 'a' || char > 'z') && (char < 'A' || char > 'Z') && (char < '0' || char > '9') && char != '_' && char != '.' {
			return Username{}, fmt.Errorf("%w: got %s", ErrInvalidNName, name)
		}
	}
	return Username{name: name}, nil
}

func (u Username) String() string {
	return u.name
}
