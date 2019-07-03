package pascaltriangle

import (
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {
	tt := []struct {
		name     string
		input    int
		expected [][]int
	}{
		{
			name:     "when numRows equals to 0",
			input:    0,
			expected: [][]int{},
		},
		{
			name:     "when rows is valid",
			input:    2,
			expected: [][]int{[]int{1}, []int{1, 1}},
		},
		{
			name:  "when rows is multiple layers",
			input: 5,
			expected: [][]int{
				[]int{1},
				[]int{1, 1},
				[]int{1, 2, 1},
				[]int{1, 3, 3, 1},
				[]int{1, 4, 6, 4, 1},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if out := generate(tc.input); !reflect.DeepEqual(tc.expected, out) {
				t.Fatalf("expected %v, but got %v", tc.expected, out)
			}
		})
	}
}
