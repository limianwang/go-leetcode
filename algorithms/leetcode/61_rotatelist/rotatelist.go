package rotatelist

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
a->b->c

k= 1 c->a->b
*/

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	tail := head
	size := 1

	for tail.Next != nil {
		tail = tail.Next
		size++
	}

	tail.Next = head

	for i := 1; i < size-k%size; i++ {
		head = head.Next
	}

	resp := head.Next
	head.Next = nil

	return resp
}
