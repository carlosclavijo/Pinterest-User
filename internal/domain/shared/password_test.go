package shared

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewPassword(t *testing.T) {
	str := "5trOng1!."

	password, err := NewPassword(str)

	require.NotEmpty(t, password)
	require.NoError(t, err)
	assert.NotNil(t, password.value)
	assert.Equal(t, str, password.value)
}

func TestPassword_Empty(t *testing.T) {
	str := ""

	password, err := NewPassword(str)

	require.NotNil(t, password)
	require.Empty(t, password)
	require.ErrorIs(t, err, ErrEmptyPassword)
}

func TestPassword_Long(t *testing.T) {
	str := "longpasswordNot!Strong5555.Butalongpasswordpleaseineedtofinishthesetestlmao"

	password, err := NewPassword(str)

	require.NotNil(t, password)
	require.Empty(t, password)
	require.ErrorIs(t, err, ErrLongPassword)
}

func TestPassword_Short(t *testing.T) {
	str := "short"

	password, err := NewPassword(str)

	require.NotNil(t, password)
	require.Empty(t, password)
	require.ErrorIs(t, err, ErrShortPassword)
}

func TestPassword_Soft(t *testing.T) {
	str := "shortpasswordaaaaaaaaaaaaaaa"

	password, err := NewPassword(str)

	require.NotNil(t, password)
	require.Empty(t, password)
	require.ErrorIs(t, err, ErrSoftPassword)
}

func TestNewHashedPassword(t *testing.T) {
	str := "$2a$10$3J9wq7F0s8G2bXHkzQvFqO5tLh8mY2nP4rZxN1uVY3sTq6aKbL1Pa"

	password, err := NewHashedPassword(str)

	require.NotNil(t, password)
	require.NotEmpty(t, password)
	require.NoError(t, err)
	assert.Equal(t, str, password.String())
}

func TestHashedPassword_Empty(t *testing.T) {
	str := ""

	password, err := NewHashedPassword(str)

	require.NotNil(t, password)
	require.Empty(t, password)
	require.ErrorIs(t, err, ErrEmptyHashedPassword)
}

func TestHashedPassword_Invalid(t *testing.T) {
	str := "invalid"

	password, err := NewHashedPassword(str)

	require.NotNil(t, password)
	require.Empty(t, password)
	require.ErrorIs(t, err, ErrInvalidHashedPassword)
}

func TestIsStrongPassword(t *testing.T) {
	str := "invalid"
	isStrong := isStrongPassword(str)

	assert.False(t, isStrong)

	str = "$2a$10$3J9wq7F0s8G2bXHkzQvFqO5tLh8mY2nP4rZxN1uVY3sTq6aKbL1Pa"
	isStrong = isStrongPassword(str)

	assert.True(t, isStrong)
}
