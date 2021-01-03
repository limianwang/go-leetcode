package fizzbuzz

/*
412 FizzBuzz
Write a program that outputs the string representation of numbers from 1 to n.

But for multiples of three it should output “Fizz” instead of the number and
for the multiples of five output “Buzz”.  For numbers which are multiples of
both three and five output “FizzBuzz”.

*/

import "fmt"

func fizzBuzz(n int) []string {
	arr := []string{}

	for i := 1; i <= n; i++ {
		entry := ""
		if i%3 == 0 {
			entry += "Fizz"
		}

		if i%5 == 0 {
			entry += "Buzz"
		}

		if entry == "" {
			entry = fmt.Sprintf("%d", i)
		}

		arr = append(arr, entry)
	}

	return arr
}
