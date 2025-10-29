package shared

import (
	"errors"
	"fmt"
)

type Country string

const (
	Canada        Country = "CA"
	Mexico        Country = "MX"
	UnitedStates  Country = "US"
	Argentina     Country = "AR"
	Bolivia       Country = "BO"
	Brazil        Country = "BR"
	UnitedKingdom Country = "GB"
	France        Country = "FR"
	Germany       Country = "DE"
	Spain         Country = "ES"
	Italy         Country = "IT"
	Japan         Country = "JP"
	China         Country = "CN"
	India         Country = "IN"
	Australia     Country = "AU"
	SouthKorea    Country = "KR"
)

var ErrNotACountry = errors.New("is not a country")

func (country Country) String() string {
	switch country {
	case Canada:
		return "Canada"
	case Mexico:
		return "Mexico"
	case UnitedStates:
		return "United States"
	case Argentina:
		return "Argentina"
	case Bolivia:
		return "Bolivia"
	case Brazil:
		return "Brazil"
	case UnitedKingdom:
		return "United Kingdom"
	case France:
		return "France"
	case Germany:
		return "Germany"
	case Spain:
		return "Spain"
	case Italy:
		return "Italy"
	case Japan:
		return "Japan"
	case China:
		return "China"
	case India:
		return "India"
	case Australia:
		return "Australia"
	case SouthKorea:
		return "South Korea"
	default:
		return "Unknown"
	}
}

func ParseCountry(country string) (Country, error) {
	switch country {
	case "CA", "Canada":
		return Canada, nil
	case "MX", "Mexico":
		return Mexico, nil
	case "US", "United States":
		return UnitedStates, nil
	case "AR", "Argentina":
		return Argentina, nil
	case "BO", "Bolivia":
		return Bolivia, nil
	case "BR", "Brazil":
		return Brazil, nil
	case "GB", "United Kingdom":
		return UnitedKingdom, nil
	case "FR", "France":
		return France, nil
	case "DE", "Germany":
		return Germany, nil
	case "ES", "Spain":
		return Spain, nil
	case "IT", "Italy":
		return Italy, nil
	case "JP", "Japan":
		return Japan, nil
	case "CN", "China":
		return China, nil
	case "IN", "India":
		return India, nil
	case "AU", "Australia":
		return Australia, nil
	case "KR", "South Korea":
		return SouthKorea, nil
	default:
		return "", fmt.Errorf("%w: got %s", ErrNotACountry, country)
	}
}

func ListCountries() []Country {
	return []Country{
		Canada,
		Mexico,
		UnitedStates,
		Argentina,
		Bolivia,
		Brazil,
		UnitedKingdom,
		France,
		Germany,
		Spain,
		Italy,
		Japan,
		China,
		India,
		Australia,
		SouthKorea,
	}
}
