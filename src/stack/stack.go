package stack

type node struct {
	next  *node
	value interface{}
}

type Stack struct {
	top *node
}

func NewStack() Stack {
	return Stack{nil}
}

func (s *Stack) Init() {
	s.top = nil
}

func (s *Stack) Push(v interface{}) {
	s.top = &node{s.top, v}
}

func (s *Stack) Pop() interface{} {
	if s.top == nil {
		return nil
	}
	v := s.top.value
	s.top = s.top.next
	return v
}

func (s Stack) Top() interface{} {
	if s.top == nil {
		return nil
	}
	return s.top.value
}

func (s Stack) IsEmpty() bool {
	return s.top == nil
}
