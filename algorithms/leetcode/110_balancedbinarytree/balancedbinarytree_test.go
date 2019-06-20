package balancedbinarytree

import "testing"

func TestIsBalanced(t *testing.T) {
	tt := []struct {
		name     string
		input    *TreeNode
		expected bool
	}{
		{
			name:     "when tree is nil",
			input:    nil,
			expected: true,
		},
		{
			name: "when tree is balanced",
			input: &TreeNode{
				Left: &TreeNode{
					val: 9,
				},
				Right: &TreeNode{
					val: 20,
					Left: &TreeNode{
						val: 15,
					},
					Right: &TreeNode{
						val: 7,
					},
				},
				val: 3,
			},
			expected: true,
		},
		{
			name: "when tree is unbalanced",
			input: &TreeNode{
				val: 1,
				Left: &TreeNode{
					val: 2,
					Left: &TreeNode{
						val: 3,
						Left: &TreeNode{
							val: 4,
						},
						Right: &TreeNode{
							val: 4,
						},
					},
					Right: &TreeNode{
						val: 2,
					},
				},
				Right: &TreeNode{
					val: 2,
				},
			},
			expected: false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if ret := isBalanced(tc.input); ret != tc.expected {
				t.Fatalf("expected %t but got %t", tc.expected, ret)
			}
		})
	}
}
