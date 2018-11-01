/**
Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.

Example:

Input: 1->2->4, 1->3->4
Output: 1->1->2->3->4->4
*/
package mergesortlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	curr := &ListNode{
		Val:  0,
		Next: nil,
	}

	root := curr

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			curr.Next = &ListNode{Val: l1.Val, Next: nil}
			l1 = l1.Next
		} else {
			curr.Next = &ListNode{Val: l2.Val, Next: nil}
			l2 = l2.Next
		}

		curr = curr.Next
	}

	if l1 == nil {
		curr.Next = l2
	} else {
		curr.Next = l1
	}

	return root.Next
}
