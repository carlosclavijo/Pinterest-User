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

func (g Gender) String() string {
	switch g {
	case Male:
		return "male"
	case Female:
		return "female"
	case Other:
		return "other"
	default:
		return "unknown"
	}
}

func ParseGender(gender string) (Gender, error) {
	switch gender {
	case "male", "M":
		return Male, nil
	case "female", "F":
		return Female, nil
	case "other", "O":
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
