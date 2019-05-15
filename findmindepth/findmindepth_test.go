package findmindepth

import (
	"fmt"
	"testing"
)

func TestMinDepth(t *testing.T) {
	tt := []struct {
		name     string
		input    *Node
		expected int
	}{
		{
			name: "success1",
			input: &Node{
				Val: 10,
				Left: &Node{
					Val: 0,
					Left: &Node{
						Val: 1,
					},
				},
				Right: nil,
			},
			expected: 3,
		},
		{
			name: "when left node has depth 1, and right node depth 3",
			input: &Node{
				Val: 10,
				Left: &Node{
					Val:  0,
					Left: nil,
				},
				Right: &Node{
					Val: 16,
					Left: &Node{
						Val: 14,
						Left: &Node{
							Val:  12,
							Left: nil,
						},
						Right: &Node{
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
			if out := minDepth(tc.input); out != tc.expected {
				t.Fatalf("expected %v but got %v", tc.expected, out)
			} else {
				fmt.Println(out, tc.expected)
			}
		})
	}
}
