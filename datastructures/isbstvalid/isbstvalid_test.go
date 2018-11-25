package isbstvalid_test

import (
	"testing"

	"github.com/limianwang/goexamples/datastructures/isbstvalid"
)

func TestIsValid(t *testing.T) {
	tt := []struct {
		name     string
		input    *isbstvalid.TreeNode
		expected bool
	}{
		{
			name: "success1",
			input: &isbstvalid.TreeNode{
				Left: &isbstvalid.TreeNode{
					Val: 1,
				},
				Right: &isbstvalid.TreeNode{
					Val: 3,
				},
				Val: 2,
			},
			expected: true,
		},
		{
			name: "success2",
			input: &isbstvalid.TreeNode{
				Left: &isbstvalid.TreeNode{
					Val: 1,
				},
				Right: &isbstvalid.TreeNode{
					Val: 4,
					Left: &isbstvalid.TreeNode{
						Val: 3,
					},
					Right: &isbstvalid.TreeNode{
						Val: 6,
					},
				},
				Val: 5,
			},
			expected: false,
		},
		{
			name: "success2",
			input: &isbstvalid.TreeNode{
				Left: &isbstvalid.TreeNode{
					Val: 5,
				},
				Right: &isbstvalid.TreeNode{
					Val: 15,
					Left: &isbstvalid.TreeNode{
						Val: 6,
					},
					Right: &isbstvalid.TreeNode{
						Val: 20,
					},
				},
				Val: 10,
			},
			expected: false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if out := isbstvalid.IsBSTValid(tc.input); out != tc.expected {
				t.Fatalf("expected to be %v but was not", tc.expected)
			}
		})
	}
}
