package utils

import "strconv"

// ConvertStringToUint16 Convert String to Uint16
func ConvertStringToUint16(str string) uint16 {
	// Convert string to uint16
	value, err := strconv.ParseUint(str, 10, 16)
	if err != nil {
		return 0
	}

	return uint16(value)
}

// ConvertStringToInteger Convert String to Integer
func ConvertStringToInteger(str string) int {
	// Convert string to int
	value, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}

	return value
}
