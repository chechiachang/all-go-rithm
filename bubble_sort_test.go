package algorithm

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	if !Equals(BubbleSort(rand), a) {
		t.Error(`TestBubbleSort failed`)
	}
}
