package buysellstock

import "testing"

func TestBuySell(t *testing.T) {
	tt := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "when the array is empty",
			input:    []int{},
			expected: 0,
		},
		{
			name:     "when there is no inflection point",
			input:    []int{1, 2, 3, 4, 5, 6},
			expected: 5,
		},
		{
			name:     "when there is no inflection point (negative)",
			input:    []int{6, 5, 4, 3, 2, 1},
			expected: 0,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if out := maxProfit(tc.input); out != tc.expected {
				t.Fatalf("expected %d but got %d", tc.expected, out)
			}
		})
	}
}
