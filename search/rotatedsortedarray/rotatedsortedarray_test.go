package rotatedsortedarray

import "testing"

func TestSearch(t *testing.T) {
	tt := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{
			name:     "success1",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   0,
			expected: 4,
		},
		{
			name:     "success2",
			nums:     []int{0, 1, 2, 4, 5, 6, 7},
			target:   4,
			expected: 3,
		},

		{
			name:     "success3",
			nums:     []int{6, 7, 0, 1, 2, 4, 5},
			target:   4,
			expected: 5,
		},
		{
			name:     "success4",
			nums:     []int{6, 7, 0, 1, 2, 4, 5},
			target:   7,
			expected: 1,
		},
		{
			name:     "success5",
			nums:     []int{6, 7, 0, 1, 2, 4, 5},
			target:   10,
			expected: -1,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if out := search(tc.nums, tc.target); out != tc.expected {
				t.Fatalf("expected %v but got %v", tc.expected, out)
			}
		})
	}
}
