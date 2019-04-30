package hascycle

import "testing"

func TestHasCycle(t *testing.T) {
	tt := []struct {
		name     string
		input    *Node
		expected bool
	}{
		{
			name: "test",
			input: &Node{
				next: &Node{
					next: &Node{
						next: nil,
					},
				},
			},
			expected: false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if out := HasCycle(tc.input); out != tc.expected {
				t.Fatalf("expected %v but got %v", tc.expected, out)
			}
		})
	}
}
