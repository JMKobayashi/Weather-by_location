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
	}{
		{
			name:        "invalid zipcode format",
			zipcode:     "1234567",
			expectError: true,
		},
		{
			name:        "non-existent zipcode",
			zipcode:     "99999999",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := service.GetWeatherByZipcode(tt.zipcode)
			if (err != nil) != tt.expectError {
				t.Errorf("GetWeatherByZipcode(%s) error = %v; expectError %v", tt.zipcode, err, tt.expectError)
			}
		})
	}
}
