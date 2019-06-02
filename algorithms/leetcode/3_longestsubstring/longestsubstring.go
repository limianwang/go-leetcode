package longestsubstring

import (
	"strings"
)

/**
Longest substing without repeating characters

abacabcbb
 ^^^
    ***
*/
func getMax(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func lengthOfLongestSubstring(s string) int {
	left := 0
	right := 0
	max := 0

	contains := map[string]string{}
	strArr := strings.Split(s, "")

	for right < len(strArr) {
		if _, ok := contains[strArr[right]]; ok {
			delete(contains, strArr[left])
			left++
		} else {
			contains[strArr[right]] = strArr[right]
			right++
			max = getMax(max, right-left)
		}
	}

	return max
}
