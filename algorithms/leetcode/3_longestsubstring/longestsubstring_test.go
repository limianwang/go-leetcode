package longestsubstring

import "testing"

func TestLengthOfSubstring(t *testing.T) {
	tt := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "when string is all duplicated",
			input:    "aaaaaaa",
			expected: 1,
		},

		{
			name:     "when we have no duplicate",
			input:    "abc",
			expected: 3,
		},
		{
			name:     "when the substring is in the middle",
			input:    "zzzabcddd",
			expected: 5,
		},
		{
			name:     "when the substring is in the middle 2",
			input:    "zzaabccdd",
			expected: 3,
		},
		{
			name:     "when the substring is what leetcode has",
			input:    "abcabcbb",
			expected: 3,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if out := lengthOfLongestSubstring(tc.input); out != tc.expected {
				t.Fatalf("expected %v but got %v", tc.expected, out)
			}
		})
	}
}
