package shared

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

type Phone struct {
	dial  DialCode
	value string
	total string
}

var (
	ErrNotNumericPhoneNumber = errors.New("phone number isn't entirely numeric")
	ErrShortPhoneNumber      = errors.New("phone number is too short, minimum size is 8")
	ErrLongPhoneNumber       = errors.New("phone number is too long, maximum size is 10")
)

func NewPhone(phone *string) (*Phone, error) {
	if phone == nil || *phone == "" {
		return nil, nil
	}

	if !isNumericPhone(*phone) {
		return nil, fmt.Errorf("%w: got %s", ErrNotNumericPhoneNumber, *phone)
	}

	parts := strings.Split(*phone, "-")
	dialPart, valuePart := parts[0], parts[1]
	if len(valuePart) < 8 {
		return nil, fmt.Errorf("%w: got %s", ErrShortPhoneNumber, *phone)
	}
	if len(valuePart) > 12 {
		return nil, fmt.Errorf("%w: got %s", ErrLongPhoneNumber, *phone)
	}

	dc, _ := ParseDialCode(dialPart)
	return &Phone{dial: dc, value: valuePart, total: *phone}, nil
}

func isNumericPhone(str string) bool {
	if str[0] != '+' {
		return false
	}
	for i := 1; i < len(str); i++ {
		if rune(str[i]) == '-' {
			continue
		}
		if !unicode.IsDigit(rune(str[i])) {
			return false
		}
	}
	return true
}

func (p Phone) String() string {
	return p.total
}

func (p Phone) Value() string {
	return p.value
}

func (p Phone) Dial() DialCode {
	return p.dial
}
