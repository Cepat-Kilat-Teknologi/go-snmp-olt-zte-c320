package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func ExtractONUID(oid string) string {
	// Menguraikan nama OID dan mengambil komponen terakhir
	parts := strings.Split(oid, ".")
	if len(parts) > 0 {
		return parts[len(parts)-1]
	}
	return ""
}

func ExtractIDOnuID(oid interface{}) int {
	// Menguraikan nama OID dan mengambil komponen terakhir
	parts := strings.Split(oid.(string), ".")
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
		// Data adalah string, gunakan langsung
		return v
	case []byte:
		// Data adalah byte slice, konversi ke string
		return string(v)
	default:
		// Tipe data tidak dikenali, Anda dapat menghandle kasus ini sesuai kebutuhan Anda.
		return ""
	}
}

// ExtractSerialNumber Fungsi ini mengambil Serial Number dari hasil SNMP Walk
func ExtractSerialNumber(oidValue interface{}) string {
	switch v := oidValue.(type) {
	case string:
		// Data adalah string, gunakan langsung
		if strings.HasPrefix(v, "1,") {
			return v[2:]
		}
		return v
	case []byte:
		// Data adalah byte slice, konversi ke string
		strValue := string(v)
		if strings.HasPrefix(strValue, "1,") {
			return strValue[2:]
		}
		return strValue
	default:
		// Tipe data tidak dikenali, Anda dapat menghandle kasus ini sesuai kebutuhan Anda.
		return ""
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

	// Convert the result to a string
	resultStr := strconv.FormatFloat(result, 'f', -1, 64)

	return resultStr, nil
}

func ExtractAndGetStatus(oidValue interface{}) string {
	// Extract the interface value to an integer type
	intValue, err := strconv.Atoi(strconv.Itoa(oidValue.(int)))
	if err != nil {
		// Handle error
		return ""
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

	// Return the integer value
	return strconv.Itoa(intValue)
}
