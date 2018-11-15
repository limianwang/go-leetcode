package findkclosest

import (
	"reflect"
	"testing"
)

func TestFindClosestElements(t *testing.T) {
	tt := []struct {
		name     string
		arr      []int
		k        int
		x        int
		expected []int
	}{
		{
			name:     "success",
			arr:      []int{0, 1, 2, 3, 4, 5},
			k:        4,
			x:        -1,
			expected: []int{0, 1, 2, 3},
		},
		{
			name:     "success1",
			arr:      []int{0, 1, 2, 3, 4, 5},
			k:        3,
			x:        4,
			expected: []int{3, 4, 5},
		},
		{
			name:     "success2",
			arr:      []int{10, 11, 12, 13, 18},
			k:        2,
			x:        15,
			expected: []int{12, 13},
		},
		{
			name:     "success3",
			arr:      []int{1, 2, 3, 4, 5},
			k:        4,
			x:        3,
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "success4",
			arr:      []int{0, 0, 1, 2, 3, 3, 4, 7, 7, 8},
			k:        3,
			x:        5,
			expected: []int{3, 3, 4},
		},
		{
			name:     "success5",
			arr:      []int{0, 0, 1, 2, 3, 3, 4, 7, 7, 8},
			k:        2,
			x:        2,
			expected: []int{1, 2},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if out := findClosestElements(tc.arr, tc.k, tc.x); !reflect.DeepEqual(out, tc.expected) {
				t.Fatalf("expected %v but got %v", tc.expected, out)
			}
		})
	}
}
