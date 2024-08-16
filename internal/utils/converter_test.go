package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestConvertStringToUint16(t *testing.T) {
	// Test cases for ConvertStringToUint16
	testCases := []struct {
		input    string
		expected uint16
	}{
		{"123", 123},
		{"65535", 65535},
		{"abc", 0}, // Invalid input, should return 0
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := ConvertStringToUint16(tc.input)
			assert.Equal(t, tc.expected, result, "Expected and actual values should be equal.")
		})
	}
}

func TestConvertStringToInteger(t *testing.T) {
	// Test cases for ConvertStringToInteger
	testCases := []struct {
		input    string
		expected int
	}{
		{"123", 123},
		{"-456", -456},
		{"abc", 0}, // Invalid input, should return 0
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := ConvertStringToInteger(tc.input)
			assert.Equal(t, tc.expected, result, "Expected and actual values should be equal.")
		})
	}
}

func TestConvertDurationToString(t *testing.T) {
	tests := []struct {
		name     string
		duration time.Duration
		expected string
	}{
		{
			name:     "Test for 0 duration",
			duration: 0,
			expected: "0 days 0 hours 0 minutes 0 seconds",
		},
		{
			name:     "Test for exact seconds",
			duration: 5 * time.Second,
			expected: "0 days 0 hours 0 minutes 5 seconds",
		},
		{
			name:     "Test for exact minutes",
			duration: 3 * time.Minute,
			expected: "0 days 0 hours 3 minutes 0 seconds",
		},
		{
			name:     "Test for hours and minutes",
			duration: 2*time.Hour + 15*time.Minute,
			expected: "0 days 2 hours 15 minutes 0 seconds",
		},
		{
			name:     "Test for days, hours, minutes, and seconds",
			duration: 1*24*time.Hour + 4*time.Hour + 23*time.Minute + 45*time.Second,
			expected: "1 days 4 hours 23 minutes 45 seconds",
		},
		{
			name:     "Test for multiple days",
			duration: 3*24*time.Hour + 6*time.Hour + 30*time.Minute + 10*time.Second,
			expected: "3 days 6 hours 30 minutes 10 seconds",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ConvertDurationToString(tt.duration)
			if result != tt.expected {
				t.Errorf("ConvertDurationToString(%v) = %v; want %v", tt.duration, result, tt.expected)
			}
		})
	}
}

// TestConvertByteArrayToDateTime tests the ConvertByteArrayToDateTime function.
func TestConvertByteArrayToDateTime(t *testing.T) {
	tests := []struct {
		name          string
		byteArray     []byte
		expected      string
		expectedError bool
	}{
		{
			name: "Valid date and time",
			byteArray: []byte{
				0x07, 0xe4, 0x08, 0x15, 0x0a, 0x1e, 0x00, 0x00,
			}, // Year 2020, Month 8, Day 21, Hour 10, Minute 30, Second 00
			expected:      "2020-08-21 10:30:00",
			expectedError: false,
		},
		{
			name:          "Invalid month",
			byteArray:     []byte{0x07, 0xe4, 0x13, 0x15, 0x0a, 0x1e, 0x00, 0x00},
			expected:      "",
			expectedError: true,
		},
		{
			name:          "Invalid day",
			byteArray:     []byte{0x07, 0xe4, 0x08, 0x32, 0x0a, 0x1e, 0x00, 0x00},
			expected:      "",
			expectedError: true,
		},
		{
			name:          "Invalid hour",
			byteArray:     []byte{0x07, 0xe4, 0x08, 0x15, 0x18, 0x1e, 0x00, 0x00},
			expected:      "",
			expectedError: true,
		},
		{
			name:          "Invalid minute",
			byteArray:     []byte{0x07, 0xe4, 0x08, 0x15, 0x0a, 0x3c, 0x00, 0x00},
			expected:      "",
			expectedError: true,
		},
		{
			name:          "Invalid second",
			byteArray:     []byte{0x07, 0xe4, 0x08, 0x15, 0x0a, 0x1e, 0x3c, 0x00},
			expected:      "",
			expectedError: true,
		},
		{
			name:          "Invalid byte array length",
			byteArray:     []byte{0x07, 0xe4, 0x08, 0x15, 0x0a, 0x1e, 0x00},
			expected:      "",
			expectedError: true,
		},
		{
			name:          "Extra byte",
			byteArray:     []byte{0x07, 0xe4, 0x08, 0x15, 0x0a, 0x1e, 0x00, 0x00, 0x01},
			expected:      "",
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ConvertByteArrayToDateTime(tt.byteArray)
			if (err != nil) != tt.expectedError {
				t.Errorf("ConvertByteArrayToDateTime() error = %v, expectedError %v", err, tt.expectedError)
				return
			}
			if result != tt.expected {
				t.Errorf("ConvertByteArrayToDateTime() = %v, expected %v", result, tt.expected)
			}
		})
	}
}
