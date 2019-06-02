package buysellstockmultiple

import "testing"

func TestMaxProfit(t *testing.T) {
	tt := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "when no input is given",
			input:    []int{},
			expected: 0,
		},
		{
			name:     "when there are no inflection point",
			input:    []int{1, 2, 5},
			expected: 4,
		},
		{
			name:     "when there are inflection point",
			input:    []int{1, 3, 2, 3},
			expected: 3,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if profit := maxProfit(tc.input); profit != tc.expected {
				t.Fatalf("expected: %d but got: %d", tc.expected, profit)
			}
		})
	}
}
