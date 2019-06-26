package rotatearray

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	tt := []struct {
		name     string
		input    []int
		k        int
		expected []int
	}{
		{
			name:     "when array rotates 2 times",
			input:    []int{1, 2, 3, 4},
			k:        2,
			expected: []int{3, 4, 1, 2},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if rotate(tc.input, tc.k); !reflect.DeepEqual(tc.expected, tc.input) {
				t.Fatalf("expected %v but got %v", tc.expected, tc.input)
			}
		})
	}
}
