package subarray

/**
nums = []int{1,2,3}
k = 6
*/

func subarraySum(nums []int, k int) int {
	count := 0

	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				count++
			}
		}
	}

	return count
}
