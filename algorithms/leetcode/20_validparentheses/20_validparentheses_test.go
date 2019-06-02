package validparentheses

import "testing"

func TestIsValid(t *testing.T) {
	tt := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "when the string is empty",
			input:    "",
			expected: true,
		},
		{
			name:     "when string does not include parenthese",
			input:    "abc",
			expected: true,
		},
		{
			name:     "when string does not match",
			input:    "{testtest]",
			expected: false,
		},
		{
			name:     "when string has long but balanced parentheses",
			input:    "{}{}{}[][]()",
			expected: true,
		},
		{
			name:     "when string has matched",
			input:    "{{{{{[()]}}}}}",
			expected: true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if out := isValid(tc.input); out != tc.expected {
				t.Fatalf("expected %v but got %v with input %s", tc.expected, out, tc.input)
			}
		})
	}
}
