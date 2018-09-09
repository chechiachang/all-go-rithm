package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArrayStack(t *testing.T) {

	capacity := 3
	s := NewIntArrayStack(capacity)

	s.Push(1) // {1}
	assert.Equal(t, 1, s.Top(), "Top should be 1 after push")

	p := s.Pop() // {}
	assert.Equal(t, 1, p, "Popped should be 1")

	p = s.Pop() // {}
	assert.Equal(t, 0, p, "Nothing to pop")

	s.Push(1) // {1}
	assert.Equal(t, 1, s.Top(), "Top should be 1 after push")

	s.Push(2) // {1, 2}
	assert.Equal(t, 2, s.Top(), "Top should be 2 after push")

	s.Push(3) // {1, 2, 3}
	assert.Equal(t, 3, s.Top(), "Top should be 3 after push")
}
