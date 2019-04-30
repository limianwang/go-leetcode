package validparentheses

import (
	"fmt"
	"strings"
)

type Stack struct {
	storage []string
}

// NewStack is the Constructor
func NewStack() *Stack {
	return &Stack{
		storage: []string{},
	}
}

func (s *Stack) Push(val string) {
	s.storage = append(s.storage, val)
}

func (s *Stack) Pop() string {
	if len(s.storage) == 0 {
		return ""
	}
	val := s.storage[len(s.storage)-1]
	s.storage = s.storage[:len(s.storage)-1]
	return val
}

func (s *Stack) Peek() string {
	if len(s.storage) == 0 {
		return ""
	}
	return s.storage[len(s.storage)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.storage) == 0
}

func isBalanced(s string) bool {
	var isValidOpen = map[string]bool{
		"{": true,
		"[": true,
		"(": true,
	}

	var isValidClose = map[string]bool{
		"}": true,
		"]": true,
		")": true,
	}

	var matchedOpenGivenClose = map[string]string{
		"}": "{",
		"]": "[",
		")": "(",
	}

	storage := NewStack()

	for _, k := range strings.Split(s, "") {
		if isValidOpen[k] {
			storage.Push(k)
		}

		if isValidClose[k] {
			open := storage.Pop()
			if open != matchedOpenGivenClose[k] {
				return false
			}
		}
	}

	return storage.IsEmpty()
}

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
