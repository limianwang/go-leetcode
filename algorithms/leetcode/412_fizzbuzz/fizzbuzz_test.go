package fizzbuzz

import (
	"reflect"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	tt := []struct {
		name   string
		input  int
		output []string
	}{
		{
			name:  "input 1",
			input: 1,
			output: []string{
				"1",
			},
		},
		{
			name:  "input 15",
			input: 15,
			output: []string{
				"1",
				"2",
				"Fizz",
				"4",
				"Buzz",
				"Fizz",
				"7",
				"8",
				"Fizz",
				"Buzz",
				"11",
				"Fizz",
				"13",
				"14",
				"FizzBuzz",
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if actual := fizzBuzz(tc.input); !reflect.DeepEqual(actual, tc.output) {
				t.Fatalf("error fizzbuzz wanted %v but got %v", tc.output, actual)
			}
		})
	}
}
