package mergesortlist

import (
	"reflect"
	"testing"
)

func TestMergeSortList(t *testing.T) {
	tt := []struct {
		name     string
		input    []*ListNode
		expected *ListNode
	}{
		{
			name: "success1",
			input: []*ListNode{
				&ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
				&ListNode{
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
			expected: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val:  4,
									Next: nil,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "success2",
			input: []*ListNode{
				&ListNode{
					Val:  0,
					Next: nil,
				},
				&ListNode{
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
			expected: &ListNode{
				Val: 0,
				Next: &ListNode{
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
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if out := mergeTwoLists(tc.input[0], tc.input[1]); !reflect.DeepEqual(tc.expected, out) {
				t.Fatalf("expected %v, but got %v", tc.expected, out)
			}
		})
	}
}
