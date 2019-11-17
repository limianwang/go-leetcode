package insertionsort

import (
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	tt := []struct {
		name     string
		input    *ListNode
		expected *ListNode
	}{
		{
			name: "when entering in the middle",
			input: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 3,
						},
					},
				},
			},
			expected: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
						},
					},
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if out := insertionSortList(tc.input); !reflect.DeepEqual(out, tc.expected) {
				t.Fatalf("expected %v but got %v", tc.expected, out)
			}
		})
	}
}
