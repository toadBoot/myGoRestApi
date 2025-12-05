package main

import (
	"testing"
)

func TestSimpleMath(t *testing.T) {
	result := 2 + 2
	if result != 4 {
		t.Errorf("Expected 4, got %d", result)
	}
}
