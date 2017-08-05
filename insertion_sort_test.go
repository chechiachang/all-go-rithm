package demo

import "testing"

func Test(t *testing.T) {
	if !Equals(InsertionSort(rand), a) {
		t.Error(`IsCorrect = false`)
	}
}

// TODO
//func TestInsertionSortResursive(t *testing.T) {
//	InsertionSortRecursive(rand2)
//	if !Equals(rand2, a) {
//		t.Error(`TestInsertionSortRecursive = false`)
//	}
//}
