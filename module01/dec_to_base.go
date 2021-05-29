package module01

import "fmt"

// DecToBase will return a string representing
// the provided decimal number in the provided base.
// This is limited to bases 2-16 for simplicity.
//
// Eg:
//
//   DecToBase(14, 16) => "E"
//   DecToBase(14, 2) => "1110"
//
func DecToBase(dec, base int) string {
	if base > 16 {
		panic("base not supported")
	}

	result := ""
	for i := dec; i != 0; {
		if base > 10 {
			result = DecToHex(i%base) + result
		} else {
			result = fmt.Sprintf("%d", i%base) + result
		}

		i = i / base
	}

	return result
}

func DecToHex(dec int) string {
	if dec < 10 && dec >= 0 {
		return fmt.Sprintf("%d", dec)
	}

	switch dec {
	case 10:
		return "A"
	case 11:
		return "B"
	case 12:
		return "C"
	case 13:
		return "D"
	case 14:
		return "E"
	case 15:
		return "F"
	default:
		panic("value not allowed")
	}
}
