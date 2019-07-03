package minpath

import "testing"

func TestMinPath(t *testing.T) {
	tt := []struct {
		name     string
		input    [][]int
		expected int
	}{
		{
			name: "when input is valid",
			input: [][]int{
				[]int{1, 3, 1},
				[]int{1, 5, 1},
				[]int{4, 2, 1},
			},
			expected: 7,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if out := minPathSum(tc.input); out != tc.expected {
				t.Fatalf("expected %v, but got %v", tc.expected, out)
			}
		})
	}
}
