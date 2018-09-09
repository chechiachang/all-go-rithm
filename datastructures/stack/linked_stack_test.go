package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedStack(t *testing.T) {
	s := NewLinkedStack()

	assert.Equal(t, 0, s.Top())
	assert.Equal(t, 0, s.Size())

	s.Push(1)                   // {1}
	assert.Equal(t, 1, s.Pop()) // {0}
	s.Push(2)                   // {2}
	s.Push(3)                   // {3}
	s.Push(4)                   // {4}
	assert.Equal(t, 4, s.Pop()) // {5}
	assert.Equal(t, 2, s.Size())
}
