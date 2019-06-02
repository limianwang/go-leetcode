package linkedlistcycle

import (
	"testing"
)

func TestCycle(t *testing.T) {

	tt := []struct {
		name     string
		input    *ListNode
		expected bool
	}{
		{
			name:     "when list are nil",
			input:    nil,
			expected: false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if out := hasCycle(tc.input); out != tc.expected {
				t.Fatalf("expected %v but got %v", tc.expected, out)
			}
		})
	}
}
