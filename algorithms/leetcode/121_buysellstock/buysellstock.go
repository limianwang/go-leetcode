package buysellstock

import "math"

func maxProfit(prices []int) int {
	minPrice := math.MaxInt16
	max := 0

	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else if price-minPrice > max {
			max = price - minPrice
		}
	}

	return max
}
