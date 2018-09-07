package easy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToLowerCase(t *testing.T) {
	assert.Equal(t, "", toLowerCase(""))
	assert.Equal(t, "a", toLowerCase("A"))
	assert.Equal(t, "abcd", toLowerCase("ABCD"))
}
