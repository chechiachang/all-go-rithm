package bst

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearch(t *testing.T) {
	bst := BinarySearchTree{}

	bst.Insert(0, "0")
	r := bst.Search(0)
	assert.Equal(t, 0, r.key)
	assert.Equal(t, "0", InorderPrint(bst.root))
	assert.Equal(t, "0", PreorderPrint(bst.root))

	bst.Insert(1, "1")
	r = bst.Search(1)
	assert.Equal(t, 1, r.key)
	assert.Equal(t, "01", InorderPrint(bst.root))
	assert.Equal(t, "01", PreorderPrint(bst.root))

	bst.Insert(2, "2")
	bst.Insert(3, "3")
	assert.Equal(t, "0123", InorderPrint(bst.root))
	assert.Equal(t, "0123", PreorderPrint(bst.root))
}
