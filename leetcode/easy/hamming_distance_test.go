package main

import (
	"testing"
)

func TestHammingDistance1(t *testing.T) {
	if HammingDistance(1, 4) != 2 {
		t.Error(`Test HammingDistance (1, 4) != 2`)
	}
}

func TestHammingDistance2(t *testing.T) {
	if HammingDistance(4, 14) != 2 {
		t.Error(`Test HammingDistance (4, 14) != 2`)
	}
}
