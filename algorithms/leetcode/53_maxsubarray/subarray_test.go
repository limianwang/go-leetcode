package maxsubarray

import "testing"

func TestMaxSubArray(t *testing.T) {
	tt := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "when there are negative and positive",
			input:    []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			expected: 6,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if res := maxSubArray(tc.input); res != tc.expected {
				t.Fatalf("expected %v but got %v", tc.expected, res)
			}
		})
	}
}
