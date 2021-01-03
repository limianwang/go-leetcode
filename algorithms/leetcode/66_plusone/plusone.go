package plusone

func plusOne(digits []int) []int {
	n := len(digits)
	carry := 0

	result := make([]int, n)

	for i := n - 1; i >= 0; i-- {
		var d int

		if i == n-1 {
			// increment
			d = digits[i] + carry + 1
		} else {
			d = digits[i] + carry
		}

		digits[i] = d % 10
		carry = d / 10

		if carry == 0 {
			return digits
		}
	}

	if carry != 0 {
		result = append([]int{1}, digits...)
	}

	return result
}
