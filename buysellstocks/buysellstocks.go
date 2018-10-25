package buysellstocks

func maxProfit(prices []int) int {
	length := len(prices)
	max := 0
	for i := length - 1; i >= 1; i-- {
		for j := 0; j < i-1; j++ {
			tmp := prices[i] - prices[j]
			if max < tmp {
				max = tmp
			}
		}
	}

	return max
}
