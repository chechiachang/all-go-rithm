package algorithm

import "testing"

func TestMerge2(t *testing.T) {
	if !Equals(Merge2(m1, m2), a) {
		t.Error(`TestMerge2 fail`)
	}
}

func TestMerge2Empty(t *testing.T) {
	if !Equals(Merge2(empty1, empty2), empty3) {
		t.Error(`TestMerge2Empty fail`)
	}
}

//func TestMerge(t *testing.T) {
//	Merge(M1, 0, len(M1)/2, len(M1))
//
//	if !Equals(M1, A1) {
//		t.Error(`TestMerge fail`)
//	}
//}

//func TestMergeSort(t *testing.T) {
//	if !Equals(MergeSort(T1, 0, len(T1)), A1) {
//		t.Error(`MergeSort Fail`)
//	}
//}
