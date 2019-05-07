package stack

type Item interface{}

type Stack struct {
	storage []Item
}

// New is the Constructor
func New() *Stack {
	return &Stack{
		storage: []Item{},
	}
}

func (s *Stack) Push(val Item) {
	s.storage = append(s.storage, val)
}

func (s *Stack) Pop() Item {
	if len(s.storage) == 0 {
		return -1
	}
	val := s.storage[len(s.storage)-1]
	s.storage = s.storage[:len(s.storage)-1]
	return val
}

func (s *Stack) Peek() Item {
	if len(s.storage) == 0 {
		return -1
	}
	return s.storage[len(s.storage)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.storage) == 0
}
