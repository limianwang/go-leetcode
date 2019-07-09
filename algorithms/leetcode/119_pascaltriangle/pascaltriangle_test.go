package pascaltriangle

import (
	"reflect"
	"testing"
)

func TestGetRow(t *testing.T) {
	tt := []struct {
		name     string
		input    int
		expected []int
	}{
		{
			name:     "when rowIndex is large",
			input:    3,
			expected: []int{1, 3, 3, 1},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if out := getRow(tc.input); !reflect.DeepEqual(out, tc.expected) {
				t.Fatalf("expected %v but got %v", tc.expected, out)
			}
		})
	}
}
