package queue

import "sync"

// Item the type of the queue
type Item interface{}

//Queue the queue of Items
type Queue struct {
	items []Item
	lock  sync.Mutex
}

// New creates a new Queue
func New() *Queue {
	s := &Queue{}
	s.items = []Item{}
	return s
}

// Enqueue adds an Item to the end of the queue
func (s *Queue) Enqueue(t Item) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.items = append(s.items, t)
}

// Dequeue removes an Item from the front
func (s *Queue) Dequeue() *Item {
	s.lock.Lock()
	defer s.lock.Unlock()
	item := s.items[0]
	s.items = s.items[1:len(s.items)]
	return &item
}

// Peek returns the Item in the front without removing it
func (s *Queue) Peek() *Item {
	s.lock.Lock()
	defer s.lock.Unlock()
	item := s.items[0]
	return &item
}

// IsEmpty returns true when the queue is empty
func (s *Queue) IsEmpty() bool {
	return len(s.items) == 0
}
