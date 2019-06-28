package findpeak

import "testing"

func TestFindPeak(t *testing.T) {
	tt := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "when arr is not empty",
			input:    []int{1, 2, 3, 1},
			expected: 2,
		},
		{
			name:     "when array is empty",
			input:    []int{},
			expected: 0,
		},
		{
			name:     "when array is weird",
			input:    []int{7, 1, 2, 3, 4, 6},
			expected: 5,
		},
		{
			name:     "when array is what leetcode provided",
			input:    []int{1, 2, 1, 3, 5, 6, 4},
			expected: 5,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if out := findPeakElement(tc.input); out != tc.expected {
				t.Fatalf("expected %d but got %d", tc.expected, out)
			}
		})
	}
}
