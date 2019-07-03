package minpath

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	dp := make([][]int, m)

	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	dp[0][0] = grid[0][0]

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j >= 1 {
				dp[i][j] = grid[i][j] + dp[i][j-1]
			} else if j == 0 && i >= 1 {
				dp[i][j] = grid[i][j] + dp[i-1][j]
			} else if j >= 1 && i >= 1 {
				dp[i][j] = grid[i][j] + min(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[m-1][n-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}
