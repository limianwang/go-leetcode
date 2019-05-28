package findmindepth

import "testing"

func TestMinDepth(t *testing.T) {
	tt := []struct {
		name     string
		input    *TreeNode
		expected int
	}{
		{
			name:     "when input is nil",
			expected: 0,
			input:    nil,
		},
		{
			name: "when left node has depth 1, and right node depth 3",
			input: &TreeNode{
				Val: 10,
				Left: &TreeNode{
					Val:  0,
					Left: nil,
				},
				Right: &TreeNode{
					Val: 16,
					Left: &TreeNode{
						Val: 14,
						Left: &TreeNode{
							Val:  12,
							Left: nil,
						},
						Right: &TreeNode{
							Val:   15,
							Right: nil,
						},
					},
				},
			},
			expected: 2,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if min := minDepth(tc.input); min != tc.expected {
				t.Fatalf("expected %v but got %v", tc.expected, min)
			}
		})
	}
}
