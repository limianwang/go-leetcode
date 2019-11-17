package insertionsort

// ListNode
type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	curr := head

	for curr != nil && curr.Next != nil {
		p := curr.Next
		if curr.Val <= p.Val {
			curr = p
			continue
		}

		curr.Next = p.Next
		pre, next := dummy, dummy.Next

		for p.Val > next.Val {
			pre = next
			next = next.Next
		}

		pre.Next = p
		p.Next = next
	}

	return dummy.Next
}
