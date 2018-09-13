package bst

type TreeNode struct {
	left   *TreeNode
	right  *TreeNode
	parent *TreeNode
	data   string
	key    int
}

type BinarySearchTree struct {
	root *TreeNode
}

func (b *BinarySearchTree) leftMost() *TreeNode {
	// TODO
	return b.root
}

func (b *BinarySearchTree) successor() *TreeNode {
	// TODO
	return b.root
}

func (b *BinarySearchTree) Insert(key int, value string) {
	newNode := &TreeNode{
		key:  key,
		data: value,
	}

	// travesal
	var newParent, current *TreeNode
	current = b.root

	// move newParent util find nil current
	for current != nil {
		newParent = current
		if newNode.key > newParent.key {
			current = newParent.right
		} else {
			current = newParent.left
		}
	}

	// current is nil. add node to current
	newNode.parent = newParent

	if newParent == nil {
		b.root = newNode
	} else if newNode.key > newParent.key {
		newParent.right = newNode
	} else {
		newParent.left = newNode
	}
}

func (b *BinarySearchTree) Search(key int) *TreeNode {
	current := b.root

	// break and return current if
	// current is nil (bottom) -> search failed, return nil
	// current has the same key -> return current
	for current != nil && key != current.key {
		if key > current.key {
			current = current.right
		} else {
			current = current.left
		}
	}

	return current
}

func InorderPrint(n *TreeNode) string {
	if n != nil {
		return InorderPrint(n.left) + n.data + InorderPrint(n.right) // L V R
	}
	return ""
}

func PreorderPrint(n *TreeNode) string {
	if n != nil {
		return n.data + PreorderPrint(n.left) + PreorderPrint(n.right) // V L R
	}
	return ""
}

func PostorderPrint(n *TreeNode) string {
	if n != nil {
		return PostorderPrint(n.left) + PostorderPrint(n.right) + n.data // L R V
	}
	return ""
}

func LevelorderPrint(n *TreeNode) string {
	queue := []*TreeNode{}
	queue = append(queue, n)

	r := ""

	for len(queue) > 0 {
		current := queue[0]
		r += current.data
		queue = queue[1:]

		if current.left != nil {
			queue = append(queue, current.left)
		}

		if current.right != nil {
			queue = append(queue, current.right)
		}
	}
	return r
}
