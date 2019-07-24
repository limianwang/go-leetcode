package sortedarray

/**
[-9,-7,0,1,3]
         ^
	   ^
**/

func sortedSquares(A []int) []int {
	N := len(A)

	dp := make([]int, N)

	j := 0

	for j < N && A[j] < 0 {
		j++
	}

	i := j - 1

	t := 0

	for i >= 0 && j < N {
		if A[i]*A[i] < A[j]*A[j] {
			dp[t] = A[i] * A[i]
			i--
		} else {
			dp[t] = A[j] * A[j]
			j++
		}
		t++
	}

	for i >= 0 {
		dp[t] = A[i] * A[i]
		i--
		t++
	}

	for j < N {
		dp[t] = A[j] * A[j]
		t++
		j++
	}

	return dp
}
