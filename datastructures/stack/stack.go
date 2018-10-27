package stack

type Stack struct {
	storage []int
}

// New is the Constructor
func New() *Stack {
	return &Stack{
		storage: []int{},
	}
}

func (s *Stack) Push(val int) {
	s.storage = append(s.storage, val)
}

func (s *Stack) Pop() int {
	if len(s.storage) == 0 {
		return -1
	}
	val := s.storage[len(s.storage)-1]
	s.storage = s.storage[:len(s.storage)-1]
	return val
}

func (s *Stack) Peek() int {
	if len(s.storage) == 0 {
		return -1
	}
	return s.storage[len(s.storage)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.storage) == 0
}
