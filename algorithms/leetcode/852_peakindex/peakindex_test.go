package peakindex

import "testing"

func TestPeak(t *testing.T) {
	tt := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "when input has a peak",
			input:    []int{0, 1, 0},
			expected: 1,
		},
		{
			name:     "when input has a peak sequence",
			input:    []int{0, 2, 1, 0},
			expected: 1,
		},
		{
			name:     "when input has no peak",
			input:    []int{0, 0, 0},
			expected: -1,
		},
		{
			name:     "when input has peak at end",
			input:    []int{0, 1, 2, 7, 3},
			expected: 3,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if out := peakIndexInMountainArray(tc.input); out != tc.expected {
				t.Fatalf("expected %d but got %d", tc.expected, out)
			}
		})
	}
}
