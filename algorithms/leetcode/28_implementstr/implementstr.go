package implementstr

func strStr(haystack string, needle string) int {
	needleLength := len(needle)
	hsLength := len(haystack)

	if needleLength > hsLength {
		return -1
	}

	for i := 0; i <= hsLength-needleLength; i++ {
		if haystack[i:i+needleLength] == needle {
			return i
		}
	}

	return -1
}
