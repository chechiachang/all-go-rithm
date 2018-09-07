package easy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func testUniqueMorseRepresentations(t *testing.T) {

	words := []string{"gin", "zen", "gig", "msg"}

	assert.Equal(t, 2, uniqueMorseRepresentations(words))
}
