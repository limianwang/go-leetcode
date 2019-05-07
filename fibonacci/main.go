package main

import "fmt"

// fib: f(n) = f(n-1) + f(n-2)
func fibonacci(n int, c chan int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		c <- a
		a, b = b, a+b
	}

	close(c)
}

func recursive(n int) int {
	if n <= 1 {
		return n
	}

	return recursive(n-1) + recursive(n-2)
}

func fibonacci_alternate(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit!!")
			return
		}
	}
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}

	// ch := make(chan int)
	// quit := make(chan int)
	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		fmt.Println(<-ch)
	// 	}

	// 	quit <- 0
	// }()

	// fibonacci_alternate(c, quit)

	fmt.Println("hello", recursive(10))
}
