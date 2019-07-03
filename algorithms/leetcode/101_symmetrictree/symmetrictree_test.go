package symmetrictree

import "testing"

func TestIsSymmetric(t *testing.T) {
	tt := []struct {
		name     string
		input    *TreeNode
		expected bool
	}{
		{
			name: "when trees are symmetric",
			input: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 7,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			expected: true,
		},
		{
			name: "when trees are not symmetric",
			input: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 7,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			expected: false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if out := isSymmetric(tc.input); out != tc.expected {
				t.Fatalf("expected %v but got %v", tc.expected, out)
			}
		})
	}
}
