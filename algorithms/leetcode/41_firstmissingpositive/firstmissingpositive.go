package firstmissingpositive

func firstMissingPositive(nums []int) int {
	specialSort(nums)

	// Once the array has been specially sorted. We now compare the index to the value that the nums are associated with. We return when the values sequentially sorted (i.e. 1, 2, 4; we then return 3).
	for index, num := range nums {
		if num != index+1 {
			return index + 1
		}
	}

	// if the value of the numbers are accurate, the next available will essentially be the length of the array + 1.
	return len(nums) + 1
}

/**
This function takes in an array and does a "special sort", where we key on the lowest value being 1 instead of 0.

When a number for a given index is lower than the length, and when the number is greater than 1, and when the number is not equal the nums[values-1], we then do a swap.


Why nums[i] >= 1?
This is what we key it on.

Why nums[i] <= len(nums)?
We don't need to worry about the numbers that are greater than length, as we know that the values are going to be less than that. (We want a fully sorted array starting with 1).

Why nums[i] != nums[nums[i]-1]?
We want to make sure that we only swap when the values we are checking are not the same. We are NOT concerned about overflow here, because we know that nums[i]-1 will always be valid, since nums[i] has already been vetted to be less than len(nums).

*/
func specialSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		if nums[i] >= 1 && nums[i] <= len(nums) && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
			i--
		}
	}
}
