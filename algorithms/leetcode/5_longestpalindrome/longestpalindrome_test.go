package longestpalindrome

import "testing"

func TestLongestPalindrome(t *testing.T) {
	tt := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "when empty string gets passed",
			input:    "",
			expected: "",
		},
		{
			name:     "when input is abacd",
			input:    "abacd",
			expected: "aba",
		},
		{
			name:     "when input is two char that is palindrome",
			input:    "aa",
			expected: "aa",
		},
		{
			name:     "when input is two char that is not palindrome",
			input:    "ab",
			expected: "a",
		},
		{
			name:     "when input is two char that has palindrome at end",
			input:    "abdd",
			expected: "dd",
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
			name:     "when there is all same characters",
			input:    "aaa",
			expected: "aaa",
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
