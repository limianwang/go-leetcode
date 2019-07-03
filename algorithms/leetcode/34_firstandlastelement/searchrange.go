package searchrange

func searchRange(nums []int, target int) []int {
	low := 0
	high := len(nums) - 1

	if len(nums) == 1 {
		if nums[0] == target {
			return []int{0, 0}
		}

		return []int{-1, -1}
	}

	for low <= high {
		mid := (high + low) / 2

		if nums[mid] == target {
			left := mid
			right := mid

			for left > 0 && nums[left-1] == target {
				left--
			}

			for right < len(nums)-1 && nums[right+1] == target {
				right++
			}

			return []int{left, right}
		}

		if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return []int{-1, -1}
}
