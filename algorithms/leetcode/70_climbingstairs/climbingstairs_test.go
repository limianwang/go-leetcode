package climbingstairs

import "testing"

func TestClimbStairs(t *testing.T) {
	tt := []struct {
		name     string
		input    int
		expected int
	}{
		{
			name:     "when input is 1",
			input:    1,
			expected: 1,
		},
		{
			name:     "when input is 2",
			input:    2,
			expected: 2,
		},
		{
			name:     "when input is 3",
			input:    3,
			expected: 3,
		},
		{
			name:     "when n = 5 should return 8",
			input:    5,
			expected: 8,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if out := climbStairs(tc.input); out != tc.expected {
				t.Fatalf("expected %v but got %v", tc.expected, out)
			}
		})
	}
}
