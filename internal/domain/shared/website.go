package shared

import (
	"errors"
	"fmt"
	"regexp"
)

type Website struct {
	value string
}

const regexWebsite = `^(https?:\/\/)?([\w\-]+\.)+[\w\-]+(\/[\w\-._~:/?#[\]@!$&'()*+,;=]*)?$`

var (
	ErrEmptyWebsite   = errors.New("website cannot be empty")
	ErrInvalidWebsite = errors.New("invalid website format")
)

func NewWebSite(website string) (Website, error) {
	if website == "" {
		return Website{}, ErrEmptyWebsite
	}
	if !regexp.MustCompile(regexWebsite).MatchString(website) {
		return Website{}, fmt.Errorf("%w: got %s", ErrInvalidWebsite, website)
	}
	return Website{value: website}, nil
}

func (w Website) String() string {
	return w.value
}
