package module01

import (
	"math"
	"strconv"
)

// BaseToDec takes in a number and the base it is currently
// in and returns the decimal equivalent as an integer.
//
// Eg:
//
//   BaseToDec("E", 16) => 14
//   BaseToDec("1110", 2) => 14
//
func BaseToDec(value string, base int) int {
	result := 0
	i := len(value) - 1
	for _, v := range value {
		digit := HexToDec(string(v))
		result += digit * int(math.Pow(float64(base), float64(i)))
		i--
	}
	return result
}

func HexToDec(hex string) int {
	if v, err := strconv.Atoi(hex); err == nil {
		return v
	}

	switch hex {
	case "A":
		return 10
	case "B":
		return 11
	case "C":
		return 12
	case "D":
		return 13
	case "E":
		return 14
	case "F":
		return 15
	default:
		panic("value not allowed")
	}
}
