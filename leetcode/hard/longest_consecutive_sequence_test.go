package hard

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestConsecutiveSequence(t *testing.T) {
	input := []int{}
	assert.Equal(t, 0, LongestConsecutiveSequence(input))

	input = []int{0}
	assert.Equal(t, 1, LongestConsecutiveSequence(input))

	input = []int{100, 4, 200, 3, 1, 2}
	assert.Equal(t, 4, LongestConsecutiveSequence(input))
}
