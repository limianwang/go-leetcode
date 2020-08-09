/**
Given a string, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.

Note: For the purpose of this problem, we define empty string as valid palindrome.
Example 1:

Input: "A man, a plan, a canal: Panama"
Output: true
Example 2:

Input: "race a car"
Output: false
**/
package Validpalindrome

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1

	for i < j {
		if !isValidChar(s[i]) {
			i++
			continue
		}

		if !isValidChar(s[j]) {
			j--
			continue
		}

		if toLower(s[i]) != toLower(s[j]) {
			return false
		}

		i++
		j--
	}

	return true
}

func isValidChar(r byte) bool {
	if (r < '0' || r > '9') && (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
		return false
	}

	return true
}

func toLower(r byte) byte {
	if r >= 'A' && r <= 'Z' {
		return 'a' + r - 'A'
	}

	return r
}
