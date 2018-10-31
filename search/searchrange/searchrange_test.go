package searchrange

import (
	"reflect"
	"testing"
)

func TestSearchRange(t *testing.T) {
	tt := []struct {
		name   string
		input  []int
		target int
		output []int
	}{
		{
			name:   "success1",
			input:  []int{5, 7, 7, 8, 8, 10},
			target: 8,
			output: []int{3, 4},
		},
		{
			name:   "success2",
			input:  []int{5, 7, 7, 8, 8, 10},
			target: 6,
			output: []int{-1, -1},
		},
		{
			name:   "success3",
			input:  []int{7, 7, 7, 7, 8, 8, 10},
			target: 7,
			output: []int{0, 3},
		},
		{
			name:   "success4",
			input:  []int{1},
			target: 1,
			output: []int{0, 0},
		},
		{
			name:   "success5",
			input:  []int{2, 2},
			target: 2,
			output: []int{0, 1},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if out := searchRange(tc.input, tc.target); !reflect.DeepEqual(tc.output, out) {
				t.Fatalf("expected: %v but got %v", tc.output, out)
			}
		})
	}
}
