package shared

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

type languageCase struct {
	name        string
	languageStr string
	language    Language
}

func TestParseLanguage(t *testing.T) {
	cases := languageCases()
	var languages []Language

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			l, err := ParseLanguage(tc.languageStr)
			languages = append(languages, l)

			require.NotEmpty(t, l)
			require.NoError(t, err)

			if len(tc.languageStr) != 2 {
				assert.Equal(t, tc.languageStr, l.String())
			}
		})
	}

	assert.Len(t, cases, len(languages))

	for v, l := range languages {
		t.Run(cases[v].name, func(t *testing.T) {
			lStr := l.String()

			if len(cases[v].languageStr) != 2 {
				assert.Equal(t, cases[v].languageStr, lStr)
			} else {
				assert.Equal(t, cases[v].language.String(), lStr)
			}
		})
	}
}

func TestParseLanguage_Invalid(t *testing.T) {
	str := "X"

	language, err := ParseLanguage(str)

	require.Empty(t, language)
	require.ErrorIs(t, err, ErrNotALanguage)

	languageStr := language.String()

	assert.Equal(t, "Unknown", languageStr)
}

func TestListLanguages(t *testing.T) {
	languages := ListLanguages()

	require.NotEmpty(t, languages)
	assert.Contains(t, languages, English)
	assert.Contains(t, languages, Spanish)
	assert.Contains(t, languages, French)
	assert.Contains(t, languages, German)
	assert.Contains(t, languages, Italian)
	assert.Contains(t, languages, Portuguese)
	assert.Contains(t, languages, Japanese)
	assert.Contains(t, languages, Chinese)
	assert.Contains(t, languages, Korean)
	assert.Contains(t, languages, Hindi)
}

func languageCases() []languageCase {
	cases := []languageCase{
		{"English 1", "EN", English},
		{"English 2", "English", English},
		{"Spanish 1", "ES", Spanish},
		{"Spanish 2", "Spanish", Spanish},
		{"French 1", "FR", French},
		{"English 2", "French", French},
		{"German 1", "DE", German},
		{"German 2", "German", German},
		{"Italian 1", "IT", Italian},
		{"Italian 2", "Italian", Italian},
		{"Portuguese 1", "PT", Portuguese},
		{"Portuguese 2", "Portuguese", Portuguese},
		{"Japanese 1", "JA", Japanese},
		{"Japanese 2", "Japanese", Japanese},
		{"Chinese 1", "CH", Chinese},
		{"Chinese 2", "Chinese", Chinese},
		{"Korean 1", "KO", Korean},
		{"Korean 2", "Korean", Korean},
		{"Hindi 1", "HI", Hindi},
		{"Hindi 2", "Hindi", Hindi},
	}

	return cases
}
