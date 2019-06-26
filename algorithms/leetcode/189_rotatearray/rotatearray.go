package rotatearray

func rotate(nums []int, k int) {
	arr := make([]int, len(nums))

	for i, val := range nums {
		arr[(i+k)%len(nums)] = val
	}

	for i, val := range arr {
		nums[i] = val
	}
}
