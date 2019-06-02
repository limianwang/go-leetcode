package firstmissingpositive

import "testing"

func TestFirstMissingPositive(t *testing.T) {
	tt := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "when array is already sorted",
			input:    []int{1, 2, 3, 4},
			expected: 5,
		},
		{
			name:     "when elements values are larger than array length",
			input:    []int{7, 9, 10},
			expected: 1,
		},
		{
			name:     "when elements are all 0s",
			input:    []int{0, 0},
			expected: 1,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if out := firstMissingPositive(tc.input); out != tc.expected {
				t.Fatalf("expected %v but got %v", tc.expected, out)
			}
		})
	}
}
