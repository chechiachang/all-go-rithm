package stack

type Stack interface {
	Top() *Stack
	Push(s *Stack)
	Pop(s *Stack)
	Size() int
}
