/*
Given an array of integers nums sorted in ascending order, find the starting and ending position of a given target value.

Your algorithm's runtime complexity must be in the order of O(log n).

If the target is not found in the array, return [-1, -1].

Example 1:

Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]
Example 2:

Input: nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]
*/
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

		if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return []int{-1, -1}
}
