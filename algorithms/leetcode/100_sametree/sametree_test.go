package sametree

import "testing"

func TestIsSameTree(t *testing.T) {
	tt := []struct {
		name     string
		inputA   *TreeNode
		inputB   *TreeNode
		expected bool
	}{
		{
			name: "when the inputs are same",
			inputA: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: 5,
				},
			},
			inputB: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: 5,
				},
			},
			expected: true,
		},
		{
			name: "when the inputs are not the same",
			inputA: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 5,
				},
			},
			inputB: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: 5,
				},
			},
			expected: false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if out := isSameTree(tc.inputA, tc.inputB); tc.expected != out {
				t.Fatalf("expected %v but got %v", tc.expected, out)
			}
		})
	}
}
