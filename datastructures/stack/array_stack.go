package stack

import (
	//"fmt"
	"sync"
)

// An int array stack by slice
type IntArrayStack struct {
	lock  sync.Mutex
	Stack []int
}

func NewIntArrayStack(capacity int) *IntArrayStack {
	return &IntArrayStack{
		sync.Mutex{},
		make([]int, 0),
	}
}

func (s *IntArrayStack) Top() int {
	if len(s.Stack) == 0 { // The stack is empty and has no top
		return 0
	}
	return s.Stack[len(s.Stack)-1]
}

// Push an element to the top of stack if has capacity
func (s *IntArrayStack) Push(i int) {
	s.Stack = append(s.Stack, i)
}

func (s *IntArrayStack) Pop() int {
	if len(s.Stack) == 0 { // The stack is empty and has no top
		return 0
	}
	r := s.Stack[len(s.Stack)-1]
	s.Stack = s.Stack[:len(s.Stack)-1]
	return r
}
