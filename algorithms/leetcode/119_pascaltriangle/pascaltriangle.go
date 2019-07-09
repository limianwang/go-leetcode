package pascaltriangle

func getRow(rowIndex int) []int {
	dp := make([][]int, rowIndex+1)

	if rowIndex >= 0 {
		dp[0] = []int{1}
	}

	if rowIndex >= 1 {
		dp[1] = []int{1, 1}
	}

	for i := 2; i <= rowIndex; i++ {
		row := make([]int, i+1)
		lastRow := dp[i-1]
		for j := 0; j <= i; j++ {
			if i == j || j == 0 {
				row[j] = 1
			} else {
				row[j] = lastRow[j-1] + lastRow[j]
			}
		}
		dp[i] = row
	}

	return dp[rowIndex]
}
