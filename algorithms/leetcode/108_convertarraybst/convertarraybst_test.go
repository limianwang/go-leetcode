package convertarraybst

import (
	"reflect"
	"testing"
)

func TestSorter(t *testing.T) {
	tt := []struct {
		name     string
		input    []int
		expected *TreeNode
	}{
		{
			name:     "when array is empty",
			input:    []int{},
			expected: nil,
		},
		{
			name:  "when array is valid and positive (odd)",
			input: []int{1, 2, 3},
			expected: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
		},
		{
			name:  "when array is valid and positive (even)",
			input: []int{1, 2, 3, 4},
			expected: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 1,
					},
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
		},
		{
			name:  "when array is valid and resulting node is nested",
			input: []int{-10, -3, 0, 5, 9},
			expected: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val: -3,
					Left: &TreeNode{
						Val: -10,
					},
				},
				Right: &TreeNode{
					Val: 9,
					Left: &TreeNode{
						Val: 5,
					},
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if out := sortedArrayToBST(tc.input); !reflect.DeepEqual(tc.expected, out) {
				t.Fatalf("expected %v but got %v", tc.expected, out)
			}
		})
	}

}
