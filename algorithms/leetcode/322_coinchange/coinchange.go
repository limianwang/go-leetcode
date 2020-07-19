package coinchange

import "math"

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount)
	return compute(coins, amount, dp)
}

func compute(coins []int, remainder int, dp []int) int {
	if remainder < 0 {
		return -1
	}

	if remainder == 0 {
		return 0
	}

	if dp[remainder-1] != 0 {
		return dp[remainder-1]
	}

	min := math.MaxInt16

	for _, coin := range coins {
		res := compute(coins, remainder-coin, dp)
		if res >= 0 && res < min {
			min = res + 1
		}
	}

	if min == math.MaxInt16 {
		dp[remainder-1] = -1
	} else {
		dp[remainder-1] = min
	}

	return dp[remainder-1]
}
