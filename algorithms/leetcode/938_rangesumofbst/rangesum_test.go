package rangerum

import "testing"

func TestRangeSum(t *testing.T) {
	tt := []struct {
		name     string
		input    *TreeNode
		L        int
		R        int
		expected int
	}{
		{
			name: "when input is valid",
			input: &TreeNode{
				Val: 10,
				Left: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 3,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
				Right: &TreeNode{
					Val: 15,

					Right: &TreeNode{
						Val: 18,
					},
				},
			},
			L:        7,
			R:        15,
			expected: 32,
		},
		{
			name: "when input is valid",
			input: &TreeNode{
				Val: 10,
				Left: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 1,
						},
					},
					Right: &TreeNode{
						Val: 7,
						Left: &TreeNode{
							Val: 6,
						},
					},
				},
				Right: &TreeNode{
					Val: 15,
					Left: &TreeNode{
						Val: 13,
					},
					Right: &TreeNode{
						Val: 18,
					},
				},
			},
			L:        6,
			R:        10,
			expected: 23,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if out := rangeSumBST(tc.input, tc.L, tc.R); out != tc.expected {
				t.Fatalf("expected %v but got %v", tc.expected, out)
			}
		})
	}
}
