package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func ExtractONUID(oid string) string {
	// Split the OID name and take the last component
	parts := strings.Split(oid, ".")
	if len(parts) > 0 {
		return parts[len(parts)-1]
	}
	return "0" // Return 0 if the OID is invalid or empty (default value)
}

func ExtractIDOnuID(oid interface{}) int {
	// Check if oid is nil
	if oid == nil {
		return 0
	}

	// Check if oid is a string
	oidStr, ok := oid.(string)
	if !ok {
		return 0
	}

	// Split the OID name and take the last component
	parts := strings.Split(oidStr, ".")
	if len(parts) > 0 {
		id, err := strconv.Atoi(parts[len(parts)-1])
		if err != nil {
			return 0
		}
		return id
	}
	return 0
}

func ExtractName(oidValue interface{}) string {
	switch v := oidValue.(type) {
	case string:
		// Data is string, return it
		return v
	case []byte:
		// Data is byte slice, convert to string
		return string(v)
	default:
		// Data type is not recognized, you can handle this case according to your needs.
		return "0" // Return 0 if the OID is invalid or empty (default value)
	}
}

// ExtractSerialNumber function is used to extract serial number from OID value
func ExtractSerialNumber(oidValue interface{}) string {
	switch v := oidValue.(type) {
	case string:
		// If the string starts with "1,", remove it from the string
		if strings.HasPrefix(v, "1,") {
			return v[2:]
		}
		return v
	case []byte:
		// Convert byte slice to string
		strValue := string(v)
		if strings.HasPrefix(strValue, "1,") {
			return strValue[2:]
		}
		return strValue // Data is byte slice, convert to string
	default:
		// Data type is not recognized, you can handle this case according to your needs.
		return "0" // Return 0 if the OID is invalid or empty (default value)
	}
}

func ConvertAndMultiply(pduValue interface{}) (string, error) {
	// Type assert pduValue to an integer type
	intValue, ok := pduValue.(int)
	if !ok {
		return "0", fmt.Errorf("value is not an integer")
	}

	// Multiply the integer by 0.002
	result := float64(intValue) * 0.002

	// Subtract 30
	result -= 30.0

	// Convert the result to a string with two decimal places
	resultStr := strconv.FormatFloat(result, 'f', 2, 64)

	return resultStr, nil
}

func ExtractAndGetStatus(oidValue interface{}) string {
	// Extract the interface value to an integer type
	intValue, err := strconv.Atoi(strconv.Itoa(oidValue.(int)))
	if err != nil {
		// Handle error
		return "Unknown"
	}

	switch intValue {
	case 1:
		return "Logging"
	case 2:
		return "LOS"
	case 3:
		return "Synchronization"
	case 4:
		return "Online"
	case 5:
		return "Dying Gasp"
	case 6:
		return "Auth Failed"
	case 7:
		return "Offline"
	default:
		return "Unknown"
	}
}
