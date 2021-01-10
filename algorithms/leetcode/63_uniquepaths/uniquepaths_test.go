package uniquepaths

import "testing"

func TestUniquePathsWithObstacles(t *testing.T) {
	tt := []struct {
		name     string
		grid     [][]int
		expected int
	}{
		{
			name: "input grid",
			grid: [][]int{
				[]int{0, 0, 0},
				[]int{0, 1, 0},
				[]int{0, 0, 0},
			},
			expected: 2,
		},
		{
			name: "input small",
			grid: [][]int{
				[]int{0, 1},
				[]int{0, 0},
			},
			expected: 1,
		},
		{
			name: "input row",
			grid: [][]int{
				[]int{0, 0},
			},
			expected: 1,
		},
		{
			name: "input wrong",
			grid: [][]int{
				[]int{1, 0},
			},
			expected: 0,
		},
		{
			name: "input small blocked",
			grid: [][]int{
				[]int{0, 0},
				[]int{1, 1},
				[]int{0, 0},
			},
			expected: 0,
		},
		{
			name: "large case",
			grid: [][]int{
				[]int{0, 1, 0, 0, 0},
				[]int{1, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 0},
			},
			expected: 0,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if output := uniquePathsWithObstacles(tc.grid); output != tc.expected {
				t.Fatalf("expected %d but got %d", tc.expected, output)
			}
		})
	}
}
