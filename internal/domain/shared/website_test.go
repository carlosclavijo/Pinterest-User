package shared

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewWebSite(t *testing.T) {
	str := "https://www.mywebsite.com/profile?v=21Asdfij84"

	website, err := NewWebSite(&str)

	require.NotEmpty(t, website)
	require.NoError(t, err)

	assert.Equal(t, str, website.String())
}

func TestWebsite_Empty(t *testing.T) {
	str := ""

	website, err := NewWebSite(&str)

	require.Nil(t, website)
	require.NoError(t, err)
}

func TestWebsite_Nil(t *testing.T) {
	website, err := NewWebSite(nil)

	require.Nil(t, website)
	require.NoError(t, err)
}

func TestWebsite_Invalid(t *testing.T) {
	str := "invalid"

	website, err := NewWebSite(&str)

	require.Nil(t, website)
	require.ErrorIs(t, err, ErrInvalidWebsite)
}
