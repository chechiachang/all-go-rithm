package easy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSumBruteForce(t *testing.T) {
	assert.Equal(
		t,
		[]int{0, 1},
		twoSumBruteForce([]int{2, 7, 11, 15}, 9),
	)

	assert.Equal(
		t,
		[]int{1, 2},
		twoSumBruteForce([]int{3, 2, 4}, 6),
	)
}

func TestTwoSumHashmap(t *testing.T) {
	assert.Equal(
		t,
		[]int{0, 1},
		twoSumHashmap([]int{2, 7, 11, 15}, 9),
	)

	assert.Equal(
		t,
		[]int{1, 2},
		twoSumHashmap([]int{3, 2, 4}, 6),
	)
}
