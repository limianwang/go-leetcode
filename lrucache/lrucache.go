package lrucache

type Node struct {
	Key   interface{}
	Value interface{}
	Prev  *Node
	Next  *Node
}

type Cacher interface {
	Get(key interface{}) interface{}
	Set(key interface{}, value interface{})
}

type Cache struct {
	cap  int
	dict map[interface{}]*Node
	head *Node
	tail *Node
}

func (c *Cache) Get(key interface{}) interface{} {
	if val, ok := c.dict[key]; ok {
		c.remove(val)
		c.add(val)
		return val.Value
	}
	return -1
}

func (c *Cache) Set(key interface{}, value interface{}) {
	if n, ok := c.dict[key]; ok {
		c.remove(n)
	}

	n := &Node{Key: key, Value: value}
	c.add(n)
	c.dict[key] = n
	if len(c.dict) > c.cap {
		n := c.head.Next
		c.remove(n)
		delete(c.dict, n.Key)
	}
}

func (c *Cache) add(node *Node) {
	prev := c.tail.Prev
	c.tail.Prev = node
	node.Next = c.tail
	node.Prev = prev
	prev.Next = node
}

func (c *Cache) remove(node *Node) {
	prev := node.Prev
	next := node.Next

	prev.Next = next
	next.Prev = prev
}

func NewCache(capacity int) Cacher {
	head := &Node{Key: 0, Value: 0}
	tail := &Node{Key: 0, Value: 0}
	head.Next = tail
	tail.Prev = head

	return &Cache{
		cap:  capacity,
		dict: map[interface{}]*Node{},
		head: head,
		tail: tail,
	}
}
