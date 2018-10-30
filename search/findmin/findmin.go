/*
Package findmin represents Leetcode #153
Find Minimum in Rotated Sorted Array

Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e.,  [0,1,2,4,5,6,7] might become  [4,5,6,7,0,1,2]).

Find the minimum element.

You may assume no duplicate exists in the array.

Example 1:

Input: [3,4,5,1,2]
Output: 1
Example 2:

Input: [4,5,6,7,0,1,2]
Output: 0
*/
package findmin

func findMin(nums []int) int {

	left := 0
	right := len(nums) - 1

	if len(nums) == 1 {
		return nums[0]
	}

	// if already sorted
	if nums[left] < nums[right] {
		return nums[left]
	}

	for left <= right {
		mid := (left + right) / 2

		// if mid is greater than next value => inflation point
		if nums[mid] > nums[mid+1] {
			return nums[mid+1]
		}

		// if mid is less than previous number => inflation point
		if nums[mid] < nums[mid-1] {
			return nums[mid]
		}

		// keep incrementing index bounds
		if nums[mid] > nums[left] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
