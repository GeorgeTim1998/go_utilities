package main

import (
	"testing"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		hasError bool
	}{
		{"a4bc2d5e", "aaaabccddddde", false},
		{"abcd", "abcd", false},
		{"45", "", true},
		{"", "", false},
		{"qwe\\4\\5", "qwe45", false},
		{"qwe\\45", "qwe44444", false},
		{"qwe\\\\5", "qwe\\\\\\\\\\", false},
	}

	for _, test := range tests {
		result, _, err := Unpack(test.input)
		if (err != nil) != test.hasError {
			t.Errorf("Unpack(%q) returned error %v, expected error: %v", test.input, err, test.hasError)
		}
		if result != test.expected {
			t.Errorf("Unpack(%q) = %q, want %q", test.input, result, test.expected)
		}
	}
}
