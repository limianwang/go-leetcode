package inorder

import (
	"reflect"
	"testing"
)

func TestIncreasingBST(t *testing.T) {
	tt := []struct {
		name     string
		input    *TreeNode
		expected *TreeNode
	}{
		{
			name: "expected result",
			input: &TreeNode{
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
					Val: 6,
					Left: &TreeNode{
						Val: 5,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			expected: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
						Right: &TreeNode{
							Val: 4,
							Right: &TreeNode{
								Val: 5,
								Right: &TreeNode{
									Val: 6,
									Right: &TreeNode{
										Val: 7,
									},
								},
							},
						},
					},
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if out := increasingBST(tc.input); !reflect.DeepEqual(tc.expected, out) {
				t.Fatalf("expected %v but got %v", tc.expected, out)
			}
		})
	}
}
