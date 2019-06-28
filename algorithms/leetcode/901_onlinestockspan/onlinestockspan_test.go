package onlinestockspan

import "testing"

func TestStockSpanner(t *testing.T) {
	tt := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "valid output",
			input:    []int{100, 80, 60, 70, 60, 75, 85},
			expected: []int{1, 1, 1, 2, 1, 4, 6},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			ss := Constructor()
			for idx, input := range tc.input {
				if output := ss.Next(input); output != tc.expected[idx] {
					t.Fatalf("expected %d but got %d", tc.expected[idx], output)
				}
			}
		})
	}
}
