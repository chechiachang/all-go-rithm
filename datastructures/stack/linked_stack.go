package stack

type LinkedStack struct {
	top *LinkedList
}

type LinkedList struct {
	v    int
	next *LinkedList
}

func NewLinkedStack() *LinkedStack {
	return &LinkedStack{}
}

func (s *LinkedStack) Top() int {
	if s.top == nil {
		return 0
	}
	return s.top.v
}

func (s *LinkedStack) Push(i int) {
	t := s.top
	s.top = &LinkedList{
		v:    i,
		next: t,
	}
}

func (s *LinkedStack) Pop() int {
	t := s.top
	s.top = s.top.next
	return t.v
}

func (s *LinkedStack) Size() int {
	size := 0
	n := s.top
	for n != nil {
		size++
		n = n.next
	}
	return size
}
