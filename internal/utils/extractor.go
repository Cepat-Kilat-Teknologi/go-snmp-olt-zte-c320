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
		// Check if the last component is a valid number
		lastComponent := parts[len(parts)-1]
		if _, err := strconv.Atoi(lastComponent); err == nil {
			return lastComponent
		}
	}
	return "" // Return an empty string if the OID is invalid or empty (default value)
}

func ExtractIDOnuID(oid interface{}) int {
	if oid == nil {
		return 0
	}

	switch v := oid.(type) {
	case string:
		parts := strings.Split(v, ".")
		if len(parts) > 0 {
			lastPart := parts[len(parts)-1]
			id, err := strconv.Atoi(lastPart)
			if err == nil {
				return id
			}
		}
		return 0
	default:
		return 0
	}
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
		return "Unknown" // Return "Unknown" if the OID is invalid or empty
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
		return "" // Return 0 if the OID is invalid or empty (default value)
	}
}

func ConvertAndMultiply(pduValue interface{}) (string, error) {
	// Type assert pduValue to an integer type
	intValue, ok := pduValue.(int)
	if !ok {
		return "", fmt.Errorf("value is not an integer")
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
	// Check if oidValue is not an integer
	intValue, ok := oidValue.(int)
	if !ok {
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

func ExtractLastOfflineReason(oidValue interface{}) string {
	// Check if oidValue is not an integer
	intValue, ok := oidValue.(int)
	if !ok {
		return "Unknown"
	}

	switch intValue {
	case 1:
		return "Unknown"
	case 2:
		return "LOS"
	case 3:
		return "LOSi"
	case 4:
		return "LOFi"
	case 5:
		return "sfi"
	case 6:
		return "loai"
	case 7:
		return "loami"
	case 8:
		return "AuthFail"
	case 9:
		return "PowerOff"
	case 10:
		return "deactiveSucc"
	case 11:
		return "deactiveFail"
	case 12:
		return "Reboot"
	case 13:
		return "Shutdown"
	default:
		return "Unknown"
	}
}

func ExtractGponOpticalDistance(oidValue interface{}) string {
	// Check if oidValue is not an integer
	intValue, ok := oidValue.(int)
	if !ok {
		return "Unknown"
	}

	return strconv.Itoa(intValue)
}
