package validparentheses

import "testing"

func TestIsValid(t *testing.T) {
	tt := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "success",
			input:    "(abc{dd})",
			expected: true,
		},
		{
			name:     "failure",
			input:    "(ddd})",
			expected: false,
		},
		{
			name:     "failure1",
			input:    "]",
			expected: false,
		},
		{
			name:     "failure2",
			input:    "[",
			expected: false,
		},
		{
			name:     "alsosuccess1",
			input:    "{}{}{}",
			expected: true,
		},
		{
			name:     "alsosuccess2",
			input:    "{}()[]",
			expected: true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if result := isValid(tc.input); result != tc.expected {
				t.Fatalf("expected %v but got %v for %v", tc.expected, result, tc.input)
			}
		})

		t.Run(tc.name, func(t *testing.T) {
			if result := isBalanced(tc.input); result != tc.expected {
				t.Fatalf("expected %v but got %v for %v", tc.expected, result, tc.input)
			}
		})
	}
}
