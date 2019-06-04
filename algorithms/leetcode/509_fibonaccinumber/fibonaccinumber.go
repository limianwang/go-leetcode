package fibonaccinumber

var cache map[int]int

func init() {
	cache = map[int]int{0: 0, 1: 1}
}

func fib(N int) int {
	if val, ok := cache[N]; ok {
		return val
	}

	result := fib(N-1) + fib(N-2)
	cache[N] = result
	return result
}
