package mincost

import "testing"

func TestMinCost(t *testing.T) {
	tt := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "input cost",
			input:    []int{1, 3, 4, 1, 2},
			expected: 4,
		},
		{
			name:     "when input is what expected",
			input:    []int{10, 15, 20},
			expected: 15,
		},
		{
			name:     "when input is what is in leetcode",
			input:    []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
			expected: 6,
		},
	}

	for _, tc := range tt {
		if out := minCostClimbingStairs(tc.input); out != tc.expected {
			t.Fatalf("expected %v, but got %v", tc.expected, out)
		}
	}
}
