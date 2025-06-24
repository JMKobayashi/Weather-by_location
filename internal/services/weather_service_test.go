package services

import (
	"testing"
)

func TestIsValidZipcode(t *testing.T) {
	tests := []struct {
		name     string
		zipcode  string
		expected bool
	}{
		{
			name:     "valid zipcode",
			zipcode:  "12345678",
			expected: true,
		},
		{
			name:     "valid zipcode with hyphens",
			zipcode:  "12345-678",
			expected: true,
		},
		{
			name:     "valid zipcode with spaces",
			zipcode:  "12345 678",
			expected: true,
		},
		{
			name:     "invalid zipcode with letters",
			zipcode:  "1234567a",
			expected: false,
		},
		{
			name:     "invalid zipcode with less digits",
			zipcode:  "1234567",
			expected: false,
		},
		{
			name:     "invalid zipcode with more digits",
			zipcode:  "123456789",
			expected: false,
		},
		{
			name:     "empty zipcode",
			zipcode:  "",
			expected: false,
		},
		{
			name:     "zipcode with special characters",
			zipcode:  "12345@678",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isValidZipcode(tt.zipcode)
			if result != tt.expected {
				t.Errorf("isValidZipcode(%s) = %v; want %v", tt.zipcode, result, tt.expected)
			}
		})
	}
}

func TestGetWeatherByZipcode(t *testing.T) {
	service := NewWeatherService("test-key")

	tests := []struct {
		name        string
		zipcode     string
		expectError bool
		errorMsg    string
	}{
		{
			name:        "invalid zipcode format",
			zipcode:     "1234567",
			expectError: true,
			errorMsg:    "invalid zipcode",
		},
		{
			name:        "invalid zipcode with letters",
			zipcode:     "1234567a",
			expectError: true,
			errorMsg:    "invalid zipcode",
		},
		{
			name:        "empty zipcode",
			zipcode:     "",
			expectError: true,
			errorMsg:    "invalid zipcode",
		},
		{
			name:        "non-existent zipcode",
			zipcode:     "99999999",
			expectError: true,
			errorMsg:    "can not find zipcode",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := service.GetWeatherByZipcode(tt.zipcode)
			if (err != nil) != tt.expectError {
				t.Errorf("GetWeatherByZipcode(%s) error = %v; expectError %v", tt.zipcode, err, tt.expectError)
				return
			}

			if tt.expectError && err != nil && err.Error() != tt.errorMsg {
				t.Errorf("GetWeatherByZipcode(%s) error message = %v; want %v", tt.zipcode, err.Error(), tt.errorMsg)
			}
		})
	}
}
