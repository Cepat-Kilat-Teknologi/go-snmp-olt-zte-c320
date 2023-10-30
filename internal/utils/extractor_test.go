package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractONUID(t *testing.T) {
	testCases := []struct {
		oid      string
		expected string
	}{
		// Test with valid OID values
		{"1.2.3.4.5", "5"},
		{"1.2.3", "3"},
		{"1", "1"},
		{"", ""},

		// Test with invalid OID values
		{"invalid.oid", ""}, // Add test case for an invalid OID
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("OID: %v", tc.oid), func(t *testing.T) {
			result := ExtractONUID(tc.oid)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestExtractIDOnuID(t *testing.T) {
	testCases := []struct {
		oid      interface{}
		expected int
	}{
		{"1.2.3.4.5", 5},
		{"1.2.3", 3},
		{"1", 1},
		{nil, 0},
		{123, 0},
		{"", 0},
		{"1.2.3.4.invalid", 0},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("OID: %v", tc.oid), func(t *testing.T) {
			result := ExtractIDOnuID(tc.oid)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestExtractName(t *testing.T) {
	testCases := []struct {
		oidValue interface{}
		expected string
		testName string
	}{
		{"test", "test", "string"},
		{[]byte("test"), "test", "byte slice"},
		{10, "Unknown", "Unknown"},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := ExtractName(tc.oidValue)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestExtractSerialNumber(t *testing.T) {
	testCases := []struct {
		oidValue interface{}
		expected string
	}{
		{"1,SerialNumber", "SerialNumber"},
		{"SerialNumber", "SerialNumber"},
		{[]byte("1,SerialNumber"), "SerialNumber"},
		{[]byte("SerialNumber"), "SerialNumber"},
		{10, ""},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("OIDValue: %v", tc.oidValue), func(t *testing.T) {
			result := ExtractSerialNumber(tc.oidValue)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestConvertAndMultiply(t *testing.T) {
	testCases := []struct {
		pduValue interface{}
		expected string
		err      bool
	}{
		{10, "-29.98", false},
		{0, "-30.00", false},
		{"string", "Unknown", true},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("PDUValue: %v", tc.pduValue), func(t *testing.T) {
			result, err := ConvertAndMultiply(tc.pduValue)
			if tc.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, result)
			}
		})
	}
}

func TestExtractAndGetStatus(t *testing.T) {
	testCases := []struct {
		oidValue interface{}
		expected string
	}{
		// Test with valid integer values
		{1, "Logging"},
		{2, "LOS"},
		{3, "Synchronization"},
		{4, "Online"},
		{5, "Dying Gasp"},
		{6, "Auth Failed"},
		{7, "Offline"},

		// Test with invalid integer input
		{"invalid", "Unknown"},
		{8, "Unknown"}, // Add test case for a value not covered in the switch
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("OIDValue: %v", tc.oidValue), func(t *testing.T) {
			result := ExtractAndGetStatus(tc.oidValue)
			assert.Equal(t, tc.expected, result)
		})
	}
}
