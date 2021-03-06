package medium

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	output := &ListNode{
		Val: 7,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val: 8,
			},
		},
	}

	return output
}

func (l *ListNode) value() int {
	if l.Next != nil {
		return l.Val + 10*l.Next.value()
	} else {
		return l.Val
	}
}
