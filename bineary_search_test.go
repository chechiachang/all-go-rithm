package algorithm

import "testing"

func TestBinearySearch(t *testing.T) {
	if BinearySearch(a, 5) != 5 {
		t.Error(`TestBinearySearch = false`)
	}
}
