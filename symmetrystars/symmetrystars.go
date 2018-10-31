package main

import (
	"fmt"
	"strings"
)

func printStars(n int) {
	arr := make([][]string, n)
	mid := n / 2

	for i := 0; i <= mid; i++ {
		row := make([]string, n)
		upper := mid + i
		lower := mid - i
		for x := 0; x < n; x++ {
			if x >= lower && x <= upper {
				row[x] = "*"
			} else {
				row[x] = " "
			}
		}
		arr[i] = row
	}

	for i := mid + 1; i < n; i++ {
		row := make([]string, n)
		delta := n - 1 - i
		upper := mid + delta
		lower := mid - delta
		for x := 0; x < n; x++ {
			if x >= lower && x <= upper {
				row[x] = "*"
			} else {
				row[x] = " "
			}
		}
		arr[i] = row
	}

	r := make([]string, n)
	for i, k := range arr {
		r[i] = strings.Join(k, "")
	}

	fmt.Println(strings.Join(r, "\n"))
}

func main() {
	printStars(5)
}
