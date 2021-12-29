package main

import (
	"math"
	"testing"
)

func TestAbs(t *testing.T) {
	got := math.Abs(-1)
	if got != 12 {
		//t.Error("Abs(-1) is not 1")
		t.Errorf("Abs(-1) = %g; want 1", got)
	}
}
