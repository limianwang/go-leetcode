package removeduplicatesfromsortedlist

/*
LeetCode #83
Given a sorted linked list, delete all duplicates such that each element appear only once.

Example 1:

Input: 1->1->2
Output: 1->2
Example 2:

Input: 1->1->2->3->3
Output: 1->2->3
*/

// ListNode represents a node
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	lastSeen := head.Val
	curr := head

	for curr != nil && curr.Next != nil {
		if curr.Next.Val == lastSeen {
			curr.Next = curr.Next.Next
		} else {
			lastSeen = curr.Next.Val
			curr = curr.Next
		}
	}

	return head
}
