package shared

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewUsername(t *testing.T) {
	str := "carlosclavijo"

	username, err := NewUsername(str)

	require.NotEmpty(t, username)
	require.NoError(t, err)
	assert.Equal(t, str, username.String())
}

func TestUsername_Empty(t *testing.T) {
	str := ""

	username, err := NewUsername(str)

	require.NotNil(t, username)
	require.Empty(t, username)
	require.ErrorIs(t, err, ErrEmptyUsername)
}

func TestUsername_Long(t *testing.T) {
	str := "longusernamemostlongerthanthirtycharacters"

	username, err := NewUsername(str)

	require.NotNil(t, username)
	require.Empty(t, username)
	require.ErrorIs(t, err, ErrLongUsername)
}

func TestUsername_Short(t *testing.T) {
	str := "sh"

	username, err := NewUsername(str)

	require.NotNil(t, username)
	require.Empty(t, username)
	require.ErrorIs(t, err, ErrShortUsername)
}

func TestUsername_Invalid(t *testing.T) {
	str := "?????????????"

	username, err := NewUsername(str)

	require.NotNil(t, username)
	require.Empty(t, username)
	require.ErrorIs(t, err, ErrInvalidUsername)
}
