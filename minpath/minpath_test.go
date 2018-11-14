package minpath

import "testing"

func TestMinPath(t *testing.T) {
	tt := []struct {
		name     string
		input    [][]int
		expected int
	}{
		{
			name: "succeed1",
			input: [][]int{
				[]int{1, 2, 3},
				[]int{4, 5, 6},
				[]int{7, 8, 9},
			},
			expected: 21,
		},
		{
			name: "succeed2",
			input: [][]int{
				[]int{1, 3, 1},
				[]int{1, 5, 1},
				[]int{4, 2, 1},
			},
			expected: 7,
		},
		{
			name: "test",
			input: [][]int{
				[]int{0},
			},
			expected: 0,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if out := minPathSum(tc.input); out != tc.expected {
				t.Fatalf("expected %v but got %v", tc.expected, out)
			}
		})
	}
}
