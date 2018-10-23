package numberguessing_test

import (
	"testing"

	"github.com/limianwang/goexamples/numberguessing"
)

func TestGuessNumber(t *testing.T) {
	tt := []struct {
		name     string
		expected int
		input    int
	}{
		{
			name:     "perfect",
			input:    5,
			expected: 4,
		},
		{
			name:     "imperfect",
			input:    40,
			expected: 20,
		},
		{
			name:     "odd",
			input:    0,
			expected: -1,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			numberguessing.Pick = tc.expected
			if r := numberguessing.GuessNumber(tc.input); r != tc.expected {
				t.Fatalf("error expected %v got %v", tc.expected, r)
			}
		})
	}
}
