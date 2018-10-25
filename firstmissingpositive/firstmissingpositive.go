package firstmissingpositive

func firstMissingPositive(nums []int) int {
	length := len(nums)

	for i := 0; i < length; i++ {
		if nums[i] >= 1 && nums[i] <= length && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
			i--
		}
	}

	for i, num := range nums {
		if num != i+1 {
			return i + 1
		}
	}

	return length + 1
}
