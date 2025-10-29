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
	ErrInvalidWebsite = errors.New("invalid website format")
)

func NewWebSite(website *string) (*Website, error) {
	if website == nil || *website == "" {
		return nil, nil
	}
	if !regexp.MustCompile(regexWebsite).MatchString(*website) {
		return nil, fmt.Errorf("%w: got %v", ErrInvalidWebsite, website)
	}
	return &Website{value: *website}, nil
}

func (website Website) String() string {
	return website.value
}
