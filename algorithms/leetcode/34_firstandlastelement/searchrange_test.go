package searchrange

import (
	"reflect"
	"testing"
)

func TestSearchRange(t *testing.T) {
	tt := []struct {
		name        string
		input       []int
		inputTarget int
		expected    []int
	}{
		{
			name:        "when input has valid data",
			input:       []int{5, 7, 7, 8, 8, 10},
			inputTarget: 8,
			expected:    []int{3, 4},
		},
		{
			name:        "when input has no valid data",
			input:       []int{5, 7, 7, 8, 8, 10},
			inputTarget: 6,
			expected:    []int{-1, -1},
		},
		{
			name:        "when input is empty array",
			input:       []int{},
			inputTarget: 10,
			expected:    []int{-1, -1},
		},
		{
			name:        "when input array is single valued",
			input:       []int{1},
			inputTarget: 1,
			expected:    []int{0, 0},
		},
		{
			name:        "when input is special",
			input:       []int{1, 4},
			inputTarget: 4,
			expected:    []int{1, 1},
		},
		{
			name:        "when input is repetitive",
			input:       []int{1, 2, 3, 3, 3, 3, 4, 5, 9},
			inputTarget: 3,
			expected:    []int{2, 5},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if out := searchRange(tc.input, tc.inputTarget); !reflect.DeepEqual(tc.expected, out) {
				t.Fatalf("expected %v but got %v", tc.expected, out)
			}
		})
	}
}
