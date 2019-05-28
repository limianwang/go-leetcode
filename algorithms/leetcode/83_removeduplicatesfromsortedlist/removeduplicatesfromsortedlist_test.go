package removeduplicatesfromsortedlist

import (
	"reflect"
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	tt := []struct {
		name     string
		input    *ListNode
		expected *ListNode
	}{
		{
			name: "three same numbers",
			input: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val:  1,
						Next: nil,
					},
				},
			},
			expected: &ListNode{
				Val:  1,
				Next: nil,
			},
		},
		{
			name: "remove dups at end",
			input: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  2,
						Next: nil,
					},
				},
			},
			expected: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  2,
					Next: nil,
				},
			},
		},
		{
			name: "remove dups with middle odd",
			input: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val:  2,
								Next: nil,
							},
						},
					},
				},
			},
			expected: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val:  2,
						Next: nil,
					},
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if node := deleteDuplicates(tc.input); !reflect.DeepEqual(node, tc.expected) {
				t.Fatalf("expected %v but got %v", tc.expected, node)
			}
		})
	}
}
