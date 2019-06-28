package stack_test

import (
	"testing"

	"github.com/limianwang/go-leetcode/datastructures/stack"
)

func TestStack(t *testing.T) {
	s := stack.New()
	s.Push(5)
	if s.Peek() != 5 {
		t.Fatalf("error peeking, got %v expected %v", s.Peek(), 5)
	}
}
