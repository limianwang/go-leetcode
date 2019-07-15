package twosum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tt := []struct {
		name        string
		inputArr    []int
		inputTarget int
		expected    []int
	}{
		{
			name:        "when value is found",
			inputArr:    []int{2, 7, 11, 15},
			inputTarget: 9,
			expected:    []int{0, 1},
		},
		{
			name:        "when value is not found",
			inputArr:    []int{2, 7, 11, 15},
			inputTarget: 11,
			expected:    []int{-1, -1},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if out := twoSum(tc.inputArr, tc.inputTarget); !reflect.DeepEqual(tc.expected, out) {
				t.Fatalf("expected %v but got %v", tc.expected, out)
			}
		})
	}

}
