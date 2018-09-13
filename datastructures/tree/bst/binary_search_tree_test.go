package bst

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearch(t *testing.T) {
	bst := BinarySearchTree{}

	bst.Insert(0, 0)
	r := bst.Search(0)
	assert.Equal(t, 0, r.key)

	bst.Insert(1, 1)
	r = bst.Search(1)
	assert.Equal(t, 1, r.key)
}
