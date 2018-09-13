package bst

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearch(t *testing.T) {
	bst := BinarySearchTree{}

	bst.Insert(1, "1")
	r := bst.Search(1)
	assert.Equal(t, 1, r.key)

	bst.Insert(2, "2")
	r = bst.Search(2)
	assert.Equal(t, 2, r.key)
}

func TestInorder(t *testing.T) {
	bst := BinarySearchTree{}

	bst.Insert(4, "4")
	bst.Insert(5, "5")
	bst.Insert(6, "6")
	bst.Insert(7, "7")
	bst.Insert(3, "3")
	bst.Insert(2, "2")
	bst.Insert(1, "1")
	assert.Equal(t, "1234567", InorderPrint(bst.root))
	assert.Equal(t, "4321567", PreorderPrint(bst.root))
	assert.Equal(t, "1237654", PostorderPrint(bst.root))
	assert.Equal(t, "4352617", LevelorderPrint(bst.root))
}
