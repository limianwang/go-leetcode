package insertintobst

import (
	"reflect"
	"testing"
)

func TestInsertBST(t *testing.T) {
	tt := []struct {
		name      string
		inputNode *TreeNode
		inputVal  int
		expected  *TreeNode
	}{
		{
			name:      "when node is nil",
			inputNode: nil,
			inputVal:  5,
			expected: &TreeNode{
				Val: 5,
			},
		},
		{
			name: "when tree is deep",
			inputNode: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
				Right: &TreeNode{
					Val: 7,
				},
			},
			inputVal: 5,
			expected: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
				Right: &TreeNode{
					Val: 7,
					Left: &TreeNode{
						Val: 5,
					},
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if out := insertIntoBST(tc.inputNode, tc.inputVal); !reflect.DeepEqual(tc.expected, out) {
				t.Fatalf("expected %v but got %v", tc.expected, out)
			}
		})
	}
}
