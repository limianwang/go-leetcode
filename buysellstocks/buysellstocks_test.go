package buysellstocks

import "testing"

func TestMaxProfile(t *testing.T) {
	tt := []struct {
		name   string
		input  []int
		output int
	}{
		{
			name:   "success1",
			input:  []int{7, 1, 5, 3, 6, 4},
			output: 5,
		},
		{
			name:   "success2",
			input:  []int{7, 6, 4, 3, 1},
			output: 0,
		},
		{
			name:   "success3",
			input:  []int{},
			output: 0,
		},
		{
			name:   "success4",
			input:  []int{1},
			output: 0,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if profit := maxProfit(tc.input); profit != tc.output {
				t.Fatalf("wanted %v but got %v", tc.output, profit)
			}
		})
	}
}
