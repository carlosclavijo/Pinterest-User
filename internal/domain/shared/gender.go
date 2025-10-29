package shared

import (
	"errors"
	"fmt"
)

type Gender string

const (
	Male   Gender = "M"
	Female Gender = "F"
	Other  Gender = "O"
)

var ErrNotAGender = errors.New("is not a gender")

func (gender Gender) String() string {
	switch gender {
	case Male:
		return "Male"
	case Female:
		return "Female"
	case Other:
		return "Other"
	default:
		return "Unknown"
	}
}

func ParseGender(gender string) (Gender, error) {
	switch gender {
	case "M", "Male":
		return Male, nil
	case "F", "Female":
		return Female, nil
	case "O", "Other":
		return Other, nil
	default:
		return "", fmt.Errorf("%w: got %s", ErrNotAGender, gender)
	}
}

func ListGender() []Gender {
	return []Gender{
		Male, Female, Other,
	}
}
