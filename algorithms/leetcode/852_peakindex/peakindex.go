package peakindex

func peakIndexInMountainArray(A []int) int {
	low := 0
	high := len(A) - 1

	for low < high {
		mid := (high + low) / 2

		if mid > 0 && mid+1 <= len(A)-1 && A[mid] > A[mid-1] && A[mid] > A[mid+1] {
			return mid
		}

		if A[mid] < A[mid+1] && A[mid] > A[mid-1] {
			low++
		} else {
			high--
		}

	}

	return -1
}
