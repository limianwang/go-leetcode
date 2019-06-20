package houserobber

import "testing"

func TestRob(t *testing.T) {
	tt := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "when there is no houses or money",
			input:    []int{},
			expected: 0,
		},
		{
			name:     "when there is only one house",
			input:    []int{5},
			expected: 5,
		},
		{
			name:     "when only two houses available",
			input:    []int{1, 2},
			expected: 2,
		},
		{
			name:     "when only two houses available",
			input:    []int{2, 1},
			expected: 2,
		},
		{
			name:     "max allowable when the array is even",
			input:    []int{1, 2, 3, 1},
			expected: 4,
		},
		{
			name:     "max allowable when the array is even",
			input:    []int{2, 7, 9, 3, 1},
			expected: 12,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if out := rob(tc.input); out != tc.expected {
				t.Fatalf("expected %v but got %v", tc.expected, out)
			}
		})
	}
}
