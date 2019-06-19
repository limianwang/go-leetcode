package lengthoflastword

import "testing"

func TestLengthOfLastWord(t *testing.T) {
	tt := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "when string has trailing spaces",
			input:    "Hello World       ",
			expected: 5,
		},
		{
			name:     "when string is empty",
			input:    "",
			expected: 0,
		},
		{
			name:     "when string is empty with spaces",
			input:    " ",
			expected: 0,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if out := lengthOfLastWord(tc.input); out != tc.expected {
				t.Fatalf("expected %v but got %v", tc.expected, out)
			}
		})
	}
}
