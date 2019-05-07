package linkedlist

import "sync"

// Item is the item stored as the value of the linkedlist
type Item interface{}

// Node is the representation of each item on the LinkedList
type Node struct {
	value Item
	next  *Node
}

// LinkedList is the datastructure
type LinkedList struct {
	size int
	head *Node
	lock sync.RWMutex
}

// AddFirst adds an Item to the beginning of the LinkedList
func (ll *LinkedList) AddFirst(i Item) {
	ll.lock.Lock()
	defer ll.lock.Unlock()
	node := Node{i, nil}
	if ll.head == nil {
		ll.head = &node
	} else {
		node.next = ll.head
		ll.head = &node
	}

	ll.size++
}

// AddLast adds an Item to the end of the LinkedList
func (ll *LinkedList) AddLast(i Item) {
	ll.lock.Lock()
	defer ll.lock.Unlock()

	node := Node{i, nil}

	if ll.head == nil {
		ll.head = &node
	} else {
		last := ll.head
		for last.next != nil {
			last = last.next
		}

		last.next = &node
	}
	ll.size++
}

// IsEmpty returns true when no elements are within the LinkedList
func (ll *LinkedList) IsEmpty() bool {
	ll.lock.RLock()
	defer ll.lock.RUnlock()
	if ll.head == nil {
		return true
	}
	return false
}
