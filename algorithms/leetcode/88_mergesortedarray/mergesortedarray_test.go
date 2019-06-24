package mergesortedarray

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tt := []struct {
		name     string
		inputA   []int
		sizeA    int
		inputB   []int
		sizeB    int
		expected []int
	}{
		{
			name:     "when input is valid",
			inputA:   []int{1, 2, 3, 0, 0, 0},
			sizeA:    3,
			inputB:   []int{2, 5, 6},
			sizeB:    3,
			expected: []int{1, 2, 2, 3, 5, 6},
		},
		{
			name:     "when input are both empty",
			inputA:   []int{},
			sizeA:    0,
			inputB:   []int{},
			sizeB:    0,
			expected: []int{},
		},
		{
			name:     "when input b is larger than a",
			inputA:   []int{1, 2, 0, 0, 0},
			sizeA:    2,
			inputB:   []int{3, 4, 5},
			sizeB:    3,
			expected: []int{1, 2, 3, 4, 5},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if merge(tc.inputA, tc.sizeA, tc.inputB, tc.sizeB); !reflect.DeepEqual(tc.inputA, tc.expected) {
				t.Fatalf("expected %v but got %v", tc.expected, tc.inputA)
			}
		})
	}
}
