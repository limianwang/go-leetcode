package findmin

import "testing"

func TestFindMin(t *testing.T) {
	tt := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "when array is entirely sorted",
			input:    []int{1, 2, 3, 4, 5},
			expected: 1,
		},
		{
			name:     "when array is rotated",
			input:    []int{4, 5, 6, 7, 0, 1, 2},
			expected: 0,
		},
		{
			name:     "when array is empty",
			input:    []int{},
			expected: -1,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if out := findMin(tc.input); out != tc.expected {
				t.Fatalf("expected %v but got %v", tc.expected, out)
			}
		})
	}
}
