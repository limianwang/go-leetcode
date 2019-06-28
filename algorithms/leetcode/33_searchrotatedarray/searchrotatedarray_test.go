package searchrotatedarray

import "testing"

func TestSearch(t *testing.T) {
	tt := []struct {
		name        string
		inputArr    []int
		inputTarget int
		expected    int
	}{
		{
			name:        "when item is found",
			inputArr:    []int{4, 5, 6, 7, 0, 1, 2},
			inputTarget: 0,
			expected:    4,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if out := search(tc.inputArr, tc.inputTarget); out != tc.expected {
				t.Fatalf("expected %d but got %d", tc.expected, out)
			}
		})
	}

}
