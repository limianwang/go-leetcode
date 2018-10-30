/**
Find the first peak
Leedcode #162
*/

package firstpeak

func findPeakElement(nums []int) int {

	left := 0
	right := len(nums) - 1

	for left <= right {
		if left == right {
			return left
		}

		if left+1 == right {
			if nums[left] > nums[right] {
				return left
			}
			return right
		}

		mid := (left + right) / 2

		if nums[mid] > nums[mid-1] && nums[mid+1] < nums[mid] {
			return mid
		}

		if nums[mid-1] < nums[mid] && nums[mid] < nums[mid+1] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
