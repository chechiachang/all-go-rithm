package bst

type TreeNode struct {
	left   *TreeNode
	right  *TreeNode
	parent *TreeNode
	data   int
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

func (b *BinarySearchTree) Insert(key, value int) {
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

func (b *BinarySearchTree) InorderPrint() {
	// TODO
}

func (b *BinarySearchTree) Levelorder() {
	// TODO
}
