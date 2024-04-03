package stringutil

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "olleh"},
		{"", ""},
		{"12345", "54321"},
	}

	for _, test := range tests {
		result := Reverse(test.input)
		if result != test.expected {
			t.Errorf("Reverse(%q) expected %q, but got %q", test.input, test.expected, result)
		}
	}
}

func TestRunesCount(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"hello", 5},
		{"", 0},
		{"12345", 5},
	}

	for _, test := range tests {
		result := RunesCount(test.input)
		if result != test.expected {
			t.Errorf("RunesCount(%q) expected %d, but got %d", test.input, test.expected, result)
		}
	}
}
