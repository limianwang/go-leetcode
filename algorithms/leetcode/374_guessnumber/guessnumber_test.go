package guessnumber

import "testing"

func TestGuessNumber(t *testing.T) {
	tt := []struct {
		name        string
		inputNumber int
		expected    int
	}{
		{
			name:        "when number is out of range",
			inputNumber: 5,
			expected:    -1,
		},
		{
			name:        "when number is within range",
			inputNumber: 20,
			expected:    10,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if out := guessNumber(tc.inputNumber); out != tc.expected {
				t.Fatalf("expected: %v but got %v", tc.expected, out)
			}
		})
	}
}
