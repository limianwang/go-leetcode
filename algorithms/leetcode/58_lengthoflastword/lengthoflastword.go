package lengthoflastword

import "strings"

func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}

	arr := strings.Split(s, " ")

	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == "" {
			continue
		} else {
			return len(arr[i])
		}
	}

	return 0
}
