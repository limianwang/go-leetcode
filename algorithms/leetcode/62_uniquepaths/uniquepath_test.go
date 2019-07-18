package uniquepath

import "testing"

func TestUniquePath(t *testing.T) {
	tt := []struct {
		name     string
		m        int
		n        int
		expected int
	}{
		{
			name:     "when input is small",
			m:        3,
			n:        2,
			expected: 3,
		},
		{
			name:     "when input is large",
			m:        7,
			n:        3,
			expected: 28,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if out := uniquePaths(tc.m, tc.n); out != tc.expected {
				t.Fatalf("expected %d but got %d", tc.expected, out)
			}
		})
	}
}
