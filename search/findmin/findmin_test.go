package findmin

import "testing"

func TestFindMin(t *testing.T) {
	tt := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "success1",
			input:    []int{3, 4, 5, 1, 2},
			expected: 1,
		},
		{
			name:     "success2",
			input:    []int{4, 5, 6, 7, 0, 1, 2},
			expected: 0,
		},
		{
			name:     "success3",
			input:    []int{0, 1, 2},
			expected: 0,
		},
		{
			name:     "success4",
			input:    []int{1},
			expected: 1,
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
