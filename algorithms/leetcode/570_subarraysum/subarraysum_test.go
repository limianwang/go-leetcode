package subarray

import "testing"

func TestSubarraySum(t *testing.T) {
	tt := []struct {
		name     string
		input    []int
		k        int
		expected int
	}{
		{
			name:     "input from leetcode",
			input:    []int{1, 1, 1},
			k:        2,
			expected: 2,
		},
		{
			name:     "input when it should still be valid",
			input:    []int{1, 2, 3},
			k:        6,
			expected: 1,
		},
		{
			name:     "customized input",
			input:    []int{1, 1, 1, 2},
			k:        2,
			expected: 3,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if out := subarraySum(tc.input, tc.k); out != tc.expected {
				t.Fatalf("expected %v but got %v", tc.expected, out)
			}
		})
	}
}
