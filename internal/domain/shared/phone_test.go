package shared

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewPhone(t *testing.T) {
	str := "+591-70926048"

	phone, err := NewPhone(&str)

	require.NotNil(t, phone)
	require.NoError(t, err)

	assert.Equal(t, str, phone.String())
	assert.Equal(t, "+591", string(phone.Dial()))
	assert.Equal(t, "70926048", phone.Value())
}

func TestPhone_Empty(t *testing.T) {
	str := ""

	phone, err := NewPhone(&str)

	require.Nil(t, phone)
	require.NoError(t, err)
}

func TestPhone_Nil(t *testing.T) {
	phone, err := NewPhone(nil)

	require.Nil(t, phone)
	require.NoError(t, err)
}

func TestPhone_NotNumeric(t *testing.T) {
	str := "+notNumeric"

	phone, err := NewPhone(&str)

	require.Nil(t, phone)
	require.ErrorIs(t, err, ErrNotNumericPhoneNumber)
}

func TestPhone_Short(t *testing.T) {
	str := "+7-78989"

	phone, err := NewPhone(&str)

	require.Nil(t, phone)
	require.ErrorIs(t, err, ErrShortPhoneNumber)
}

func TestPhone_Long(t *testing.T) {
	str := "+1-778878787878787887878"

	phone, err := NewPhone(&str)

	require.Nil(t, phone)
	require.ErrorIs(t, err, ErrLongPhoneNumber)
}

func TestIsNumeric(t *testing.T) {
	str := "+591-70826048"
	isNumeric := isNumericPhone(str)

	assert.True(t, isNumeric)

	str = "-5ja0f32f"
	isNumeric = isNumericPhone(str)

	assert.False(t, isNumeric)

}
