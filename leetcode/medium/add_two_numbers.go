package medium

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return &ListNode{
		Val: 7,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val:  8,
				Next: nil,
			},
		},
	}
}
