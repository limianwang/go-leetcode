package longestpalindrome

import "testing"

func TestLongestPalindrome(t *testing.T) {
	tt := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "when input is abacd",
			input:    "abacd",
			expected: "aba",
		},
		{
			name:     "when input is a string char",
			input:    "a",
			expected: "a",
		},
		{
			name:     "when the input is greater than 3",
			input:    "banana",
			expected: "anana",
		},
		{
			name:     "when there is only single character palindrome",
			input:    "abcd",
			expected: "a",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if out := longestPalindrome(tc.input); out != tc.expected {
				t.Fatalf("expected: %v but got %v", tc.expected, out)
			}
		})
	}
}
