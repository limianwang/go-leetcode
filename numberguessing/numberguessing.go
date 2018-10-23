package numberguessing

// Pick is the anticipated number
var Pick = 17

func guess(n int) int {
	if n < Pick {
		return -1
	}

	if n > Pick {
		return 1
	}
	return 0
}

// GuessNumber guesses the number that's Picked from 1 ... n. Returns -1 if cannot find.
func GuessNumber(n int) int {
	low := 1
	high := n

	for low <= high {
		mid := low + (high-low)/2
		result := guess(mid)
		// what we guessed is higher
		if result > 0 {
			high = mid - 1
		} else if result < 0 {
			low = mid + 1
		} else {
			return mid
		}
	}

	return -1
}
