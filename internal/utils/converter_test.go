package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
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
