package implementstr

import "testing"

func TestStrFunc(t *testing.T) {
	tt := []struct {
		name     string
		haystack string
		needle   string
		expected int
	}{
		{
			name:     "when needle is present in haystack",
			haystack: "hello",
			needle:   "ll",
			expected: 2,
		},
		{
			name:     "when needle is not present",
			haystack: "aaaaa",
			needle:   "bba",
			expected: -1,
		},
		{
			name:     "when the needle is empty",
			haystack: "aaaa",
			needle:   "",
			expected: 0,
		},
		{
			name:     "needle and haystack are same",
			haystack: "a",
			needle:   "a",
			expected: 0,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if out := strStr(tc.haystack, tc.needle); out != tc.expected {
				t.Fatalf("expected %v but got %v", tc.expected, out)
			}
		})
	}
}
