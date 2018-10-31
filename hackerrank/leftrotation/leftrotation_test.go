package leftrotation

import (
	"reflect"
	"testing"
)

func TestLeftRotation(t *testing.T) {
	tt := []struct {
		name      string
		input     []int
		iteration int
		expected  []int
	}{
		{
			name:      "success",
			input:     []int{1, 2, 3, 4, 5},
			iteration: 4,
			expected:  []int{5, 1, 2, 3, 4},
		},
		{
			name:      "success",
			input:     []int{1, 2, 3, 4, 5},
			iteration: 2,
			expected:  []int{3, 4, 5, 1, 2},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if out := rotateArray(tc.input, tc.iteration); !reflect.DeepEqual(tc.expected, out) {
				t.Fatalf("expected %v but got %v", tc.expected, out)
			}
		})
	}
}
