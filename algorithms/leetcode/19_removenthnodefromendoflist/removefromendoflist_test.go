package removenthnodefromendoflist

import (
	"reflect"
	"testing"
)

func TestRemove(t *testing.T) {
	tt := []struct {
		name  string
		input struct {
			head *ListNode
			n    int
		}
		expected *ListNode
	}{
		{
			name: "when number is valid",
			input: struct {
				head *ListNode
				n    int
			}{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val:  4,
								Next: nil,
							},
						},
					},
				},
				n: 2,
			},
			expected: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  4,
						Next: nil,
					},
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if node := removeNthFromEnd(tc.input.head, tc.input.n); !reflect.DeepEqual(node, tc.expected) {
				t.Fatalf("expected %v but got %v", tc.input.head, tc.expected)
			}
		})
	}
}
