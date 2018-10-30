package firstpeak

import "testing"

func TestFirstPeak(t *testing.T) {
	tt := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "success1",
			input:    []int{1, 2, 1, 3, 5, 6, 4},
			expected: 1,
		},
		{
			name:     "success2",
			input:    []int{1, 2, 3, 1},
			expected: 2,
		},
		{
			name:     "success3",
			input:    []int{1, 2, 3, 4},
			expected: 3,
		},
		{
			name:     "success4",
			input:    []int{4, 3, 2, 1, 0},
			expected: 0,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if out := findPeakElement(tc.input); out != tc.expected {
				t.Fatalf("expected %v but got %v", tc.expected, out)
			}
		})
	}
}
