package fibonaccinumber

import "testing"

func TestFib(t *testing.T) {
	tt := []struct {
		name     string
		input    int
		expected int
	}{
		{
			name:     "when N = 1 it should return 1",
			input:    1,
			expected: 1,
		},
		{
			name:     "when n = 2 it should return 2",
			input:    2,
			expected: 1,
		},
		{
			name:     "when n = 50 it should return a large number",
			input:    50,
			expected: 12586269025,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if out := fib(tc.input); tc.expected != out {
				t.Fatalf("expected %v but got %v", tc.expected, tc.input)
			}
		})
	}
}
