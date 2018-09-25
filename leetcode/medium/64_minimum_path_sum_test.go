package medium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinPathSum(t *testing.T) {
	input := [][]int{}
	assert.Equal(t, 0, minPathSum(input))

	input = [][]int{
		[]int{1},
		[]int{1},
		[]int{1},
	}
	assert.Equal(t, 3, minPathSum(input))

	input = [][]int{
		[]int{1, 3, 1},
		[]int{1, 5, 1},
		[]int{4, 2, 1},
	}
	assert.Equal(t, 7, minPathSum(input))

	input = [][]int{
		[]int{1, 2, 3},
		[]int{1, 1, 1},
		[]int{1, 2, 3},
	}
	assert.Equal(t, 7, minPathSum(input))
}
