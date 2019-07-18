package houserobber

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	return max(robber(nums[:len(nums)-1]), robber(nums[1:]))
}

func robber(nums []int) int {
	dp := make([]int, len(nums))

	length := len(nums)

	if length >= 1 {
		dp[0] = nums[0]
	}

	if length >= 2 {
		dp[1] = max(nums[0], nums[1])
	}

	for i := 2; i < length; i++ {
		dp[i] = max(dp[i-1], nums[i]+dp[i-2])
	}

	return dp[length-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
