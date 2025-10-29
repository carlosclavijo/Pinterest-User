package shared

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestNewBirthDate(t *testing.T) {
	date := time.Now().AddDate(-20, 0, 0)

	birth, err := NewBirthDate(date)

	require.NotEmpty(t, birth)
	require.NoError(t, err)
	assert.Equal(t, date, birth.Time())
}

func TestBirthDate_Future(t *testing.T) {
	date := time.Now().AddDate(10, 0, 0)

	birth, err := NewBirthDate(date)

	require.NotNil(t, birth)
	require.Empty(t, birth)
	assert.ErrorIs(t, err, ErrFutureDate)
}

func TestBirthDate_Underage(t *testing.T) {
	date := time.Now().AddDate(-10, 0, 0)

	birth, err := NewBirthDate(date)

	require.NotNil(t, birth)
	require.Empty(t, birth)
	assert.ErrorIs(t, err, ErrUnderTwelve)
}
