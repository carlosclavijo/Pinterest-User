package shared

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

type dialCodeCase struct {
	name        string
	dialCodeStr string
	dialCode    DialCode
}

func TestParseDialCode(t *testing.T) {
	cases := dialCodesCases()
	var dialCodes []DialCode

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			dc, err := ParseDialCode(tc.dialCodeStr)
			dialCodes = append(dialCodes, dc)

			require.NotEqual(t, "", dc)
			require.NoError(t, err)
			assert.Equal(t, tc.dialCode, dc)
		})
	}

	assert.Len(t, cases, len(dialCodes))

	for v, dc := range dialCodes {
		t.Run(cases[v].name, func(t *testing.T) {
			dcStr := dc.String()
			assert.Equal(t, cases[v].name, dcStr)
		})
	}
}

func TestParseDialCode_Invalid(t *testing.T) {
	str := "X"

	dc, err := ParseDialCode(str)

	require.Empty(t, dc)
	require.ErrorIs(t, err, ErrInvalidDial)

	dcStr := dc.String()
	assert.Equal(t, "Unknown", dcStr)

}

func TestListDialCodes(t *testing.T) {
	dialCodes := ListDialCodes()

	require.NotEmpty(t, dialCodes)
	assert.Contains(t, dialCodes, UnitedStatesDc)
	assert.Contains(t, dialCodes, MexicoDC)
	assert.Contains(t, dialCodes, ArgentinaDC)
	assert.Contains(t, dialCodes, BoliviaDC)
	assert.Contains(t, dialCodes, BrazilDC)
	assert.Contains(t, dialCodes, UnitedKingdomDC)
	assert.Contains(t, dialCodes, FranceDC)
	assert.Contains(t, dialCodes, GermanyDC)
	assert.Contains(t, dialCodes, SpainDC)
	assert.Contains(t, dialCodes, ItalyDC)
	assert.Contains(t, dialCodes, JapanDC)
	assert.Contains(t, dialCodes, ChinaDC)
	assert.Contains(t, dialCodes, IndiaDC)
	assert.Contains(t, dialCodes, AustraliaDC)
	assert.Contains(t, dialCodes, SouthKoreaDC)
}

func dialCodesCases() []dialCodeCase {
	cases := []dialCodeCase{
		{"United States", "+1", UnitedStatesDc},
		{"Mexico", "+52", MexicoDC},
		{"Argentina", "+54", ArgentinaDC},
		{"Bolivia", "+591", BoliviaDC},
		{"Brazil", "+55", BrazilDC},
		{"United Kingdom", "+44", UnitedKingdomDC},
		{"France", "+33", FranceDC},
		{"Germany", "+49", GermanyDC},
		{"Spain", "+34", SpainDC},
		{"Italy", "+39", ItalyDC},
		{"Japan", "+81", JapanDC},
		{"China", "+86", ChinaDC},
		{"India", "+91", IndiaDC},
		{"Australia", "+61", AustraliaDC},
		{"South Korea", "+82", SouthKoreaDC},
	}
	return cases
}
