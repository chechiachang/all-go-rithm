package medium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinPathSum(t *testing.T) {
	input := [][]int{
		[]int{1},
		[]int{1},
		[]int{1},
	}
	assert.Equal(t, 3, minPathSum(input))
}
