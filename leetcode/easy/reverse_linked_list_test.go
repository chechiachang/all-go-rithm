package easy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseLinkedList(t *testing.T) {
	l := LinkedList{
		v: 1,
	}

	assert.Equal(t, []int{1}, l.ToSlice())
	assert.Equal(t, []int{1}, l.Reverse().ToSlice())

	l2 := LinkedList{
		v:    1,
		next: &l,
	}

	assert.Equal(t, []int{2, 1}, l2.Reverse().ToSlice())
}
