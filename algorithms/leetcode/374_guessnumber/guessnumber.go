package guessnumber

var secretNumber = 10

func guess(num int) int {
	if num == secretNumber {
		return 0
	}

	if num < secretNumber {
		return 1
	}

	return -1
}

func guessNumber(n int) int {
	low := 0
	high := n

	for low < high {
		mid := low + (high-low)/2

		guessed := guess(mid)
		if guessed < 0 {
			low = mid + 1
		} else if guessed > 0 {
			high = mid - 1
		} else {
			return mid
		}
	}

	return -1
}
