package firstmissingpositive

import "testing"

func TestFirstMissingPositive(t *testing.T) {
	tt := []struct {
		name   string
		input  []int
		output int
	}{
		{
			name:   "success",
			input:  []int{1, 2, 0},
			output: 3,
		},
		{

			name:   "success1",
			input:  []int{3, 4, -1, 1},
			output: 2,
		},
		{
			name:   "success2",
			input:  []int{7, 8, 9, 11, 12},
			output: 1,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if out := firstMissingPositive(tc.input); out != tc.output {
				t.Fatalf("error in getting firstMissingPositive. expected %v but got %v", tc.output, out)
			}
		})
	}
}
