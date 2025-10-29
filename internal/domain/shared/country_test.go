package shared

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

type countryCase struct {
	name       string
	countryStr string
	country    Country
}

func TestParseCountry(t *testing.T) {
	cases := countryCases()
	var countries []Country

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			c, err := ParseCountry(tc.countryStr)
			countries = append(countries, c)

			require.NotEqual(t, "", c)
			require.NoError(t, err)
			assert.Equal(t, tc.country, c)
		})
	}

	assert.Len(t, cases, len(countries))

	for v, c := range countries {
		cStr := c.String()
		t.Run(cases[v].name, func(t *testing.T) {
			if len(cases[v].countryStr) != 2 {
				assert.Equal(t, cases[v].countryStr, cStr)
			} else {
				assert.Equal(t, cases[v].country.String(), cStr)
			}
		})
	}
}

func TestParseCountryError(t *testing.T) {
	country, err := ParseCountry("New Zealand")

	require.NotNil(t, country)
	assert.NotEqual(t, "", country.String())
	assert.ErrorIs(t, err, ErrNotACountry)
}

func TestListCountries(t *testing.T) {
	countries := ListCountries()

	require.NotEmpty(t, countries)
	assert.Contains(t, countries, Canada)
	assert.Contains(t, countries, Mexico)
	assert.Contains(t, countries, UnitedStates)
	assert.Contains(t, countries, Argentina)
	assert.Contains(t, countries, Bolivia)
	assert.Contains(t, countries, Brazil)
	assert.Contains(t, countries, UnitedKingdom)
	assert.Contains(t, countries, France)
	assert.Contains(t, countries, Germany)
	assert.Contains(t, countries, Spain)
	assert.Contains(t, countries, Italy)
	assert.Contains(t, countries, Japan)
	assert.Contains(t, countries, China)
	assert.Contains(t, countries, India)
	assert.Contains(t, countries, Australia)
	assert.Contains(t, countries, SouthKorea)
}

func countryCases() []countryCase {
	cases := []countryCase{
		{"Canada 1", "CA", Canada},
		{"Canada 2", "Canada", Canada},
		{"Mexico 1", "MX", Mexico},
		{"Mexico 2", "Mexico", Mexico},
		{"United States 1", "US", UnitedStates},
		{"United States 2", "United States", UnitedStates},
		{"Argentina 1", "AR", Argentina},
		{"Argentina 2", "Argentina", Argentina},
		{"Bolivia 1", "BO", Bolivia},
		{"Bolivia 2", "Bolivia", Bolivia},
		{"Brazil 1", "BR", Brazil},
		{"Brazil 2", "Brazil", Brazil},
		{"United Kingdom 1", "GB", UnitedKingdom},
		{"United Kingdom 2", "United Kingdom", UnitedKingdom},
		{"France 1", "FR", France},
		{"France 2", "France", France},
		{"Germany 1", "DE", Germany},
		{"Germany 2", "Germany", Germany},
		{"Spain 1", "ES", Spain},
		{"Spain 2", "Spain", Spain},
		{"Italy 1", "IT", Italy},
		{"Italy 2", "Italy", Italy},
		{"Japan 1", "JP", Japan},
		{"Japan 2", "Japan", Japan},
		{"China 1", "CN", China},
		{"China 2", "China", China},
		{"India 1", "IN", India},
		{"India 2", "India", India},
		{"Australia 1", "AU", Australia},
		{"Australia 2", "Australia", Australia},
		{"South Korea 1", "KR", SouthKorea},
		{"South Korea 2", "South Korea", SouthKorea},
	}
	return cases
}
