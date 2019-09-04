package mostwater

import "testing"

func TestMaxArea(t *testing.T) {
	tt := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "when the input is 7x7",
			input:    []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			expected: 49,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if out := maxArea(tc.input); tc.expected != out {
				t.Fatalf("expected: %v but got %v", tc.expected, out)
			}
		})
	}
}
