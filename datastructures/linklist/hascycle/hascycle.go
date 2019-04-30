package hascycle

type Node struct {
	next *Node
}

func HasCycle(node *Node) bool {
	slow := node
	fast := node.next

	for slow != nil && fast != nil && fast.next != nil {
		if slow == fast {
			return true
		}

		slow = slow.next
		fast = fast.next.next
	}

	return false
}
