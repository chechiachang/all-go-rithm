package demo

import "testing"

func TestLinearSearch(t *testing.T) {
	if LinearSearch(a, 3) != 3 {
		t.Error(`Is3Correct = false"`)
	}
}
