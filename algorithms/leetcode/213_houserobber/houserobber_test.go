package houserobber

import "testing"

func TestRob(t *testing.T) {
	tt := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "when input is odd",
			input:    []int{2, 3, 2},
			expected: 3,
		},
		{
			name:     "when input is even",
			input:    []int{1, 2, 3, 1},
			expected: 4,
		},
		{
			name:     "when max is first",
			input:    []int{2, 1},
			expected: 2,
		},
		{
			name:     "when max is in middle",
			input:    []int{1, 2, 1, 1},
			expected: 3,
		},
		{
			name:     "when input is empty",
			input:    []int{},
			expected: 0,
		},

		{
			name:     "when input is single",
			input:    []int{1},
			expected: 1,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if out := rob(tc.input); out != tc.expected {
				t.Fatalf("expected %d but got %d", tc.expected, out)
			}
		})
	}
}
