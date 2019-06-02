package validatebinarysearchtree

import "testing"

func TestIsValid(t *testing.T) {
	tt := []struct {
		name     string
		input    *TreeNode
		expected bool
	}{
		{
			name:     "when BST is nil",
			input:    nil,
			expected: true,
		},
		{
			name: "when BST is balance",
			input: &TreeNode{
				Val: 10,
				Left: &TreeNode{
					Val: 5,
				},
				Right: &TreeNode{
					Val: 20,
				},
			},
			expected: true,
		},
		{
			name: "when BST is not valid",
			input: &TreeNode{
				Val: 20,
				Left: &TreeNode{
					Val: 10,
					Left: &TreeNode{
						Val: 7,
					},
					Right: &TreeNode{
						Val: 21,
					},
				},
				Right: &TreeNode{
					Val: 27,
				},
			},
			expected: false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if out := isValidBST(tc.input); out != tc.expected {
				t.Fatalf("expected %v but got %v", tc.expected, out)
			}
		})
	}
}
