package mostwater

func maxArea(height []int) int {
	maxArea := 0
	left := 0
	right := len(height) - 1

	for left < right {
		maxArea = getMax(maxArea, getMin(height[left], height[right])*(right-left))

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}

func getMax(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func getMin(a, b int) int {
	if a > b {
		return b
	}
	return a
}
