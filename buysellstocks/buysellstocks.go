package buysellstocks

import (
	"math"
)

func maxProfit(prices []int) int {
	length := len(prices)
	max := 0
	// Default set to max INT
	minPrice := math.MaxInt64

	for i := 0; i < length; i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > max {
			max = prices[i] - minPrice
		}
	}

	return max
}
