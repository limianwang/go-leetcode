package onlinestockspan_test

import (
	"testing"

	"github.com/limianwang/goexamples/onlinestockspan"
)

func TestSpan(t *testing.T) {
	tt := []struct {
		name   string
		input  []int
		output []int
	}{
		{
			name:   "success1",
			input:  []int{100, 80, 60, 70, 60, 75, 85},
			output: []int{1, 1, 1, 2, 1, 4, 6},
		},
		{
			name:   "success1",
			input:  []int{29, 91, 62, 76, 51},
			output: []int{1, 2, 1, 2, 1},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			s := &onlinestockspan.StockSpanner{}
			for i, in := range tc.input {
				if out := s.Next(in); tc.output[i] != out {
					t.Fatalf("expected %v but got %v for index %v on val %v", tc.output[i], out, i, tc.input[i])
				}
			}
		})
	}
}
