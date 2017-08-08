package algorithm

import "testing"

func test5Digit(t *testing.T) {
	if !Equals(BitAddition(bit1, bit2), bit3) {
		t.Error(`BitAddition = false`)
	}
}
