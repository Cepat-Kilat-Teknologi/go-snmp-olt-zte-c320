package utils

import (
	"encoding/binary"
	"errors"
	"fmt"
	"strconv"
	"time"
)

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

// ConvertDurationToString Convert duration to human-readable format
func ConvertDurationToString(duration time.Duration) string {
	days := int(duration / (24 * time.Hour))
	duration = duration % (24 * time.Hour)
	hours := int(duration / time.Hour)
	duration = duration % time.Hour
	minutes := int(duration / time.Minute)
	duration = duration % time.Minute
	seconds := int(duration / time.Second)

	return strconv.Itoa(days) + " days " + strconv.Itoa(hours) + " hours " + strconv.Itoa(minutes) + " minutes " + strconv.Itoa(seconds) + " seconds"
}

// ConvertByteArrayToDateTime Convert byte array to human-readable date time
func ConvertByteArrayToDateTime(byteArray []byte) (string, error) {

	// Check if byteArray length is exactly 8
	if len(byteArray) != 8 {
		return "", errors.New("invalid byte array length: expected 8 bytes")
	}

	// Extract the year from the first two bytes
	year := int(binary.BigEndian.Uint16(byteArray[0:2]))

	// Extract other components
	month := time.Month(byteArray[2]) // Month
	day := int(byteArray[3])          // Day
	hour := int(byteArray[4])         // Hour
	minute := int(byteArray[5])       // Minute
	second := int(byteArray[6])       // Second

	// Validate extracted values
	if month < 1 || month > 12 {
		return "", fmt.Errorf("invalid month: %d", month)
	}
	if day < 1 || day > 31 {
		return "", fmt.Errorf("invalid day: %d", day)
	}
	if hour < 0 || hour > 23 {
		return "", fmt.Errorf("invalid hour: %d", hour)
	}
	if minute < 0 || minute > 59 {
		return "", fmt.Errorf("invalid minute: %d", minute)
	}
	if second < 0 || second > 59 {
		return "", fmt.Errorf("invalid second: %d", second)
	}

	// Create a time.Time object UTC
	datetime := time.Date(year, month, day, hour, minute, second, 0, time.UTC)

	// Convert to Unix epoch time (seconds since Jan 1, 1970)
	return datetime.Format("2006-01-02 15:04:05"), nil
}
