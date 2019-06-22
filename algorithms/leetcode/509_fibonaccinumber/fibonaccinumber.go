package fibonaccinumber

func fib(N int) int {
	dp := make([]int, N)

	if N >= 1 {
		dp[0] = 1
	}
	if N >= 2 {
		dp[1] = 1
	}

	for i := 2; i < N; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[N-1]
}
