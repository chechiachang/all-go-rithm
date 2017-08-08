package algorithm

import "testing"

func test10(t *testing.T) {
	if !Equals(SelectionSort(rand), a) {
		t.Error(`testSelectionSort = false`)
	}
}
