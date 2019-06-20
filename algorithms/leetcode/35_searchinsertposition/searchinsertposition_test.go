package main

import "testing"

func TestSearchInsert(t *testing.T) {
	tt := []struct {
		name        string
		inputArr    []int
		inputTarget int
		expected    int
	}{
		{
			name:        "when item is found",
			inputArr:    []int{1, 3, 5, 6},
			inputTarget: 5,
			expected:    2,
		},
		{
			name:        "when the item is not found but should be in middle",
			inputTarget: 2,
			inputArr:    []int{1, 3, 5, 6},
			expected:    1,
		},
		{
			name:        "when the item is not found but should be last",
			inputTarget: 7,
			inputArr:    []int{1, 3, 5, 6},
			expected:    4,
		},
		{
			name:        "when the item is first",
			inputTarget: 0,
			inputArr:    []int{1, 3, 5, 6},
			expected:    0,
		},
		{
			name:        "when the input is not a valid array",
			inputTarget: 1,
			inputArr:    []int{},
			expected:    0,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if out := searchInsert(tc.inputArr, tc.inputTarget); out != tc.expected {
				t.Fatalf("expected %v but got %v", tc.expected, out)
			}
		})
	}
}
