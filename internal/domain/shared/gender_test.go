package shared

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

type genderCase struct {
	name      string
	genderStr string
	gender    Gender
}

func TestParseGender(t *testing.T) {
	cases := genderCases()
	var genders []Gender

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			g, err := ParseGender(tc.genderStr)
			genders = append(genders, g)

			require.NotEmpty(t, g)
			require.NoError(t, err)
			assert.Equal(t, tc.gender, g)
		})
	}

	assert.Len(t, cases, len(genders))

	for v, g := range genders {
		t.Run(cases[v].name, func(t *testing.T) {
			if len(cases[v].genderStr) > 1 {
				assert.Equal(t, cases[v].genderStr, g.String())
			} else {
				assert.Equal(t, cases[v].gender, g)
			}
		})
	}
}

func TestParseGender_Invalid(t *testing.T) {
	str := "X"

	gender, err := ParseGender(str)

	require.Empty(t, gender)
	require.ErrorIs(t, err, ErrNotAGender)

	genderStr := gender.String()

	assert.Equal(t, "Unknown", genderStr)
}

func TestListGender(t *testing.T) {
	genders := ListGender()

	require.NotEmpty(t, genders)
	assert.Contains(t, genders, Male)
	assert.Contains(t, genders, Female)
	assert.Contains(t, genders, Other)
}

func genderCases() []genderCase {
	cases := []genderCase{
		{"Male 1", "M", Male},
		{"Male 2", "Male", Male},
		{"Female 1", "F", Female},
		{"Female 2", "Female", Female},
		{"Other 1", "O", Other},
		{"Other 2", "Other", Other},
	}

	return cases
}
