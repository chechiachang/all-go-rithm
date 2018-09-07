package easy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJudgeCircle(t *testing.T) {
	assert.Equal(t, true, judgeCircle(""))
	assert.Equal(t, true, judgeCircle("UUDD"))
	assert.Equal(t, true, judgeCircle("UUDDRRLL"))
}
