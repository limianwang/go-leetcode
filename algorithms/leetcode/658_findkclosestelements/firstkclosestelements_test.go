package findkclosestelements

import (
	"reflect"
	"testing"
)

func TestClosests(t *testing.T) {
	tt := []struct {
		name     string
		inputArr []int
		inputK   int
		inputX   int
		expected []int
	}{
		{
			name:     "when everything just works",
			inputArr: []int{1, 2, 3, 4, 5},
			inputK:   2,
			inputX:   1,
			expected: []int{1, 2},
		},
		{
			name:     "when things should still work",
			inputArr: []int{3, 4, 5, 6},
			inputK:   2,
			inputX:   -1,
			expected: []int{3, 4},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if out := findClosestElements(tc.inputArr, tc.inputK, tc.inputX); !reflect.DeepEqual(out, tc.expected) {
				t.Fatalf("expected %v but instead got %v", tc.expected, out)
			}
		})
	}
}
