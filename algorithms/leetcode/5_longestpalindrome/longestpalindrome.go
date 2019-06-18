package longestpalindrome

func longestPalindrome(s string) string {
	length := len(s)

	start := 0
	maxLength := 1

	dp := make([][]bool, length)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, length)
	}

	// Single letters palindromes
	for i := 0; i < length; i++ {
		dp[i][i] = true
	}

	for i := 0; i < length-1; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = true
			start = 1
			maxLength = 2
		}
	}

	for k := 3; k <= length; k++ {
		for i := 0; i < length-k+1; i++ {
			j := i + k - 1
			if s[i] == s[j] && dp[i+1][j-1] {
				dp[i][j] = true
				start = i
				maxLength = j + 1
			}
		}
	}

	return s[start:maxLength]
}
