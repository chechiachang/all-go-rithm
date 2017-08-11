package main

import (
	"testing"
)

func TestHammingDistance(t *testing.T) {
	if HammingDistance(1, 4) != 2 {
		t.Error(`TestHammingDistance (1,4) != 2`)
	}
}
