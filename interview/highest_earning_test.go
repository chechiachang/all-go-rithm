package interview

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHighestEarning(t *testing.T) {
	assert.Equal(t, 0, HighestEarning(""))
	assert.Equal(t, 0, HighestEarning("   "))
	assert.Equal(t, 12, HighestEarning("1 2"))
	assert.Equal(t, 0, HighestEarning("adsf"))
	assert.Equal(t, 3, HighestEarning("1,2"))
	assert.Equal(t, 0, HighestEarning("1,,,,2"))
	assert.Equal(t, 0, HighestEarning("1,a,b,c,2"))
	assert.Equal(t, 3, HighestEarning("1 , 2 "))
	assert.Equal(t, 0, HighestEarning("-1"))
	assert.Equal(t, 0, HighestEarning("0,0,0"))
	assert.Equal(t, 4, HighestEarning("1,3,-1,1"))
	assert.Equal(t, 5, HighestEarning("1,3,-1,-2, 4"))
	assert.Equal(t, 0, HighestEarning("-1,-3,-1,-2,-4"))
	assert.Equal(t, 7, HighestEarning("0,5,0,1,-9,2,0,5"))
}
