package easy

type LinkedList struct {
	v    int
	next *LinkedList
}

func (l *LinkedList) ToSlice() []int {
	s := []int{l.v}
	for l.next != nil {
		s = append(l.next.ToSlice())
	}
	return s
}

func (l *LinkedList) Reverse() *LinkedList {
	return l
}
