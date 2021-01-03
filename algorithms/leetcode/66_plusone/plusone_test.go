package plusone

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	tt := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{0, 1},
			expected: []int{0, 2},
		},
		{
			input:    []int{1, 2, 3},
			expected: []int{1, 2, 4},
		},
		{
			input:    []int{4, 3, 2, 1},
			expected: []int{4, 3, 2, 2},
		},
		{
			input:    []int{4, 3, 2, 9},
			expected: []int{4, 3, 3, 0},
		},
		{
			input:    []int{9, 9},
			expected: []int{1, 0, 0},
		},
	}

	for _, tc := range tt {
		output := plusOne(tc.input)
		if !reflect.DeepEqual(output, tc.expected) {
			t.Errorf("%v does not equal %v", output, tc.expected)
		}
	}
}
