package lrucache

// DoublyLinkedList is a variation of a List Node
type DoublyLinkedList struct {
	Prev  *DoublyLinkedList
	Next  *DoublyLinkedList
	Value int
	Key   int
}

// LRUCache is a cache based on Least Recently Used policy
type LRUCache struct {
	capacity int
	dict     map[int]*DoublyLinkedList
	head     *DoublyLinkedList
	tail     *DoublyLinkedList
}

func (c *LRUCache) put(key int, value int) {
	if node, ok := c.dict[key]; ok {
		c.remove(node)
	}

	node := &DoublyLinkedList{
		Value: value,
		Key:   key,
	}

	c.add(node)
	c.dict[key] = node

	if len(c.dict) > c.capacity {
		node := c.tail.Prev
		c.remove(node)
		delete(c.dict, node.Key)
	}
}

func (c *LRUCache) get(key int) int {
	node := c.dict[key]
	if node != nil {
		c.remove(node)
		c.add(node)
		return node.Value
	}

	return -1
}

func (c *LRUCache) add(node *DoublyLinkedList) {
	next := c.head.Next
	node.Next = next
	node.Prev = c.head
	next.Prev = node
	c.head.Next = node
}

func (c *LRUCache) remove(node *DoublyLinkedList) {
	prev := node.Prev
	next := node.Next

	prev.Next = next
	next.Prev = prev
}

// Constructor is the default exposed version
func Constructor(capacity int) *LRUCache {
	head := &DoublyLinkedList{
		Prev:  nil,
		Next:  nil,
		Value: -1,
		Key:   -1,
	}

	tail := &DoublyLinkedList{
		Prev:  nil,
		Next:  nil,
		Value: -1,
		Key:   -1,
	}

	head.Next = tail
	tail.Prev = head

	return &LRUCache{
		capacity: capacity,
		dict:     map[int]*DoublyLinkedList{},
		head:     head,
		tail:     tail,
	}
}
