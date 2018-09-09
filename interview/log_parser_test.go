package interview

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseLog(t *testing.T) {
	assert.Equal(t, "1:errA", ParseLog("1:errA", "2:errB", 0))
	assert.Equal(t, "7:errA", ParseLog(
		"1:errA,6:errB,10:errC",
		"3:errB,7:errA,9:errA",
		3))
}
