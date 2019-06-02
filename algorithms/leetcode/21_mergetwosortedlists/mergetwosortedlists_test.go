package mergetwosortedlists

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	tt := []struct {
		name     string
		input    []*ListNode
		expected *ListNode
	}{
		{
			name: "when two lists are different",
			input: []*ListNode{
				&ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val:  5,
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
									Val:  5,
									Next: nil,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "when two lists are empty",
			input: []*ListNode{
				nil,
				nil,
			},
			expected: nil,
		},
		{
			name: "when only one of the list is available",
			input: []*ListNode{
				&ListNode{
					Val:  0,
					Next: nil,
				},
				nil,
			},
			expected: &ListNode{0, nil},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if out := mergeTwoLists(tc.input[0], tc.input[1]); !reflect.DeepEqual(out, tc.expected) {
				expected, _ := json.Marshal(tc.expected)
				output, _ := json.Marshal(out)
				t.Fatalf("expected %s but got %s", expected, output)
			}
		})
	}
}
