package maxsubarray

func maxSubArray(nums []int) int {
	curr, max := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		if curr < 0 {
			curr = nums[i]
		} else {
			curr += nums[i]
		}

		if curr > max {
			max = curr
		}
	}

	return max
}
