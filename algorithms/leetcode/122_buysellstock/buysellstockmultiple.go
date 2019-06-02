package buysellstockmultiple

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	prev := prices[0]
	max := 0

	for i := 1; i < len(prices); i++ {
		if prices[i]-prev > 0 {
			max += prices[i] - prev
		}
		prev = prices[i]
	}
	return max
}
