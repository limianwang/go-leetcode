package instersectiontwolinkedlist

import (
	"reflect"
	"testing"
)

func TestGetIntersectionNode(t *testing.T) {
	commonNode := &ListNode{
		Val: 10,
		Next: &ListNode{
			Val:  7,
			Next: nil,
		},
	}

	tt := []struct {
		name     string
		input    []*ListNode
		expected *ListNode
	}{
		{
			name: "when input has intersection",
			input: []*ListNode{
				&ListNode{
					Val: 1,
					Next: &ListNode{
						Val:  3,
						Next: commonNode,
					},
				},
				&ListNode{
					Val:  5,
					Next: commonNode,
				},
			},
			expected: commonNode,
		},
		{
			name: "when input has no intersection", input: []*ListNode{
				&ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 10,
						Next: &ListNode{
							Val:  7,
							Next: nil,
						},
					},
				},
				&ListNode{
					Val: 7,
					Next: &ListNode{
						Val:  8,
						Next: nil,
					},
				},
			},
			expected: nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if out := getIntersectionNode(tc.input[0], tc.input[1]); !reflect.DeepEqual(tc.expected, out) {
				t.Fatalf("expected %v but got %v", tc.expected, out)
			}
		})
	}
}
