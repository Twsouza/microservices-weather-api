package services

import (
	"context"
	"testing"
)

// TestIsValidCEP tests the IsValidCEP method.
func TestIsValidCEP(t *testing.T) {
	service := &ZipCodeValidator{}
	tests := []struct {
		cep     string
		isValid bool
	}{
		{"12345678", true},
		{"1234567", false},
		{"abcdefgh", false},
		{"123456789", false},
	}

	for _, test := range tests {
		result := service.IsValidCEP(context.Background(), test.cep)
		if result != test.isValid {
			t.Errorf("IsValidCEP(%s) = %v; want %v", test.cep, result, test.isValid)
		}
	}
}
