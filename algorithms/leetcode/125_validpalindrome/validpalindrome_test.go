package Validpalindrome

import "testing"

func TestIsValidpalindrome(t *testing.T) {
	tt := []struct {
		name   string
		input  string
		output bool
	}{
		{
			name:   "test valid",
			input:  "A man, a plan, a canal: Panama",
			output: true,
		},
		{
			name:   "test invalid",
			input:  "race a car",
			output: false,
		},
		{
			name:   "test should be valid",
			input:  "0P",
			output: false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if isPalindrome(tc.input) != tc.output {
				t.Fatalf("expected to be %v", tc.output)
			}
		})
	}
}
