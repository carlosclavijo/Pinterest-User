package shared

import (
	"errors"
	"fmt"
)

type Language string

const (
	English    Language = "en"
	Spanish    Language = "es"
	French     Language = "fr"
	German     Language = "de"
	Italian    Language = "it"
	Portuguese Language = "pt"
	Japanese   Language = "ja"
	Chinese    Language = "zh"
	Korean     Language = "ko"
	Hindi      Language = "hi"
)

var ErrNotALanguage = errors.New("is not a language")

func (l Language) String() string {
	switch l {
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
		return "unknown"
	}
}

func ParseLanguage(language string) (Language, error) {
	switch language {
	case "en", "English":
		return English, nil
	case "es", "Spanish":
		return Spanish, nil
	case "fr", "French":
		return French, nil
	case "de", "German":
		return German, nil
	case "it", "Italian":
		return Italian, nil
	case "pt", "Portuguese":
		return Portuguese, nil
	case "ja", "Japanese":
		return Japanese, nil
	case "zh", "Chinese":
		return Chinese, nil
	case "ko", "Korean":
		return Korean, nil
	case "hi", "Hindi":
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
