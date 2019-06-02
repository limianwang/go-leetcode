package findkclosestelements

/*
This is a sorted array.

We need to do a search to find the smallest differences

k = number of elements that are closests to the number (x).

*/
func findClosestElements(arr []int, k int, x int) []int {
	low := 0
	high := len(arr) - k

	for low < high {
		mid := low + (high-low)/2
		if x-arr[mid] > arr[mid+k]-x {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return arr[low : low+k]
}
