package validparentheses

import (
	"strings"
)

type stack struct {
	arr []string
}

func (s *stack) push(val string) {
	s.arr = append(s.arr, val)
}

func (s *stack) pop() string {
	if len(s.arr) == 0 {
		return ""
	}

	result := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return result
}

func (s *stack) peek() string {
	if len(s.arr) == 0 {
		return ""
	}

	return s.arr[len(s.arr)-1]
}

func newStack() *stack {
	return &stack{
		arr: []string{},
	}
}

func isValid(s string) bool {
	var isOpen = map[string]bool{
		"{": true,
		"[": true,
		"(": true,
	}

	var isClose = map[string]bool{
		"}": true,
		"]": true,
		")": true,
	}

	var isMatched = map[string]string{
		"}": "{",
		"]": "[",
		")": "(",
	}

	stackStorage := newStack()

	for _, k := range strings.Split(s, "") {
		if isOpen[k] {
			stackStorage.push(k)
		}

		if isClose[k] {
			val := stackStorage.peek()
			if val != isMatched[k] {
				return false
			}
			stackStorage.pop()
		}
	}

	return len(stackStorage.arr) == 0
}
