package findmindepth

import "testing"

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
				Right: &Node{},
			},
			expected: 3,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if out := minDepth(tc.input); out == tc.expected {
				t.Fatalf("expected %v but got %v", tc.expected, out)
			}
		})
	}
}
