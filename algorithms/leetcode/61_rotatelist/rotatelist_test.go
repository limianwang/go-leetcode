package rotatelist

import (
	"reflect"
	"testing"
)

func TestRotateList(t *testing.T) {
	tt := []struct {
		name      string
		inputList *ListNode
		inputK    int
		expected  *ListNode
	}{
		{
			name: "when the input is long and rotate multiple times",
			inputList: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 5,
							},
						},
					},
				},
			},
			inputK: 2,
			expected: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val: 3,
							},
						},
					},
				},
			},
		},
		{
			name: "when the input is valid",
			inputList: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
			inputK: 1,
			expected: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
					},
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if out := rotateRight(tc.inputList, tc.inputK); !reflect.DeepEqual(tc.expected, out) {
				t.Fatalf("expected %v but got %v", tc.expected, out)
			}
		})
	}
}
