package validparentheses

import (
	"fmt"
	"strings"
)

func isValid(s string) bool {
	open := []string{}

	var validOpen = map[string]bool{
		"{": true,
		"[": true,
		"(": true,
	}
	var validClose = map[string]bool{
		"}": true,
		"]": true,
		")": true,
	}

	var mapping = map[string]bool{
		"{}": true,
		"()": true,
		"[]": true,
	}

	for _, s := range strings.Split(s, "") {
		if validOpen[s] {
			open = append(open, s)
		}

		if validClose[s] {
			if len(open) == 0 {
				return false
			}
			if !mapping[fmt.Sprintf("%s%s", open[len(open)-1], s)] {
				return false
			}
			open = open[:len(open)-1]
		}
	}

	if len(open) == 0 {
		return true
	}
	return false
}
