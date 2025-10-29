package shared

import (
	"errors"
	"fmt"
)

type DialCode string

const (
	UnitedStatesDc  DialCode = "+1"
	MexicoDC        DialCode = "+52"
	ArgentinaDC     DialCode = "+54"
	BoliviaDC       DialCode = "+591"
	BrazilDC        DialCode = "+55"
	UnitedKingdomDC DialCode = "+44"
	FranceDC        DialCode = "+33"
	GermanyDC       DialCode = "+49"
	SpainDC         DialCode = "+34"
	ItalyDC         DialCode = "+39"
	JapanDC         DialCode = "+81"
	ChinaDC         DialCode = "+86"
	IndiaDC         DialCode = "+91"
	AustraliaDC     DialCode = "+61"
	SouthKoreaDC    DialCode = "+82"
)

var ErrInvalidDial = errors.New("is not a valid dial code")

func (dialcode DialCode) String() string {
	switch dialcode {
	case UnitedStatesDc:
		return "United States"
	case MexicoDC:
		return "Mexico"
	case ArgentinaDC:
		return "Argentina"
	case BoliviaDC:
		return "Bolivia"
	case BrazilDC:
		return "Brazil"
	case UnitedKingdomDC:
		return "United Kingdom"
	case FranceDC:
		return "France"
	case GermanyDC:
		return "Germany"
	case SpainDC:
		return "Spain"
	case ItalyDC:
		return "Italy"
	case JapanDC:
		return "Japan"
	case ChinaDC:
		return "China"
	case IndiaDC:
		return "India"
	case AustraliaDC:
		return "Australia"
	case SouthKoreaDC:
		return "South Korea"
	default:
		return "Unknown"
	}
}

func ParseDialCode(dial string) (DialCode, error) {
	switch dial {
	case "+1":
		return UnitedStatesDc, nil
	case "+52":
		return MexicoDC, nil
	case "+54":
		return ArgentinaDC, nil
	case "+591":
		return BoliviaDC, nil
	case "+55":
		return BrazilDC, nil
	case "+44":
		return UnitedKingdomDC, nil
	case "+33":
		return FranceDC, nil
	case "+49":
		return GermanyDC, nil
	case "+34":
		return SpainDC, nil
	case "+39":
		return ItalyDC, nil
	case "+81":
		return JapanDC, nil
	case "+86":
		return ChinaDC, nil
	case "+91":
		return IndiaDC, nil
	case "+61":
		return AustraliaDC, nil
	case "+82":
		return SouthKoreaDC, nil
	default:
		return "", fmt.Errorf("%w: got %s", ErrInvalidDial, dial)
	}
}

func ListDialCodes() []DialCode {
	return []DialCode{
		UnitedStatesDc,
		MexicoDC,
		ArgentinaDC,
		BoliviaDC,
		BrazilDC,
		UnitedKingdomDC,
		FranceDC,
		GermanyDC,
		SpainDC,
		ItalyDC,
		JapanDC,
		ChinaDC,
		IndiaDC,
		AustraliaDC,
		SouthKoreaDC,
	}
}
