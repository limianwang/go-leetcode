/**
Leetcode #64
Given a m x n grid filled with non-negative numbers, find a path from top left to bottom right which minimizes the sum of all numbers along its path.

Note: You can only move either down or right at any point in time.

Example:

Input:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
Output: 7
Explanation: Because the path 1→3→1→1→1 minimizes the sum.
**/
package minpath

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minPathSum(grid [][]int) int {

	m := len(grid)
	n := len(grid[0])
	// initialization code
	result := make([][]int, m)

	for i := range result {
		result[i] = make([]int, n)
	}

	result[0][0] = grid[0][0]

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i >= 1 && j >= 1 {
				result[i][j] = grid[i][j] + min(result[i-1][j], result[i][j-1])
			} else if i == 0 && j >= 1 {
				result[i][j] = grid[i][j] + result[i][j-1]
			} else if j == 0 && i >= 1 {
				result[i][j] = grid[i][j] + result[i-1][j]
			}
		}
	}

	return result[m-1][n-1]
}
