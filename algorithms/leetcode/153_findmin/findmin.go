package findmin

// [4,5,6,7,0,1,2] => 0

func findMin(nums []int) int {
	left := 0
	right := len(nums) - 1

	if len(nums) == 0 {
		return -1
	}

	if nums[left] < nums[right] {
		return nums[left]
	}

	for left <= right {
		mid := (right + left) / 2

		// inflection point
		if nums[mid] > nums[mid+1] {
			return nums[mid+1]
		}

		// second inflection point
		if nums[mid] < nums[mid-1] {
			return nums[mid]
		}

		if nums[mid] > nums[left] {
			left = mid + 1
		} else {
			right = mid - 1
		}

	}

	return -1
}
