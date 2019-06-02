package mergetwosortedlists

// ListNode is a node in a list
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{0, nil}
	curr := head

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			curr.Next = &ListNode{
				Val:  l1.Val,
				Next: nil,
			}
			l1 = l1.Next
		} else {
			curr.Next = &ListNode{
				Val:  l2.Val,
				Next: nil,
			}

			l2 = l2.Next
		}

		curr = curr.Next
	}

	if l1 == nil {
		curr.Next = l2
	} else {
		curr.Next = l1
	}

	return head.Next
}

/**
1->2->4
i
1->3->8
j


1->2->3->4->8
*/
