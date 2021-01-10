/**
Leetcode #63 Unique Paths II

A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).

The robot can only move either down or right at any point in time. The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).

Now consider if some obstacles are added to the grids. How many unique paths would there be?

An obstacle and space is marked as 1 and 0 respectively in the grid.
*/
package uniquepaths

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)

	if m <= 0 {
		return 0
	}

	dp := make([][]int, m)

	n := len(obstacleGrid[0])
	if n <= 0 {
		return 0
	}

	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	if obstacleGrid[0][0] == 1 {
		return 0
	}
	dp[0][0] = 1

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
				continue
			}

			if i > 0 && j > 0 {
				dp[i][j] = dp[i][j-1] + dp[i-1][j]
			} else {
				if j == 0 && i > 0 {
					dp[i][j] = dp[i-1][j]
				}

				if i == 0 && j > 0 {
					dp[i][j] = dp[i][j-1]
				}
			}
		}
	}

	return dp[m-1][n-1]
}
