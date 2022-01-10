package gocollections

type IStack interface {
	Push(interface{})
	Pop() interface{}
	Peek() interface{}
	IsEmpty() bool
	Size() int
}

type IntStack struct {
	size  int
	stack []int
}

func (s *IntStack) Push(e interface{}) {
	s.stack = append(s.stack, e.(int))
	s.size += 1
}

func (s *IntStack) Pop() interface{} {
	if s.size == 0 {
		return nil
	}
	s.size -= 1
	val := s.stack[s.size]
	s.stack = s.stack[:s.size]
	return val
}

func (s *IntStack) Peek() interface{} {
	if s.size == 0 {
		return nil
	}
	return s.stack[s.size-1]
}

func (s *IntStack) IsEmpty() bool {
	return s.size == 0
}

func (s *IntStack) Size() int {
	return s.size
}
