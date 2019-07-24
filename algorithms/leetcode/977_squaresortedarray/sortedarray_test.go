package sortedarray

import (
	"reflect"
	"testing"
)

func TestSortedArray(t *testing.T) {
	tt := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "when input is small",
			input:    []int{-4, -3, 0, 1, 2},
			expected: []int{0, 1, 4, 9, 16},
		},
		{
			name:     "when input is normal",
			input:    []int{-7, -3, 2, 3, 11},
			expected: []int{4, 9, 9, 49, 121},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if out := sortedSquares(tc.input); !reflect.DeepEqual(out, tc.expected) {
				t.Fatalf("expected %v but got %v", tc.expected, out)
			}
		})
	}
}
