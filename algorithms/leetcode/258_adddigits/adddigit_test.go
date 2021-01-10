package adddigits

import "testing"

func TestAddDigits(t *testing.T) {
	tt := []struct {
		input    int
		expected int
		name     string
	}{
		{
			name:     "adding one digit",
			expected: 2,
			input:    38,
		},
		{
			name:     "test zero",
			expected: 0,
			input:    0,
		},
		{
			name:     "test nine",
			expected: 9,
			input:    9,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if output := addDigits(tc.input); output != tc.expected {
				t.Fatalf("error %v is not equal to %v", tc.expected, output)
			}
		})
	}
}
