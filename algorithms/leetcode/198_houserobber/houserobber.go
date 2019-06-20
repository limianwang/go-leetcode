package houserobber

func rob(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	dp := make([]int, length)

	if length >= 1 {
		dp[0] = nums[0]
	}

	if length >= 2 {
		dp[1] = max(nums[1], nums[0])
	}

	if length >= 3 {
		dp[2] = max(nums[0]+nums[2], nums[1])
	}

	for i := 3; i < length; i++ {
		dp[i] = max(nums[i]+dp[i-2], dp[i-1])
	}

	return dp[length-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
