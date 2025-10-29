package shared

import (
	"errors"
	"fmt"
)

type Language string

const (
	English    Language = "EN"
	Spanish    Language = "ES"
	French     Language = "FR"
	German     Language = "DE"
	Italian    Language = "IT"
	Portuguese Language = "PT"
	Japanese   Language = "JA"
	Chinese    Language = "CH"
	Korean     Language = "KO"
	Hindi      Language = "HI"
)

var ErrNotALanguage = errors.New("is not a language")

func (language Language) String() string {
	switch language {
	case English:
		return "English"
	case Spanish:
		return "Spanish"
	case French:
		return "French"
	case German:
		return "German"
	case Italian:
		return "Italian"
	case Portuguese:
		return "Portuguese"
	case Japanese:
		return "Japanese"
	case Chinese:
		return "Chinese"
	case Korean:
		return "Korean"
	case Hindi:
		return "Hindi"
	default:
		return "Unknown"
	}
}

func ParseLanguage(language string) (Language, error) {
	switch language {
	case "EN", "English":
		return English, nil
	case "ES", "Spanish":
		return Spanish, nil
	case "FR", "French":
		return French, nil
	case "DE", "German":
		return German, nil
	case "IT", "Italian":
		return Italian, nil
	case "PT", "Portuguese":
		return Portuguese, nil
	case "JA", "Japanese":
		return Japanese, nil
	case "CH", "Chinese":
		return Chinese, nil
	case "KO", "Korean":
		return Korean, nil
	case "HI", "Hindi":
		return Hindi, nil
	default:
		return "", fmt.Errorf("%w: got %s", ErrNotALanguage, language)
	}
}

func ListLanguages() []Language {
	return []Language{
		English,
		Spanish,
		French,
		German,
		Italian,
		Portuguese,
		Japanese,
		Chinese,
		Korean,
		Hindi,
	}
}
