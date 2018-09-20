package easy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortArrayByParity(t *testing.T) {
	input := []int{3, 1, 2, 4}
	output := []int{2, 4, 3, 1}
	assert.Equal(t, output, sortArrayByParity(input))
}
